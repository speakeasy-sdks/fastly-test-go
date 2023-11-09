# LoggingFtpResponseFormatVersion

The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.



## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `LoggingFtpResponseFormatVersionOne` | 1                                    |
| `LoggingFtpResponseFormatVersionTwo` | 2                                    |