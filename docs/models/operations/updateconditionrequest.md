# UpdateConditionRequest


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Condition`                                                   | [*components.Condition](../../models/components/condition.md) | :heavy_minus_sign:                                            | N/A                                                           |                                                               |
| `ConditionName`                                               | *string*                                                      | :heavy_check_mark:                                            | Name of the condition. Required.                              | test-condition                                                |
| `ServiceID`                                                   | *string*                                                      | :heavy_check_mark:                                            | Alphanumeric string identifying the service.                  | SU1Z0isxPaozGVKXdv0eY                                         |
| `VersionID`                                                   | *int64*                                                       | :heavy_check_mark:                                            | Integer identifying a service version.                        | 1                                                             |