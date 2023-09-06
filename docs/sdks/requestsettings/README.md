# RequestSettings

## Overview

Settings used to customize Fastly's request handling. When used with [Conditions](#condition) the Request Settings object allows you to fine tune how specific types of requests are handled.

<https://developer.fastly.com/reference/api/vcl-services/request-settings>
### Available Operations

* [DeleteRequestSettings](#deleterequestsettings) - Delete a Request Settings object
* [GetRequestSettings](#getrequestsettings) - Get a Request Settings object
* [ListRequestSettings](#listrequestsettings) - List Request Settings objects
* [UpdateRequestSettings](#updaterequestsettings) - Update a Request Settings object

## DeleteRequestSettings

Removes the specified Request Settings object.

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
    operationSecurity := operations.DeleteRequestSettingsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.RequestSettings.DeleteRequestSettings(ctx, operations.DeleteRequestSettingsRequest{
        RequestSettingsName: "test-request-setting",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteRequestSettings200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteRequestSettingsRequest](../../models/operations/deleterequestsettingsrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.DeleteRequestSettingsSecurity](../../models/operations/deleterequestsettingssecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.DeleteRequestSettingsResponse](../../models/operations/deleterequestsettingsresponse.md), error**


## GetRequestSettings

Gets the specified Request Settings object.

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
    operationSecurity := operations.GetRequestSettingsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.RequestSettings.GetRequestSettings(ctx, operations.GetRequestSettingsRequest{
        RequestSettingsName: "test-request-setting",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetRequestSettingsRequest](../../models/operations/getrequestsettingsrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.GetRequestSettingsSecurity](../../models/operations/getrequestsettingssecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.GetRequestSettingsResponse](../../models/operations/getrequestsettingsresponse.md), error**


## ListRequestSettings

Returns a list of all Request Settings objects for the given service and version.

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
    operationSecurity := operations.ListRequestSettingsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.RequestSettings.ListRequestSettings(ctx, operations.ListRequestSettingsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestSettingsResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListRequestSettingsRequest](../../models/operations/listrequestsettingsrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.ListRequestSettingsSecurity](../../models/operations/listrequestsettingssecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.ListRequestSettingsResponse](../../models/operations/listrequestsettingsresponse.md), error**


## UpdateRequestSettings

Updates the specified Request Settings object.

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
    operationSecurity := operations.UpdateRequestSettingsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.RequestSettings.UpdateRequestSettings(ctx, operations.UpdateRequestSettingsRequest{
        RequestSettings2: &shared.RequestSettings2{
            Action: shared.RequestSettingsActionPass.ToPointer(),
            BypassBusyWait: sdk.Int64(424089),
            DefaultHost: sdk.String("ducimus"),
            ForceMiss: sdk.Int64(554688),
            ForceSsl: sdk.Int64(427834),
            GeoHeaders: sdk.Int64(287051),
            HashKeys: sdk.String("possimus"),
            MaxStaleAge: sdk.Int64(706575),
            Name: sdk.String("test-request-setting"),
            RequestCondition: sdk.String("null"),
            TimerSupport: sdk.Int64(738227),
            Xff: shared.RequestSettingsXffAppend.ToPointer(),
        },
        RequestSettingsName: "test-request-setting",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateRequestSettingsRequest](../../models/operations/updaterequestsettingsrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.UpdateRequestSettingsSecurity](../../models/operations/updaterequestsettingssecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.UpdateRequestSettingsResponse](../../models/operations/updaterequestsettingsresponse.md), error**

