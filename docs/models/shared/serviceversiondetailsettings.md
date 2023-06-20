# ServiceVersionDetailSettings

List of default settings for this service.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `GeneralDefaultHost`                                                         | **string*                                                                    | :heavy_minus_sign:                                                           | The default host name for the version.                                       |
| `GeneralDefaultTTL`                                                          | **int64*                                                                     | :heavy_minus_sign:                                                           | The default time-to-live (TTL) for the version.                              |
| `GeneralStaleIfError`                                                        | **bool*                                                                      | :heavy_minus_sign:                                                           | Enables serving a stale object if there is an error.                         |
| `GeneralStaleIfErrorTTL`                                                     | **int64*                                                                     | :heavy_minus_sign:                                                           | The default time-to-live (TTL) for serving the stale object for the version. |