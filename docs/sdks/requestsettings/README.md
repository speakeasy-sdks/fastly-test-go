# RequestSettings
(*RequestSettings*)

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
	fastly "Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestSettings.DeleteRequestSettings(ctx, operations.DeleteRequestSettingsRequest{
        RequestSettingsName: "test-request-setting",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteRequestSettings200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.DeleteRequestSettingsRequest](../../models/operations/deleterequestsettingsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


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
	fastly "Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestSettings.GetRequestSettings(ctx, operations.GetRequestSettingsRequest{
        RequestSettingsName: "test-request-setting",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetRequestSettingsRequest](../../models/operations/getrequestsettingsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


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
	fastly "Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestSettings.ListRequestSettings(ctx, operations.ListRequestSettingsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestSettingsResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListRequestSettingsRequest](../../models/operations/listrequestsettingsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


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
	fastly "Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RequestSettings.UpdateRequestSettings(ctx, operations.UpdateRequestSettingsRequest{
        RequestSettings2: &shared.RequestSettings2{
            Action: shared.RequestSettingsActionPass.ToPointer(),
            BypassBusyWait: fastly.Int64(828657),
            DefaultHost: fastly.String("nemo"),
            ForceMiss: fastly.Int64(924967),
            ForceSsl: fastly.Int64(397533),
            GeoHeaders: fastly.Int64(46007),
            HashKeys: fastly.String("cum"),
            MaxStaleAge: fastly.Int64(232627),
            Name: fastly.String("test-request-setting"),
            RequestCondition: fastly.String("null"),
            TimerSupport: fastly.Int64(449083),
            Xff: shared.RequestSettingsXffLeave.ToPointer(),
        },
        RequestSettingsName: "test-request-setting",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateRequestSettingsRequest](../../models/operations/updaterequestsettingsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.UpdateRequestSettingsResponse](../../models/operations/updaterequestsettingsresponse.md), error**

