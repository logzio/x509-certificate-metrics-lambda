package main

import (
	"fmt"
	"github.com/influxdata/telegraf/config"
	_tls "github.com/influxdata/telegraf/plugins/common/tls"
	"log"
	"os"
	"testing"
	"time"
)

const LOGZIO_APP_URL = "https://app.logz.io:443"

func TestURLNegative(t *testing.T) {
	url := []string{"https://app.logz.io:44"}
	err, _ := UrlTest(url)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestURLPositive(t *testing.T) {
	url := []string{LOGZIO_APP_URL}
	err, _ := UrlTest(url)
	if err != nil {
		t.Error(err.Error())
	}
}

func UrlTest(url []string) (err error, metrics []X509CertMetrics) {
	err, metrics = GatherFromUrl(url)
	if err != nil {
		return fmt.Errorf("invalid URL - gather metrics timed out"), nil
	}

	if len(metrics) == 0 {
		return fmt.Errorf("invalid URL - gather metrics timed out"), nil
	}

	return nil, metrics
}

func TestCertificateTags(t *testing.T) {
	var expected_tags map[string]string
	expected_tags = map[string]string{
		"common_name":          "logz.io",
		"source":               LOGZIO_APP_URL,
		"serial_number":        "f513049a18ad046cdc18a9408275195",
		"public_key_algorithm": "ECDSA",
		"organization":         "Logz.io, Inc.",
		"verification":         "valid",
	}
	url := []string{LOGZIO_APP_URL}
	err, metrics := UrlTest(url)
	if err != nil {
		t.Errorf(err.Error())
	}

	if metrics != nil {
		for key, value := range expected_tags {
			if value != metrics[0].Tags[key] {
				t.Errorf("Unmatching tags for the certificate")
			}
		}
	}
}

func TestCertificateStartEndDate(t *testing.T) {
	var expectedTags map[string]int64
	expectedTags = map[string]int64{
		"startdate": 1639353600,
		"enddate":   1663718399,
	}
	url := []string{LOGZIO_APP_URL}
	err, metrics := UrlTest(url)
	if err != nil {
		t.Errorf(err.Error())
	}

	if metrics != nil {
		for key, value := range expectedTags {
			log.Printf("matching %v with %v", value, metrics[0].Fields[key])
			if value != metrics[0].Fields[key] {
				t.Errorf("Unmatching fields for the certificate")
			}
		}
	}
}

func GatherFromUrl(url []string) (error, []X509CertMetrics) {
	cert := X509Cert{
		Sources:          url,
		Timeout:          config.Duration(15 * time.Second),
		ServerName:       "",
		ExcludeRootCerts: false,
		tlsCfg:           nil,
		ClientConfig:     _tls.ClientConfig{},
		locations:        nil,
		globpaths:        nil,
	}

	log.Printf("Initializing certificate struct")
	cert.Init()
	log.Printf("Gathering metrics")
	err, metrics := cert.Gather()
	return err, metrics
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
