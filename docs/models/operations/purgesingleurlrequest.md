# PurgeSingleURLRequest


## Fields

| Field                                                                                                                                                                                                  | Type                                                                                                                                                                                                   | Required                                                                                                                                                                                               | Description                                                                                                                                                                                            | Example                                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `CachedURL`                                                                                                                                                                                            | *string*                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                     | URL of object in cache to be purged.                                                                                                                                                                   | www.example.com/path/to/object-to-purge                                                                                                                                                                |
| `FastlySoftPurge`                                                                                                                                                                                      | **int64*                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                     | If present, this header triggers the purge to be 'soft', which marks the affected object as stale rather than making it inaccessible.  Typically set to "1" when used, but the value is not important. | 1                                                                                                                                                                                                      |