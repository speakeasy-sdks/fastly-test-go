# ListServicesRequest


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Direction`                                           | [*shared.Direction](../../models/shared/direction.md) | :heavy_minus_sign:                                    | Direction in which to sort results.                   | ascend                                                |
| `Page`                                                | **int64*                                              | :heavy_minus_sign:                                    | Current page.                                         | 1                                                     |
| `PerPage`                                             | **int64*                                              | :heavy_minus_sign:                                    | Number of records per page.                           | 20                                                    |
| `Sort`                                                | **string*                                             | :heavy_minus_sign:                                    | Field on which to sort.                               | created                                               |