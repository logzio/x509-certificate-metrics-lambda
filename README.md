# X509 Certificate Lambda
A lambda designed to collect X509 certificate metrics from a URL.
The collected metrics will be exported to the Logzio app.
This lambda is using the logzio log extension as the log layer, for additional information visit: https://docs.logz.io/shipping/log-sources/lambda-extensions.html

## Getting Started

To start just press the button and follow the instructions:

[![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/create/template?templateURL=https://logzio-aws-integrations-us-east-1.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/autodeployment.yaml&stackName=logzio-x509-certiricate-metrics)

### Metrics list:
- x509_age (duration in milliseconds)
- x509_expiry (duration in milliseconds)
- x509_start_date (in seconds passed since 1.1.1970)
- x509_end_date (in seconds passed since 1.1.1970)


#### Full list of configurable environment variables

| Environment variable                           | Description                                                                                                                                                                                              |
|------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `LogzioMetricsToken (Required)`                | Token for shipping metrics to your Logz.io account. Find it under Settings > Manage accounts. [_How do I look up my Metrics account token?_](/user-guide/accounts/finding-your-metrics-account-token/)   |
| `LogzioListener (Required)`                    | Your logzio listener url for your region, with no port. i.e: https://listener.logz.io                                                                                                                    |
 | `LambdaName (Optional)`                        | Name for the current lambda function. Default: x509-Certificate-Metrics-Lambda                                                                                                                           |
| `CertificateURL (Required)`                    | The URL to collect x509 certificate metrics from, including port. i.e: https://app.logz.io:443                                                                                                           |
| `LambdaTimeout (Optional)`                     | The amount of time that Lambda allows a function to run before stopping it. Minimum value is 1 second and Maximum value is 900 seconds. We recommend to start with 300 seconds (5 minutes). Default: 300 |
| `CloudWatchEventScheduleExpression (Optional)` | The scheduling expression that determines when and how often the Lambda function runs. We recommend to start with 10 hour rate. Default: rate(10 hours)                                                  |
| `LogzioLogsToken`                              | Your Logz.io log shipping [token](https://app.logz.io/#/dashboard/settings/manage-tokens/data-shipping).                                                                                                 | Required         |


Build and package:
```
GOARCH=amd64 GOOS=linux go build -o bootstrap && zip function.zip bootstrap
```


### Changelog:

- v0.0.1: Initial release