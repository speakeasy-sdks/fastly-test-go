# GetWafRuleRevisionRequest


## Fields

| Field                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                          | Example                                                                                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Include`                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                   | Include relationships. Optional, comma-separated values. Permitted values: `waf_rule`, `vcl`, and `source`. The `vcl` and `source` relationships show the WAF VCL and corresponding ModSecurity source. These fields are blank unless the relationship is included.<br/> | source,vcl,waf_rule                                                                                                                                                                                                                                                  |
| `WafRuleID`                                                                                                                                                                                                                                                          | *string*                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                   | Alphanumeric string identifying a WAF rule.                                                                                                                                                                                                                          | 3krg2uUGZzb2W9Euo4moOR                                                                                                                                                                                                                                               |
| `WafRuleRevisionNumber`                                                                                                                                                                                                                                              | *int64*                                                                                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                                   | Revision number.                                                                                                                                                                                                                                                     | 2                                                                                                                                                                                                                                                                    |