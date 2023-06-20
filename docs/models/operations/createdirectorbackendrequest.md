# CreateDirectorBackendRequest


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  | Example                                      |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `BackendName`                                | *string*                                     | :heavy_check_mark:                           | The name of the backend.                     | test-backend                                 |
| `DirectorName`                               | *string*                                     | :heavy_check_mark:                           | Name for the Director.                       | test-director                                |
| `ServiceID`                                  | *string*                                     | :heavy_check_mark:                           | Alphanumeric string identifying the service. | SU1Z0isxPaozGVKXdv0eY                        |
| `VersionID`                                  | *int64*                                      | :heavy_check_mark:                           | Integer identifying a service version.       | 1                                            |