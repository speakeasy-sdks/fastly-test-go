# Attributes


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ModsecRuleID`                                                                        | **int64*                                                                              | :heavy_minus_sign:                                                                    | The ModSecurity rule ID of the associated rule revision.                              |
| `Revision`                                                                            | [*components.WafRuleRevisionOrLatest](../../models/shared/wafrulerevisionorlatest.md) | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `Status`                                                                              | [*components.BulkWafActiveRuleStatus](../../models/shared/bulkwafactiverulestatus.md) | :heavy_minus_sign:                                                                    | Describes the behavior for the particular rule revision within this firewall version. |