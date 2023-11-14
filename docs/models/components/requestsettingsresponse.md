# RequestSettingsResponse


## Fields

| Field                                                                                                         | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   | Example                                                                                                       |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `Action`                                                                                                      | [*components.RequestSettingsResponseAction](../../models/components/requestsettingsresponseaction.md)         | :heavy_minus_sign:                                                                                            | Allows you to terminate request handling and immediately perform an action.                                   |                                                                                                               |
| `BypassBusyWait`                                                                                              | **int64*                                                                                                      | :heavy_minus_sign:                                                                                            | Disable collapsed forwarding, so you don't wait for other objects to origin.                                  |                                                                                                               |
| `CreatedAt`                                                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                                                    | :heavy_minus_sign:                                                                                            | Date and time in ISO 8601 format.                                                                             | 2020-04-09T18:14:30Z                                                                                          |
| `DefaultHost`                                                                                                 | **string*                                                                                                     | :heavy_minus_sign:                                                                                            | Sets the host header.                                                                                         |                                                                                                               |
| `DeletedAt`                                                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                                                    | :heavy_minus_sign:                                                                                            | Date and time in ISO 8601 format.                                                                             | 2020-04-09T18:14:30Z                                                                                          |
| `ForceMiss`                                                                                                   | **int64*                                                                                                      | :heavy_minus_sign:                                                                                            | Allows you to force a cache miss for the request. Replaces the item in the cache if the content is cacheable. |                                                                                                               |
| `ForceSsl`                                                                                                    | **int64*                                                                                                      | :heavy_minus_sign:                                                                                            | Forces the request use SSL (redirects a non-SSL to SSL).                                                      |                                                                                                               |
| `GeoHeaders`                                                                                                  | **int64*                                                                                                      | :heavy_minus_sign:                                                                                            | Injects Fastly-Geo-Country, Fastly-Geo-City, and Fastly-Geo-Region into the request headers.                  |                                                                                                               |
| `HashKeys`                                                                                                    | **string*                                                                                                     | :heavy_minus_sign:                                                                                            | Comma separated list of varnish request object fields that should be in the hash key.                         |                                                                                                               |
| `MaxStaleAge`                                                                                                 | **int64*                                                                                                      | :heavy_minus_sign:                                                                                            | How old an object is allowed to be to serve stale-if-error or stale-while-revalidate.                         |                                                                                                               |
| `Name`                                                                                                        | **string*                                                                                                     | :heavy_minus_sign:                                                                                            | Name for the request settings.                                                                                | test-request-setting                                                                                          |
| `RequestCondition`                                                                                            | **string*                                                                                                     | :heavy_minus_sign:                                                                                            | Condition which, if met, will select this configuration during a request. Optional.                           | null                                                                                                          |
| `ServiceID`                                                                                                   | **string*                                                                                                     | :heavy_minus_sign:                                                                                            | N/A                                                                                                           | SU1Z0isxPaozGVKXdv0eY                                                                                         |
| `TimerSupport`                                                                                                | **int64*                                                                                                      | :heavy_minus_sign:                                                                                            | Injects the X-Timer info into the request for viewing origin fetch durations.                                 |                                                                                                               |
| `UpdatedAt`                                                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                                                    | :heavy_minus_sign:                                                                                            | Date and time in ISO 8601 format.                                                                             | 2020-04-09T18:14:30Z                                                                                          |
| `Version`                                                                                                     | **int64*                                                                                                      | :heavy_minus_sign:                                                                                            | N/A                                                                                                           | 1                                                                                                             |
| `Xff`                                                                                                         | [*components.RequestSettingsResponseXff](../../models/components/requestsettingsresponsexff.md)               | :heavy_minus_sign:                                                                                            | Short for X-Forwarded-For.                                                                                    |                                                                                                               |