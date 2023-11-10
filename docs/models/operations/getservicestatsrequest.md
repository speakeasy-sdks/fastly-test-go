# GetServiceStatsRequest


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   | Example                                       |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `EndTime`                                     | **int64*                                      | :heavy_minus_sign:                            | Epoch timestamp. Limits the results returned. | 1608560817                                    |
| `Month`                                       | **string*                                     | :heavy_minus_sign:                            | 2-digit month.                                | 05                                            |
| `ServiceID`                                   | *string*                                      | :heavy_check_mark:                            | Alphanumeric string identifying the service.  | SU1Z0isxPaozGVKXdv0eY                         |
| `StartTime`                                   | **int64*                                      | :heavy_minus_sign:                            | Epoch timestamp. Limits the results returned. | 1608560817                                    |
| `Year`                                        | **string*                                     | :heavy_minus_sign:                            | 4-digit year.                                 | 2020                                          |