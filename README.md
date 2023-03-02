# X509 Certificate Lambda
A lambda designed to collect X509 certificate metrics from a URL.
The collected metrics will be exported to the Logzio app.
This lambda is using the logzio log extension as the log layer, for additional information visit: https://docs.logz.io/shipping/log-sources/lambda-extensions.html

## Getting Started

To start just press the button and follow the instructions:

| Region           | Deployment                                                                                                                                                                                                                                                                                                                                                             |
|------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `us-east-1`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/create/template?templateURL=https://logzio-aws-integrations-us-east-1.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert)           | 
| `us-east-2`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/create/template?templateURL=https://logzio-aws-integrations-us-east-2.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert)           | 
| `us-west-1`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=us-west-1#/stacks/create/template?templateURL=https://logzio-aws-integrations-us-west-1.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert)           | 
| `us-west-2`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=us-west-2#/stacks/create/template?templateURL=https://logzio-aws-integrations-us-west-2.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert)           | 
| `eu-central-1`   | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-central-1#/stacks/create/template?templateURL=https://logzio-aws-integrations-eu-central-1.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert)     | 
| `eu-north-1`     | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-north-1#/stacks/create/template?templateURL=https://logzio-aws-integrations-eu-north-1.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert)         | 
| `eu-west-1`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-1#/stacks/create/template?templateURL=https://logzio-aws-integrations-eu-west-1.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert)           | 
| `eu-west-2`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-2#/stacks/create/template?templateURL=https://logzio-aws-integrations-eu-west-2.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert)           | 
| `eu-west-3`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-3#/stacks/create/template?templateURL=https://logzio-aws-integrations-eu-west-3.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert)           | 
| `sa-east-1`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=sa-east-1#/stacks/create/template?templateURL=https://logzio-aws-integrations-sa-east-1.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert)           | 
| `ap-northeast-1` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-1#/stacks/create/template?templateURL=https://logzio-aws-integrations-ap-northeast-1.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert) | 
| `ap-northeast-2` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-2#/stacks/create/template?templateURL=https://logzio-aws-integrations-ap-northeast-2.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert) | 
| `ap-northeast-3` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-3#/stacks/create/template?templateURL=https://logzio-aws-integrations-ap-northeast-3.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert) | 
| `ap-south-1`     | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-south-1#/stacks/create/template?templateURL=https://logzio-aws-integrations-ap-south-1.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert)         | 
| `ap-southeast-1` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-1#/stacks/create/template?templateURL=https://logzio-aws-integrations-ap-southeast-1.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert) | 
| `ap-southeast-2` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-2#/stacks/create/template?templateURL=https://logzio-aws-integrations-ap-southeast-2.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert) | 
| `ca-central-1`   | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ca-central-1#/stacks/create/template?templateURL=https://logzio-aws-integrations-ca-central-1.s3.amazonaws.com/x509-certificate-metricts-auto-deployment/0.0.3/sam-template.yaml&stackName=logizo-x509-cert)     |

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
| `CertificateURL (Required)`                    | The URL to collect x509 certificate metrics from, including port. i.e: https://app.logz.io:443                                                                                                           |
| `LambdaTimeout (Optional)`                     | The amount of time that Lambda allows a function to run before stopping it. Minimum value is 1 second and Maximum value is 900 seconds. We recommend to start with 300 seconds (5 minutes). Default: 300 |
| `CloudWatchEventScheduleExpression (Optional)` | The scheduling expression that determines when and how often the Lambda function runs. We recommend to start with 10 hour rate. Rate below 6 minutes will cause the lambda to behave unexpectedly due to cold start and custom resource invocation. Default: rate(10 hours)|
| `LogzioLogsToken (Required)`                   | Your Logz.io log shipping [token](https://app.logz.io/#/dashboard/settings/manage-tokens/data-shipping).                                                                                                 |


Build and package:
```
GOARCH=amd64 GOOS=linux go build -o bootstrap && zip function.zip bootstrap
```


### Changelog:

- 0.0.3:
  - Auto deployment upon release.
  - Support more regions.
  - Remove LambdaName parameter. Lambda name will derive from Stack Name.

- v0.0.2:
  * Changed IAMRole and EventRule names to have dynamic names.
  * Added custom resource call in autodeployment.yaml to invoke the lambda upon creation.

- v0.0.1: Initial release