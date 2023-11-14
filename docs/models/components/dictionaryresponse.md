# DictionaryResponse


## Fields

| Field                                                                                                                                        | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  | Example                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `CreatedAt`                                                                                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                   | :heavy_minus_sign:                                                                                                                           | Date and time in ISO 8601 format.                                                                                                            | 2020-04-09T18:14:30Z                                                                                                                         |
| `DeletedAt`                                                                                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                   | :heavy_minus_sign:                                                                                                                           | Date and time in ISO 8601 format.                                                                                                            | 2020-04-09T18:14:30Z                                                                                                                         |
| `ID`                                                                                                                                         | **string*                                                                                                                                    | :heavy_minus_sign:                                                                                                                           | N/A                                                                                                                                          | 3vjTN8v1O7nOAY7aNDGOL                                                                                                                        |
| `Name`                                                                                                                                       | **string*                                                                                                                                    | :heavy_minus_sign:                                                                                                                           | Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace). | test_dictionary                                                                                                                              |
| `ServiceID`                                                                                                                                  | **string*                                                                                                                                    | :heavy_minus_sign:                                                                                                                           | N/A                                                                                                                                          | SU1Z0isxPaozGVKXdv0eY                                                                                                                        |
| `UpdatedAt`                                                                                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                   | :heavy_minus_sign:                                                                                                                           | Date and time in ISO 8601 format.                                                                                                            | 2020-04-09T18:14:30Z                                                                                                                         |
| `Version`                                                                                                                                    | **int64*                                                                                                                                     | :heavy_minus_sign:                                                                                                                           | N/A                                                                                                                                          | 1                                                                                                                                            |
| `WriteOnly`                                                                                                                                  | **bool*                                                                                                                                      | :heavy_minus_sign:                                                                                                                           | Determines if items in the dictionary are readable or not.                                                                                   |                                                                                                                                              |