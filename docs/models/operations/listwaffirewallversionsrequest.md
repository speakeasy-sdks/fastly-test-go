# ListWafFirewallVersionsRequest


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     | Example                                         |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `FirewallID`                                    | *string*                                        | :heavy_check_mark:                              | Alphanumeric string identifying a WAF Firewall. | fW7g2uUGZzb2W9Euo4Mo0r                          |
| `Include`                                       | **string*                                       | :heavy_minus_sign:                              | Include relationships. Optional.                | waf_firewall                                    |
| `PageNumber`                                    | **int64*                                        | :heavy_minus_sign:                              | Current page.                                   | 1                                               |
| `PageSize`                                      | **int64*                                        | :heavy_minus_sign:                              | Number of records per page.                     | 20                                              |