# UpdateCustomVclRequest


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       | Example                                           |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `ServiceID`                                       | *string*                                          | :heavy_check_mark:                                | Alphanumeric string identifying the service.      | SU1Z0isxPaozGVKXdv0eY                             |
| `Vcl`                                             | [*components.Vcl](../../models/components/vcl.md) | :heavy_minus_sign:                                | N/A                                               |                                                   |
| `VclName`                                         | *string*                                          | :heavy_check_mark:                                | The name of this VCL.                             | test-vcl                                          |
| `VersionID`                                       | *int64*                                           | :heavy_check_mark:                                | Integer identifying a service version.            | 1                                                 |