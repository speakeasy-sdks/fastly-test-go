# Gzip

## Overview

Gzip configuration allows you to choose resources to automatically compress.  For more information about compressing and decompressing data with Fastly, check out our [concept guide to compression](https://developer.fastly.com/learning/concepts/compression/).

<https://developer.fastly.com/reference/api/vcl-services/gzip>
### Available Operations

* [CreateGzipConfig](#creategzipconfig) - Create a gzip configuration
* [DeleteGzipConfig](#deletegzipconfig) - Delete a gzip configuration
* [GetGzipConfigs](#getgzipconfigs) - Get a gzip configuration
* [ListGzipConfigs](#listgzipconfigs) - List gzip configurations
* [UpdateGzipConfig](#updategzipconfig) - Update a gzip configuration

## CreateGzipConfig

Create a named gzip configuration on a particular service and version.

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

    ctx := context.Background()
    res, err := s.Gzip.CreateGzipConfig(ctx, operations.CreateGzipConfigRequest{
        Gzip: &shared.Gzip{
            CacheCondition: sdk.String("null"),
            ContentTypes: sdk.String("id"),
            Extensions: sdk.String("saepe"),
            Name: sdk.String("test-gzip"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateGzipConfigSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GzipResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateGzipConfigRequest](../../models/operations/creategzipconfigrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.CreateGzipConfigSecurity](../../models/operations/creategzipconfigsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.CreateGzipConfigResponse](../../models/operations/creategzipconfigresponse.md), error**


## DeleteGzipConfig

Delete a named gzip configuration on a particular service and version.

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

    ctx := context.Background()
    res, err := s.Gzip.DeleteGzipConfig(ctx, operations.DeleteGzipConfigRequest{
        GzipName: "test-gzip",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteGzipConfigSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteGzipConfig200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteGzipConfigRequest](../../models/operations/deletegzipconfigrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.DeleteGzipConfigSecurity](../../models/operations/deletegzipconfigsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.DeleteGzipConfigResponse](../../models/operations/deletegzipconfigresponse.md), error**


## GetGzipConfigs

Get the gzip configuration for a particular service, version, and name.

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

    ctx := context.Background()
    res, err := s.Gzip.GetGzipConfigs(ctx, operations.GetGzipConfigsRequest{
        GzipName: "test-gzip",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetGzipConfigsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GzipResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetGzipConfigsRequest](../../models/operations/getgzipconfigsrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetGzipConfigsSecurity](../../models/operations/getgzipconfigssecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetGzipConfigsResponse](../../models/operations/getgzipconfigsresponse.md), error**


## ListGzipConfigs

List all gzip configurations for a particular service and version.

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

    ctx := context.Background()
    res, err := s.Gzip.ListGzipConfigs(ctx, operations.ListGzipConfigsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListGzipConfigsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GzipResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListGzipConfigsRequest](../../models/operations/listgzipconfigsrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ListGzipConfigsSecurity](../../models/operations/listgzipconfigssecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ListGzipConfigsResponse](../../models/operations/listgzipconfigsresponse.md), error**


## UpdateGzipConfig

Update a named gzip configuration on a particular service and version.

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

    ctx := context.Background()
    res, err := s.Gzip.UpdateGzipConfig(ctx, operations.UpdateGzipConfigRequest{
        Gzip: &shared.Gzip{
            CacheCondition: sdk.String("null"),
            ContentTypes: sdk.String("eius"),
            Extensions: sdk.String("aspernatur"),
            Name: sdk.String("test-gzip"),
        },
        GzipName: "test-gzip",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.UpdateGzipConfigSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GzipResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateGzipConfigRequest](../../models/operations/updategzipconfigrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.UpdateGzipConfigSecurity](../../models/operations/updategzipconfigsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.UpdateGzipConfigResponse](../../models/operations/updategzipconfigresponse.md), error**

