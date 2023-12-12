# ConfigStoreItemResponse


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `CreatedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | Date and time in ISO 8601 format.          | 2020-04-09 18:14:30 +0000 UTC              |
| `DeletedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | Date and time in ISO 8601 format.          | 2020-04-09 18:14:30 +0000 UTC              |
| `ItemKey`                                  | **string*                                  | :heavy_minus_sign:                         | Item key, maximum 256 characters.          | test-key                                   |
| `ItemValue`                                | **string*                                  | :heavy_minus_sign:                         | Item value, maximum 8000 characters.       | test-value                                 |
| `StoreID`                                  | **string*                                  | :heavy_minus_sign:                         | N/A                                        | 7Lsb7Y76rChV9hSrv3KgFl                     |
| `UpdatedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | Date and time in ISO 8601 format.          | 2020-04-09 18:14:30 +0000 UTC              |