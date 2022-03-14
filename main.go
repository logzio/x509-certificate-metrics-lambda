/*
Code is taken and modified from:
github.com/influxdata/telegraf
*/

package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/influxdata/telegraf/config"
	_tls "github.com/influxdata/telegraf/plugins/common/tls"
	metricsExporter "github.com/logzio/go-metrics-sdk"
	"github.com/pion/dtls/v2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
	"net"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"
)

var infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
var errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) (string, error) {
	infoLogger.Printf("Checking certificate metrics for URL: " + os.Getenv("CERTIFICATE_URL"))
	err := run(ctx)
	if err != nil {
		errorLogger.Printf("Encountered an error: %s", err.Error())
		return "", err
	}

	return "Lambda finished", nil
}

// X509Cert holds the configuration
type X509Cert struct {
	Sources          []string        `toml:"sources"`
	Timeout          config.Duration `toml:"timeout"`
	ServerName       string          `toml:"server_name"`
	ExcludeRootCerts bool            `toml:"exclude_root_certs"`
	tlsCfg           *tls.Config
	_tls.ClientConfig
	locations []*url.URL
	globpaths []*GlobPath
}

type X509CertMetrics struct {
	Tags   map[string]string
	Fields map[string]int64
}

func (c *X509Cert) sourcesToURLs() error {
	for _, source := range c.Sources {
		u, err := url.Parse(source)
		if err != nil {
			return fmt.Errorf("failed to parse cert location - %s", err.Error())
		}
		c.locations = append(c.locations, u)
	}

	return nil
}

func (c *X509Cert) serverName(u *url.URL) (string, error) {
	if c.tlsCfg.ServerName != "" {
		if c.ServerName != "" {
			return "", fmt.Errorf("both server_name (%q) and tls_server_name (%q) are set, but they are mutually exclusive", c.ServerName, c.tlsCfg.ServerName)
		}
		return c.tlsCfg.ServerName, nil
	}
	if c.ServerName != "" {
		return c.ServerName, nil
	}
	return u.Hostname(), nil
}

func (c *X509Cert) getCert(u *url.URL, timeout time.Duration) ([]*x509.Certificate, error) {
	protocol := u.Scheme
	switch u.Scheme {
	case "udp", "udp4", "udp6":
		ipConn, err := net.DialTimeout(u.Scheme, u.Host, timeout)
		if err != nil {
			return nil, err
		}
		defer ipConn.Close()

		serverName, err := c.serverName(u)
		if err != nil {
			return nil, err
		}

		dtlsCfg := &dtls.Config{
			InsecureSkipVerify: true,
			Certificates:       c.tlsCfg.Certificates,
			RootCAs:            c.tlsCfg.RootCAs,
			ServerName:         serverName,
		}
		conn, err := dtls.Client(ipConn, dtlsCfg)
		if err != nil {
			return nil, err
		}
		defer conn.Close()

		rawCerts := conn.ConnectionState().PeerCertificates
		var certs []*x509.Certificate
		for _, rawCert := range rawCerts {
			parsed, err := x509.ParseCertificate(rawCert)
			if err != nil {
				return nil, err
			}

			if parsed != nil {
				certs = append(certs, parsed)
			}
		}

		return certs, nil
	case "https":
		protocol = "tcp"
		fallthrough
	case "tcp", "tcp4", "tcp6":
		ipConn, err := net.DialTimeout(protocol, u.Host, timeout)
		if err != nil {
			return nil, err
		}
		defer ipConn.Close()

		serverName, err := c.serverName(u)
		if err != nil {
			return nil, err
		}
		c.tlsCfg.ServerName = serverName

		c.tlsCfg.InsecureSkipVerify = true
		conn := tls.Client(ipConn, c.tlsCfg)
		defer conn.Close()

		// reset SNI between requests
		defer func() { c.tlsCfg.ServerName = "" }()

		hsErr := conn.Handshake()
		if hsErr != nil {
			return nil, hsErr
		}

		certs := conn.ConnectionState().PeerCertificates
		return certs, nil
	default:
		return nil, fmt.Errorf("unsupported scheme '%s' in location %s", u.Scheme, u.String())
	}
}

func getFields(cert *x509.Certificate, now time.Time) map[string]int64 {
	age := int(now.Sub(cert.NotBefore).Milliseconds())
	expiry := int(cert.NotAfter.Sub(now).Milliseconds())
	startdate := cert.NotBefore.Unix()
	enddate := cert.NotAfter.Unix()

	fields := map[string]int64{
		"age":       int64(age),
		"expiry":    int64(expiry),
		"startdate": startdate,
		"enddate":   enddate,
	}

	return fields
}

func getTags(cert *x509.Certificate, location string) map[string]string {
	tags := map[string]string{
		"source":               location,
		"common_name":          cert.Subject.CommonName,
		"serial_number":        cert.SerialNumber.Text(16),
		"signature_algorithm":  cert.SignatureAlgorithm.String(),
		"public_key_algorithm": cert.PublicKeyAlgorithm.String(),
	}

	if len(cert.Subject.Organization) > 0 {
		tags["organization"] = cert.Subject.Organization[0]
	}
	if len(cert.Subject.OrganizationalUnit) > 0 {
		tags["organizational_unit"] = cert.Subject.OrganizationalUnit[0]
	}
	if len(cert.Subject.Country) > 0 {
		tags["country"] = cert.Subject.Country[0]
	}
	if len(cert.Subject.Province) > 0 {
		tags["province"] = cert.Subject.Province[0]
	}
	if len(cert.Subject.Locality) > 0 {
		tags["locality"] = cert.Subject.Locality[0]
	}

	tags["issuer_common_name"] = cert.Issuer.CommonName
	tags["issuer_serial_number"] = cert.Issuer.SerialNumber

	san := append(cert.DNSNames, cert.EmailAddresses...)
	for _, ip := range cert.IPAddresses {
		san = append(san, ip.String())
	}
	for _, uri := range cert.URIs {
		san = append(san, uri.String())
	}
	tags["san"] = strings.Join(san, ",")

	return tags
}

