# LoggingS3Placement

Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.



## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `LoggingS3PlacementNone`                   | none                                       |
| `LoggingS3PlacementWafDebug`               | waf_debug                                  |
| `LoggingS3PlacementLessThanNilGreaterThan` | <nil>                                      |