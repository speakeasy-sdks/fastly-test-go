# ListWafTagsRequest


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `FilterName`                                                  | **string*                                                     | :heavy_minus_sign:                                            | Limit the returned tags to a specific name.                   |                                                               |
| `Include`                                                     | [*shared.WafTagInclude](../../models/shared/waftaginclude.md) | :heavy_minus_sign:                                            | Include relationships. Optional.                              | waf_rules                                                     |
| `PageNumber`                                                  | **int64*                                                      | :heavy_minus_sign:                                            | Current page.                                                 | 1                                                             |
| `PageSize`                                                    | **int64*                                                      | :heavy_minus_sign:                                            | Number of records per page.                                   | 20                                                            |