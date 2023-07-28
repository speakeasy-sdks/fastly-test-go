# CacheSettings

## Overview

Configures cache lifetime for objects stored in the Fastly cache, overriding cache freshness information that would otherwise be determined from cache-related headers on the HTTP response. When used in conjunction with conditions, cache settings objects provide detailed control over how long content persists in the cache.

<https://developer.fastly.com/reference/api/vcl-services/cache-settings>
### Available Operations

* [CreateCacheSettings](#createcachesettings) - Create a cache settings object
* [DeleteCacheSettings](#deletecachesettings) - Delete a cache settings object
* [GetCacheSettings](#getcachesettings) - Get a cache settings object
* [ListCacheSettings](#listcachesettings) - List cache settings objects
* [UpdateCacheSettings](#updatecachesettings) - Update a cache settings object

## CreateCacheSettings

Create a cache settings object.

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
    operationSecurity := operations.CreateCacheSettingsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.CacheSettings.CreateCacheSettings(ctx, operations.CreateCacheSettingsRequest{
        CacheSetting1: &shared.CacheSetting1{
            Action: shared.CacheSettingActionCache.ToPointer(),
            CacheCondition: sdk.String("null"),
            Name: sdk.String("test-cache-setting"),
            StaleTTL: sdk.Int64(918236),
            TTL: sdk.Int64(64147),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CacheSettingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateCacheSettingsRequest](../../models/operations/createcachesettingsrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.CreateCacheSettingsSecurity](../../models/operations/createcachesettingssecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.CreateCacheSettingsResponse](../../models/operations/createcachesettingsresponse.md), error**


## DeleteCacheSettings

Delete a specific cache settings object.

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
    operationSecurity := operations.DeleteCacheSettingsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.CacheSettings.DeleteCacheSettings(ctx, operations.DeleteCacheSettingsRequest{
        CacheSettingsName: "test-cache-setting",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteCacheSettings200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteCacheSettingsRequest](../../models/operations/deletecachesettingsrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.DeleteCacheSettingsSecurity](../../models/operations/deletecachesettingssecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.DeleteCacheSettingsResponse](../../models/operations/deletecachesettingsresponse.md), error**


## GetCacheSettings

Get a specific cache settings object.

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
    operationSecurity := operations.GetCacheSettingsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.CacheSettings.GetCacheSettings(ctx, operations.GetCacheSettingsRequest{
        CacheSettingsName: "test-cache-setting",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CacheSettingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetCacheSettingsRequest](../../models/operations/getcachesettingsrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.GetCacheSettingsSecurity](../../models/operations/getcachesettingssecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.GetCacheSettingsResponse](../../models/operations/getcachesettingsresponse.md), error**


## ListCacheSettings

Get a list of all cache settings for a particular service and version.

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
    operationSecurity := operations.ListCacheSettingsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.CacheSettings.ListCacheSettings(ctx, operations.ListCacheSettingsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CacheSettingResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListCacheSettingsRequest](../../models/operations/listcachesettingsrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ListCacheSettingsSecurity](../../models/operations/listcachesettingssecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.ListCacheSettingsResponse](../../models/operations/listcachesettingsresponse.md), error**


## UpdateCacheSettings

Update a specific cache settings object.

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
    operationSecurity := operations.UpdateCacheSettingsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.CacheSettings.UpdateCacheSettings(ctx, operations.UpdateCacheSettingsRequest{
        CacheSetting1: &shared.CacheSetting1{
            Action: shared.CacheSettingActionPass.ToPointer(),
            CacheCondition: sdk.String("null"),
            Name: sdk.String("test-cache-setting"),
            StaleTTL: sdk.Int64(692472),
            TTL: sdk.Int64(565189),
        },
        CacheSettingsName: "test-cache-setting",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CacheSettingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateCacheSettingsRequest](../../models/operations/updatecachesettingsrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.UpdateCacheSettingsSecurity](../../models/operations/updatecachesettingssecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.UpdateCacheSettingsResponse](../../models/operations/updatecachesettingsresponse.md), error**

