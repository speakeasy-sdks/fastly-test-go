# Gzip


## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `CacheCondition`                                                                                         | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Name of the cache condition controlling when this configuration applies.                                 | null                                                                                                     |
| `ContentTypes`                                                                                           | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Space-separated list of content types to compress. If you omit this field a default list will be used.   |                                                                                                          |
| `Extensions`                                                                                             | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Space-separated list of file extensions to compress. If you omit this field a default list will be used. |                                                                                                          |
| `Name`                                                                                                   | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Name of the gzip configuration.                                                                          | test-gzip                                                                                                |