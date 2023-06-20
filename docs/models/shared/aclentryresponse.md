# ACLEntryResponse

OK


## Fields

| Field                                                                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                                                                                                                                                                                                          | Required                                                                                                                                                                                                                                                                                                                                      | Description                                                                                                                                                                                                                                                                                                                                   | Example                                                                                                                                                                                                                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ACLID`                                                                                                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                            | N/A                                                                                                                                                                                                                                                                                                                                           | 6tUXdegLTf5BCig0zGFrU3                                                                                                                                                                                                                                                                                                                        |
| `Comment`                                                                                                                                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                            | A freeform descriptive note.                                                                                                                                                                                                                                                                                                                  |                                                                                                                                                                                                                                                                                                                                               |
| `CreatedAt`                                                                                                                                                                                                                                                                                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                            | Date and time in ISO 8601 format.                                                                                                                                                                                                                                                                                                             | 2020-04-09T18:14:30Z                                                                                                                                                                                                                                                                                                                          |
| `DeletedAt`                                                                                                                                                                                                                                                                                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                            | Date and time in ISO 8601 format.                                                                                                                                                                                                                                                                                                             | 2020-04-09T18:14:30Z                                                                                                                                                                                                                                                                                                                          |
| `ID`                                                                                                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                            | N/A                                                                                                                                                                                                                                                                                                                                           | 6yxNzlOpW1V7JfSwvLGtOc                                                                                                                                                                                                                                                                                                                        |
| `IP`                                                                                                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                            | An IP address.                                                                                                                                                                                                                                                                                                                                | 127.0.0.1                                                                                                                                                                                                                                                                                                                                     |
| `Negated`                                                                                                                                                                                                                                                                                                                                     | [*ACLEntryResponseNegated](../../models/shared/aclentryresponsenegated.md)                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                            | Whether to negate the match. Useful primarily when creating individual exceptions to larger subnets.                                                                                                                                                                                                                                          | 0                                                                                                                                                                                                                                                                                                                                             |
| `ServiceID`                                                                                                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                            | N/A                                                                                                                                                                                                                                                                                                                                           | SU1Z0isxPaozGVKXdv0eY                                                                                                                                                                                                                                                                                                                         |
| `Subnet`                                                                                                                                                                                                                                                                                                                                      | **int64*                                                                                                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                            | Number of bits for the subnet mask applied to the IP address. For IPv4 addresses, a value of 32 represents the smallest subnet mask (1 address), 24 represents a class C subnet mask (256 addresses), 16 represents a class B subnet mask (65k addresses), and 8 is class A subnet mask (16m addresses). If not provided, no mask is applied. | 8                                                                                                                                                                                                                                                                                                                                             |
| `UpdatedAt`                                                                                                                                                                                                                                                                                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                            | Date and time in ISO 8601 format.                                                                                                                                                                                                                                                                                                             | 2020-04-09T18:14:30Z                                                                                                                                                                                                                                                                                                                          |