// Gather adds metrics into the accumulator.
func (c *X509Cert) Gather() (error, []X509CertMetrics) {
	metrics := make([]X509CertMetrics, 0)
	now := time.Now()

	for _, location := range append(c.locations) {
		certs, err := c.getCert(location, time.Duration(c.Timeout))
		if err != nil {
			return err, nil
		}

		for i, cert := range certs {
			fields := getFields(cert, now)
			tags := getTags(cert, location.String())

			// The first certificate is the leaf/end-entity certificate which needs DNS
			// name validation against the URL hostname.
			opts := x509.VerifyOptions{
				Intermediates: x509.NewCertPool(),
				KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
			}
			if i == 0 {
				opts.DNSName, err = c.serverName(location)
				if err != nil {
					return err, nil
				}
				for j, cert := range certs {
					if j != 0 {
						opts.Intermediates.AddCert(cert)
					}
				}
			}
			if c.tlsCfg.RootCAs != nil {
				opts.Roots = c.tlsCfg.RootCAs
			}

			_, err = cert.Verify(opts)
			if err == nil {
				tags["verification"] = "valid"
				fields["verification_code"] = 0
			} else {
				tags["verification"] = "invalid"
				fields["verification_code"] = 1
				fields["verification_error"] = reflect.ValueOf(err.Error()).Int()
			}

			metrics = append(metrics, X509CertMetrics{
				Tags:   tags,
				Fields: fields})
			if c.ExcludeRootCerts {
				break
			}
		}
	}

	return nil, metrics
}

func (c *X509Cert) Init() error {
	err := c.sourcesToURLs()
	if err != nil {
		return err
	}

	tlsCfg, err := c.ClientConfig.TLSConfig()
	if err != nil {
		return err
	}
	if tlsCfg == nil {
		tlsCfg = &tls.Config{}
	}

	if tlsCfg.ServerName != "" && c.ServerName == "" {
		c.ServerName = tlsCfg.ServerName
		tlsCfg.ServerName = ""
	}
	c.tlsCfg = tlsCfg

	return nil
}

func run(ctx context.Context) error {
	cont, err := createController()
	if err != nil {
		panic(fmt.Errorf("error creating controller: %v", err))
	}

	meter := cont.Meter("x509_certificate_metrics")
	if err != nil {
		panic(err)
	}

	cert := &X509Cert{
		Sources:          nil,
		Timeout:          0,
		ServerName:       "",
		ExcludeRootCerts: false,
		tlsCfg:           nil,
		ClientConfig:     _tls.ClientConfig{},
		locations:        nil,
		globpaths:        nil,
	}

	cert.Sources = []string{os.Getenv("CERTIFICATE_URL")}
	err = cert.Init()
	if err != nil {
		return err
	}

	var certMetrics []X509CertMetrics
	err, certMetrics = cert.Gather()
	if err != nil {
		return err
	}

	defer func() {
		handleErr(cont.Stop(ctx))
	}()
	debugLogger.Printf("Extracting metrics from certificates")
	for _, certMetric := range certMetrics {
		labels := make([]attribute.KeyValue, 0)
		for key, value := range certMetric.Tags {
			labels = append(labels, attribute.String(key, value))
		}

		ageObserverCallback := func(_ context.Context, result metric.Int64ObserverResult) {
			result.Observe(certMetric.Fields["age"],
				labels...)
		}

		expiryObserverCallback := func(_ context.Context, result metric.Int64ObserverResult) {
			result.Observe(certMetric.Fields["expiry"],
				labels...)
		}

		startdateObserverCallback := func(_ context.Context, result metric.Int64ObserverResult) {
			result.Observe(certMetric.Fields["startdate"],
				labels...)
		}

		enddateObserverCallback := func(_ context.Context, result metric.Int64ObserverResult) {
			result.Observe(certMetric.Fields["enddate"],
				labels...)
		}

		_ = metric.Must(meter).NewInt64GaugeObserver(
			"x509_age",
			ageObserverCallback,
			metric.WithDescription("x509 certificate age"),
		)

		_ = metric.Must(meter).NewInt64GaugeObserver(
			"x509_expiry",
			expiryObserverCallback,
			metric.WithDescription("x509 certificate expiry"),
		)

		_ = metric.Must(meter).NewInt64GaugeObserver(
			"x509_start_date",
			startdateObserverCallback,
			metric.WithDescription("x509 certificate start date"),
		)

		_ = metric.Must(meter).NewInt64GaugeObserver(
			"x509_end_date",
			enddateObserverCallback,
			metric.WithDescription("x509 certificate end date"),
		)
	}

	/* Time for the controller to send the data before the program ends */
	infoLogger.Printf("Metrics sent successfully to logzio")
	return nil
}

func handleErr(err error) {
	if err != nil {
		panic(fmt.Errorf("something went wrong: %v", err))
	}
}

func createController() (*controller.Controller, error) {
	debugLogger.Printf("Creating controller")
	config := metricsExporter.Config{
		LogzioMetricsListener: os.Getenv("LOGZIO_METRICS_LISTENER"),
		RemoteTimeout:         30 * time.Second,
		PushInterval:          15 * time.Second,
		LogzioMetricsToken:    os.Getenv("LOGZIO_METRICS_TOKEN"),
	}

	return metricsExporter.InstallNewPipeline(config,
		controller.WithCollectPeriod(10*time.Second),
		controller.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
		),
		),
	)
}
