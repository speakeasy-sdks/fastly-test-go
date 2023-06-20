# VclResponse

OK


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Content`                                                   | **string*                                                   | :heavy_minus_sign:                                          | The VCL code to be included.                                |                                                             |
| `CreatedAt`                                                 | [*time.Time](https://pkg.go.dev/time#Time)                  | :heavy_minus_sign:                                          | Date and time in ISO 8601 format.                           | 2020-04-09T18:14:30Z                                        |
| `DeletedAt`                                                 | [*time.Time](https://pkg.go.dev/time#Time)                  | :heavy_minus_sign:                                          | Date and time in ISO 8601 format.                           | 2020-04-09T18:14:30Z                                        |
| `Main`                                                      | **bool*                                                     | :heavy_minus_sign:                                          | Set to `true` when this is the main VCL, otherwise `false`. |                                                             |
| `Name`                                                      | **string*                                                   | :heavy_minus_sign:                                          | The name of this VCL.                                       | test-vcl                                                    |
| `ServiceID`                                                 | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         | SU1Z0isxPaozGVKXdv0eY                                       |
| `UpdatedAt`                                                 | [*time.Time](https://pkg.go.dev/time#Time)                  | :heavy_minus_sign:                                          | Date and time in ISO 8601 format.                           | 2020-04-09T18:14:30Z                                        |
| `Version`                                                   | **int64*                                                    | :heavy_minus_sign:                                          | N/A                                                         | 1                                                           |