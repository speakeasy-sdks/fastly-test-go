# UpdateBackendRequest


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Backend`                                                 | [*components.Backend](../../models/components/backend.md) | :heavy_minus_sign:                                        | N/A                                                       |                                                           |
| `BackendName`                                             | *string*                                                  | :heavy_check_mark:                                        | The name of the backend.                                  | test-backend                                              |
| `ServiceID`                                               | *string*                                                  | :heavy_check_mark:                                        | Alphanumeric string identifying the service.              | SU1Z0isxPaozGVKXdv0eY                                     |
| `VersionID`                                               | *int64*                                                   | :heavy_check_mark:                                        | Integer identifying a service version.                    | 1                                                         |