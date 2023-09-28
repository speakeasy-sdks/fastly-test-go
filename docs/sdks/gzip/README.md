# Gzip
(*Gzip*)

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
    res, err := s.Gzip.CreateGzipConfig(ctx, operations.CreateGzipConfigRequest{
        Gzip: &shared.Gzip{
            CacheCondition: fastly.String("null"),
            ContentTypes: fastly.String("sint"),
            Extensions: fastly.String("accusantium"),
            Name: fastly.String("test-gzip"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateGzipConfigRequest](../../models/operations/creategzipconfigrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
    res, err := s.Gzip.DeleteGzipConfig(ctx, operations.DeleteGzipConfigRequest{
        GzipName: "test-gzip",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteGzipConfigRequest](../../models/operations/deletegzipconfigrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
    res, err := s.Gzip.GetGzipConfigs(ctx, operations.GetGzipConfigsRequest{
        GzipName: "test-gzip",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetGzipConfigsRequest](../../models/operations/getgzipconfigsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
    res, err := s.Gzip.ListGzipConfigs(ctx, operations.ListGzipConfigsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListGzipConfigsRequest](../../models/operations/listgzipconfigsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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
    res, err := s.Gzip.UpdateGzipConfig(ctx, operations.UpdateGzipConfigRequest{
        Gzip: &shared.Gzip{
            CacheCondition: fastly.String("null"),
            ContentTypes: fastly.String("mollitia"),
            Extensions: fastly.String("reiciendis"),
            Name: fastly.String("test-gzip"),
        },
        GzipName: "test-gzip",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateGzipConfigRequest](../../models/operations/updategzipconfigrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateGzipConfigResponse](../../models/operations/updategzipconfigresponse.md), error**

