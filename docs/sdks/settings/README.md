# Settings

## Overview

Handles default settings for a particular version of a service.

<https://developer.fastly.com/reference/api/vcl-services/settings>
### Available Operations

* [GetServiceSettings](#getservicesettings) - Get service settings
* [UpdateServiceSettings](#updateservicesettings) - Update service settings

## GetServiceSettings

Get the settings for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.GetServiceSettingsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Settings.GetServiceSettings(ctx, operations.GetServiceSettingsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.SettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetServiceSettingsRequest](../../models/operations/getservicesettingsrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.GetServiceSettingsSecurity](../../models/operations/getservicesettingssecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.GetServiceSettingsResponse](../../models/operations/getservicesettingsresponse.md), error**


## UpdateServiceSettings

Update the settings for a particular service and version. NOTE: If you override TTLs with custom VCL, any general.default_ttl value will not be honored and the expected behavior may change.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.UpdateServiceSettingsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Settings.UpdateServiceSettings(ctx, operations.UpdateServiceSettingsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        Settings: &shared.Settings{
            GeneralDefaultHost: sdk.String("dignissimos"),
            GeneralDefaultTTL: sdk.Int64(950953),
            GeneralStaleIfError: sdk.Bool(false),
            GeneralStaleIfErrorTTL: sdk.Int64(891523),
        },
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.SettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateServiceSettingsRequest](../../models/operations/updateservicesettingsrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.UpdateServiceSettingsSecurity](../../models/operations/updateservicesettingssecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.UpdateServiceSettingsResponse](../../models/operations/updateservicesettingsresponse.md), error**

