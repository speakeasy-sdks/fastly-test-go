# GetWafRuleRequest


## Fields

| Field                                                                                                            | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `Include`                                                                                                        | **string*                                                                                                        | :heavy_minus_sign:                                                                                               | Include relationships. Optional, comma-separated values. Permitted values: `waf_tags` and `waf_rule_revisions`.<br/> | waf_tags,waf_rule_revisions                                                                                      |
| `WafRuleID`                                                                                                      | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Alphanumeric string identifying a WAF rule.                                                                      | 3krg2uUGZzb2W9Euo4moOR                                                                                           |