# Http3

## Overview

Supports the use of the HTTP/3 (QUIC) protocol.

<https://developer.fastly.com/reference/api/vcl-services/http3>
### Available Operations

* [CreateHttp3](#createhttp3) - Enable support for HTTP/3
* [DeleteHttp3](#deletehttp3) - Disable support for HTTP/3
* [GetHttp3](#gethttp3) - Get HTTP/3 status

## CreateHttp3

Enable HTTP/3 (QUIC) support for a particular service and version.

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
    res, err := s.Http3.CreateHttp3(ctx, operations.CreateHttp3Request{
        Http3Input: &shared.Http3Input{
            FeatureRevision: sdk.Int64(396060),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateHttp3Security{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Http3 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateHttp3Request](../../models/operations/createhttp3request.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.CreateHttp3Security](../../models/operations/createhttp3security.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.CreateHttp3Response](../../models/operations/createhttp3response.md), error**


## DeleteHttp3

Disable HTTP/3 (QUIC) support for a particular service and version.

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
    res, err := s.Http3.DeleteHttp3(ctx, operations.DeleteHttp3Request{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteHttp3Security{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteHttp3200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteHttp3Request](../../models/operations/deletehttp3request.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.DeleteHttp3Security](../../models/operations/deletehttp3security.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.DeleteHttp3Response](../../models/operations/deletehttp3response.md), error**


## GetHttp3

Get the status of HTTP/3 (QUIC) support for a particular service and version.

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
    res, err := s.Http3.GetHttp3(ctx, operations.GetHttp3Request{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetHttp3Security{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Http3 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetHttp3Request](../../models/operations/gethttp3request.md)   | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.GetHttp3Security](../../models/operations/gethttp3security.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.GetHttp3Response](../../models/operations/gethttp3response.md), error**

