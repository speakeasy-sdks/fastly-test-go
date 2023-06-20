# ResourceResponse

OK


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `CreatedAt`                                           | [*time.Time](https://pkg.go.dev/time#Time)            | :heavy_minus_sign:                                    | Date and time in ISO 8601 format.                     | 2020-04-09T18:14:30Z                                  |
| `DeletedAt`                                           | [*time.Time](https://pkg.go.dev/time#Time)            | :heavy_minus_sign:                                    | Date and time in ISO 8601 format.                     | 2020-04-09T18:14:30Z                                  |
| `Href`                                                | **string*                                             | :heavy_minus_sign:                                    | The path to the resource.                             | /resources/stores/object/3vjTN8v1O7nOAY7aNDGOL        |
| `ID`                                                  | **string*                                             | :heavy_minus_sign:                                    | An alphanumeric string identifying the resource link. | 7Lsb7Y76rChV9hSrv3KgFl                                |
| `Name`                                                | **string*                                             | :heavy_minus_sign:                                    | The name of the resource link.                        | test-resource                                         |
| `ResourceID`                                          | **string*                                             | :heavy_minus_sign:                                    | The ID of the underlying linked resource.             | 3vjTN8v1O7nOAY7aNDGOL                                 |
| `ResourceType`                                        | [*TypeResource](../../models/shared/typeresource.md)  | :heavy_minus_sign:                                    | Resource type                                         |                                                       |
| `ServiceID`                                           | **string*                                             | :heavy_minus_sign:                                    | Alphanumeric string identifying the service.          | SU1Z0isxPaozGVKXdv0eY                                 |
| `UpdatedAt`                                           | [*time.Time](https://pkg.go.dev/time#Time)            | :heavy_minus_sign:                                    | Date and time in ISO 8601 format.                     | 2020-04-09T18:14:30Z                                  |
| `Version`                                             | **int64*                                              | :heavy_minus_sign:                                    | Integer identifying a service version.                | 1                                                     |