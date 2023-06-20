# GetWafFirewallRequest


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `FilterServiceVersionNumber`                                      | **string*                                                         | :heavy_minus_sign:                                                | Limit the results returned to a specific service version.         |                                                                   |
| `FirewallID`                                                      | *string*                                                          | :heavy_check_mark:                                                | Alphanumeric string identifying a WAF Firewall.                   | fW7g2uUGZzb2W9Euo4Mo0r                                            |
| `Include`                                                         | [*shared.FirewallInclude](../../models/shared/firewallinclude.md) | :heavy_minus_sign:                                                | Include related objects. Optional.                                |                                                                   |