# LoggingSyslogResponse

OK


## Fields

| Field                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                | Example                                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Address`                                                                                                                                                                                                                  | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | A hostname or IPv4 address.                                                                                                                                                                                                | example.com                                                                                                                                                                                                                |
| `CreatedAt`                                                                                                                                                                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                         | Date and time in ISO 8601 format.                                                                                                                                                                                          | 2020-04-09T18:14:30Z                                                                                                                                                                                                       |
| `DeletedAt`                                                                                                                                                                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                         | Date and time in ISO 8601 format.                                                                                                                                                                                          | 2020-04-09T18:14:30Z                                                                                                                                                                                                       |
| `Format`                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).                                                                                                                                        | %h %l %u %t "%r" %&gt;s %b                                                                                                                                                                                                 |
| `FormatVersion`                                                                                                                                                                                                            | [*LoggingSyslogResponseFormatVersion](../../models/shared/loggingsyslogresponseformatversion.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                         | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.<br/> | 2                                                                                                                                                                                                                          |
| `Hostname`                                                                                                                                                                                                                 | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The hostname used for the syslog endpoint.                                                                                                                                                                                 |                                                                                                                                                                                                                            |
| `Ipv4`                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The IPv4 address used for the syslog endpoint.                                                                                                                                                                             |                                                                                                                                                                                                                            |
| `MessageType`                                                                                                                                                                                                              | [*LoggingMessageType](../../models/shared/loggingmessagetype.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                         | How the message should be formatted.                                                                                                                                                                                       | classic                                                                                                                                                                                                                    |
| `Name`                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The name for the real-time logging configuration.                                                                                                                                                                          | test-log-endpoint                                                                                                                                                                                                          |
| `Placement`                                                                                                                                                                                                                | [*LoggingSyslogResponsePlacement](../../models/shared/loggingsyslogresponseplacement.md)                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                         | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.<br/>              | null                                                                                                                                                                                                                       |
| `Port`                                                                                                                                                                                                                     | **int64*                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                         | The port number.                                                                                                                                                                                                           |                                                                                                                                                                                                                            |
| `ResponseCondition`                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The name of an existing condition in the configured endpoint, or leave blank to always execute.                                                                                                                            | null                                                                                                                                                                                                                       |
| `ServiceID`                                                                                                                                                                                                                | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | N/A                                                                                                                                                                                                                        | SU1Z0isxPaozGVKXdv0eY                                                                                                                                                                                                      |
| `TLSCaCert`                                                                                                                                                                                                                | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | A secure certificate to authenticate a server with. Must be in PEM format.                                                                                                                                                 |                                                                                                                                                                                                                            |
| `TLSClientCert`                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The client certificate used to make authenticated requests. Must be in PEM format.                                                                                                                                         |                                                                                                                                                                                                                            |
| `TLSClientKey`                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The client private key used to make authenticated requests. Must be in PEM format.                                                                                                                                         |                                                                                                                                                                                                                            |
| `TLSHostname`                                                                                                                                                                                                              | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.                                                 |                                                                                                                                                                                                                            |
| `Token`                                                                                                                                                                                                                    | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | Whether to prepend each message with a specific token.                                                                                                                                                                     |                                                                                                                                                                                                                            |
| `UpdatedAt`                                                                                                                                                                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                         | Date and time in ISO 8601 format.                                                                                                                                                                                          | 2020-04-09T18:14:30Z                                                                                                                                                                                                       |
| `UseTLS`                                                                                                                                                                                                                   | [*LoggingUseTLS](../../models/shared/loggingusetls.md)                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                         | Whether to use TLS.                                                                                                                                                                                                        |                                                                                                                                                                                                                            |
| `Version`                                                                                                                                                                                                                  | **int64*                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                         | N/A                                                                                                                                                                                                                        | 1                                                                                                                                                                                                                          |