# UpdateHealthcheckRequest


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Healthcheck`                                             | [*shared.Healthcheck](../../models/shared/healthcheck.md) | :heavy_minus_sign:                                        | N/A                                                       |                                                           |
| `HealthcheckName`                                         | *string*                                                  | :heavy_check_mark:                                        | The name of the health check.                             | test-healthcheck                                          |
| `ServiceID`                                               | *string*                                                  | :heavy_check_mark:                                        | Alphanumeric string identifying the service.              | SU1Z0isxPaozGVKXdv0eY                                     |
| `VersionID`                                               | *int64*                                                   | :heavy_check_mark:                                        | Integer identifying a service version.                    | 1                                                         |