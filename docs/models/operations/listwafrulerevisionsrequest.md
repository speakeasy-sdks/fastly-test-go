# ListWafRuleRevisionsRequest


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Include`                                   | *interface{}*                               | :heavy_minus_sign:                          | Include relationships. Optional.            |                                             |
| `PageNumber`                                | **int64*                                    | :heavy_minus_sign:                          | Current page.                               | 1                                           |
| `PageSize`                                  | **int64*                                    | :heavy_minus_sign:                          | Number of records per page.                 | 20                                          |
| `WafRuleID`                                 | *string*                                    | :heavy_check_mark:                          | Alphanumeric string identifying a WAF rule. | 3krg2uUGZzb2W9Euo4moOR                      |