# UpdateCacheSettingsRequest


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CacheSetting`                                              | [*shared.CacheSetting](../../models/shared/cachesetting.md) | :heavy_minus_sign:                                          | N/A                                                         |                                                             |
| `CacheSettingsName`                                         | *string*                                                    | :heavy_check_mark:                                          | Name for the cache settings object.                         | test-cache-setting                                          |
| `ServiceID`                                                 | *string*                                                    | :heavy_check_mark:                                          | Alphanumeric string identifying the service.                | SU1Z0isxPaozGVKXdv0eY                                       |
| `VersionID`                                                 | *int64*                                                     | :heavy_check_mark:                                          | Integer identifying a service version.                      | 1                                                           |