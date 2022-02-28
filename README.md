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

| Environment variable                         | Description                                                                                                                                                                                              |
|----------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `LogzioToken (Required)`                       | Token for shipping metrics to your Logz.io account. Find it under Settings > Manage accounts. [_How do I look up my Metrics account token?_](/user-guide/accounts/finding-your-metrics-account-token/)   |
| `LogzioListener (Required)`                    | Your logzio listener url (for metrics).                                                                                                                                                                  |
 | `LambdaName (Optional)`                        | Name for the current lambda function. Default: x509-Certificate-Metrics-Lambda                                                                                                                           |
| `CertificateURL (Required)`                    | The URL to collect x509 certificate metrics from, including port. i.e: https://app.logz.io:443                                                                                                           |
| `LambdaTimeout (Optional)`                     | The amount of time that Lambda allows a function to run before stopping it. Minimum value is 1 second and Maximum value is 900 seconds. We recommend to start with 300 seconds (5 minutes). Default: 300 |
| `CloudWatchEventScheduleExpression (Optional)` | The scheduling expression that determines when and how often the Lambda function runs. We recommend to start with 10 hour rate. Default: rate(10 hours)                                                  |
| `LOGZIO_LOGS_TOKEN`                          | Your Logz.io log shipping [token](https://app.logz.io/#/dashboard/settings/manage-tokens/data-shipping).                                                                                                                                                                                                                                    | Required         |
| `LOGZIO_LISTENER`                            | Your  Logz.io listener address, with port 8070 (http) or 8071 (https). For example: `https://listener.logz.io:8071`                                                                                                                                                                                                                         | Required         |
| `LOGS_EXT_LOG_LEVEL`                         | Log level of the extension. Can be set to one of the following: `debug`, `info`, `warn`, `error`, `fatal`, `panic`.                                                                                                                                                                                                                         | Default: `info`  |
| `ENABLE_EXTENSION_LOGS`                      | Set to `true` if you wish the extension logs will be shipped to your Logz.io account.                                                                                                                                                                                                                                                       | Default: `false` |
| `ENABLE_PLATFORM_LOGS`                       | The platform log captures runtime or execution environment errors. Set to `true` if you wish the platform logs will be shipped to your Logz.io account.                                                                                                                                                                                     | Default: `false` |
| `GROK_PATTERNS`                              | Must be set with `LOGS_FORMAT`. Use this if you want to parse your logs into fields. A minified JSON list that contains the field name and the regex that will match the field. To understand more see the [parsing logs](https://github.com/logzio/logzio-lambda-extensions/tree/main/logzio-lambda-extensions-logs#parsing-logs) section. | -                |
| `LOGS_FORMAT`                                | Must be set with `GROK_PATTERNS`. Use this if you want to parse your logs into fields. The format in which the logs will appear, in accordance to grok conventions. To understand more see the [parsing logs](https://github.com/logzio/logzio-lambda-extensions/tree/main/logzio-lambda-extensions-logs#parsing-logs) section.             | -                |
| `CUSTOM_FIELDS`                              | Include additional fields with every message sent, formatted as `fieldName1=fieldValue1;fieldName2=fieldValue2` (**NO SPACES**). A custom key that clashes with a key from the log itself will be ignored.                                                                                                                                  | -                |


Build and package:
```
GOARCH=amd64 GOOS=linux go build -o bootstrap && zip function.zip bootstrap
```


### Changelog:

- v0.0.1: Initial release