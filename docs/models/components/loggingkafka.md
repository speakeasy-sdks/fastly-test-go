# LoggingKafka


## Fields

| Field                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                | Example                                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AuthMethod`                                                                                                                                                                                                               | [*components.AuthMethod](../../models/components/authmethod.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                         | SASL authentication method.                                                                                                                                                                                                |                                                                                                                                                                                                                            |
| `Brokers`                                                                                                                                                                                                                  | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | A comma-separated list of IP addresses or hostnames of Kafka brokers. Required.                                                                                                                                            |                                                                                                                                                                                                                            |
| `CompressionCodec`                                                                                                                                                                                                         | [*components.LoggingKafkaCompressionCodec](../../models/components/loggingkafkacompressioncodec.md)                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                         | The codec used for compression of your logs.                                                                                                                                                                               |                                                                                                                                                                                                                            |
| `Format`                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).                                                                                                                                        | %h %l %u %t "%r" %&gt;s %b                                                                                                                                                                                                 |
| `FormatVersion`                                                                                                                                                                                                            | [*components.LoggingKafkaFormatVersion](../../models/components/loggingkafkaformatversion.md)                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                         | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.<br/> | 2                                                                                                                                                                                                                          |
| `Name`                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The name for the real-time logging configuration.                                                                                                                                                                          | test-log-endpoint                                                                                                                                                                                                          |
| `ParseLogKeyvals`                                                                                                                                                                                                          | **bool*                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                         | Enables parsing of key=value tuples from the beginning of a logline, turning them into [record headers](https://cwiki.apache.org/confluence/display/KAFKA/KIP-82+-+Add+Record+Headers).                                    |                                                                                                                                                                                                                            |
| `Password`                                                                                                                                                                                                                 | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | SASL password.                                                                                                                                                                                                             |                                                                                                                                                                                                                            |
| `Placement`                                                                                                                                                                                                                | [*components.LoggingKafkaPlacement](../../models/components/loggingkafkaplacement.md)                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                         | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.<br/>              | null                                                                                                                                                                                                                       |
| `RequestMaxBytes`                                                                                                                                                                                                          | **int64*                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                         | The maximum number of bytes sent in one request. Defaults `0` (no limit).                                                                                                                                                  |                                                                                                                                                                                                                            |
| `RequiredAcks`                                                                                                                                                                                                             | [*components.RequiredAcks](../../models/components/requiredacks.md)                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                         | The number of acknowledgements a leader must receive before a write is considered successful.                                                                                                                              |                                                                                                                                                                                                                            |
| `ResponseCondition`                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The name of an existing condition in the configured endpoint, or leave blank to always execute.                                                                                                                            | null                                                                                                                                                                                                                       |
| `TLSCaCert`                                                                                                                                                                                                                | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | A secure certificate to authenticate a server with. Must be in PEM format.                                                                                                                                                 |                                                                                                                                                                                                                            |
| `TLSClientCert`                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The client certificate used to make authenticated requests. Must be in PEM format.                                                                                                                                         |                                                                                                                                                                                                                            |
| `TLSClientKey`                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The client private key used to make authenticated requests. Must be in PEM format.                                                                                                                                         |                                                                                                                                                                                                                            |
| `TLSHostname`                                                                                                                                                                                                              | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.                                                 |                                                                                                                                                                                                                            |
| `Topic`                                                                                                                                                                                                                    | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The Kafka topic to send logs to. Required.                                                                                                                                                                                 |                                                                                                                                                                                                                            |
| `UseTLS`                                                                                                                                                                                                                   | [*components.LoggingUseTLS](../../models/components/loggingusetls.md)                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                         | Whether to use TLS.                                                                                                                                                                                                        |                                                                                                                                                                                                                            |
| `User`                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | SASL user.                                                                                                                                                                                                                 |                                                                                                                                                                                                                            |