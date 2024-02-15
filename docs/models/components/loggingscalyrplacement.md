# LoggingScalyrPlacement

Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.



## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `LoggingScalyrPlacementNone`     | none                             |
| `LoggingScalyrPlacementWafDebug` | waf_debug                        |