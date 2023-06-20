# ListACLEntriesRequest


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ACLID`                                               | *string*                                              | :heavy_check_mark:                                    | Alphanumeric string identifying a ACL.                | 6tUXdegLTf5BCig0zGFrU3                                |
| `Direction`                                           | [*shared.Direction](../../models/shared/direction.md) | :heavy_minus_sign:                                    | Direction in which to sort results.                   | ascend                                                |
| `Page`                                                | **int64*                                              | :heavy_minus_sign:                                    | Current page.                                         | 1                                                     |
| `PerPage`                                             | **int64*                                              | :heavy_minus_sign:                                    | Number of records per page.                           | 20                                                    |
| `ServiceID`                                           | *string*                                              | :heavy_check_mark:                                    | Alphanumeric string identifying the service.          | SU1Z0isxPaozGVKXdv0eY                                 |
| `Sort`                                                | **string*                                             | :heavy_minus_sign:                                    | Field on which to sort.                               | created                                               |