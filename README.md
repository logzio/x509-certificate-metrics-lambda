# X509 Certificate Lambda
A lambda designed to collect X509 certificate metrics from a URL.
The collected metrics will be exported to the Logzio app.


#### Full list of configurable environment variables

| Environment variable                         | Description                                                                                                                                                                                              |
|----------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| LogzioToken (Required)                       | Token for shipping metrics to your Logz.io account. Find it under Settings > Manage accounts. [_How do I look up my Metrics account token?_](/user-guide/accounts/finding-your-metrics-account-token/)   |
| LogzioListener (Required)                    | Your logzio listener url (for metrics).                                                                                                                                                                  |
| LogLevel (Optional)                          | Lambda log level. Default: ERROR                                                                                                                                                                         |
 | LambdaName (Optional)                        | Name for the current lambda function. Default: x509-Certificate-Metrics-Lambda                                                                                                                           |
| CertificateURL (Required)                    | The URL to collect x509 certificate metrics from, including port. i.e: https://app.logz.io:443                                                                                                           |
| LambdaTimeout (Optional)                     | The amount of time that Lambda allows a function to run before stopping it. Minimum value is 1 second and Maximum value is 900 seconds. We recommend to start with 300 seconds (5 minutes). Default: 300 |
| CloudWatchEventScheduleExpression (Optional) | The scheduling expression that determines when and how often the Lambda function runs. We recommend to start with 10 hour rate. Default: rate(10 hours)                                                  |