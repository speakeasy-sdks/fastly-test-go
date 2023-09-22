# Backend

## Overview

A backend (also sometimes called an origin server) is a server identified by IP address or hostname, from which Fastly will fetch your content. There can be multiple backends attached to a service, but each backend is specific to one service. By default, the first backend added to a service configuration will be used for all requests (provided it meets any [conditions](/reference/api/vcl-services/condition) attached to it). If multiple backends are defined for a service, the first one that has no attached conditions, or whose condition is satisfied for the current request, will be used, unless that behavior is modified using the `auto_loadbalance` field described below.

<https://developer.fastly.com/reference/api/services/backend>
### Available Operations

* [CreateBackend](#createbackend) - Create a backend
* [DeleteBackend](#deletebackend) - Delete a backend
* [GetBackend](#getbackend) - Describe a backend
* [ListBackends](#listbackends) - List backends
* [UpdateBackend](#updatebackend) - Update a backend

## CreateBackend

Create a backend for a particular service and version.

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
    res, err := s.Backend.CreateBackend(ctx, operations.CreateBackendRequest{
        Backend: &shared.Backend{
            Address: fastly.String("80923 Paxton Spurs"),
            AutoLoadbalance: fastly.Bool(false),
            BetweenBytesTimeout: fastly.Int64(528895),
            ClientCert: fastly.String("iusto"),
            Comment: fastly.String("excepturi"),
            ConnectTimeout: fastly.Int64(392785),
            FirstByteTimeout: fastly.Int64(925597),
            Healthcheck: fastly.String("temporibus"),
            Hostname: fastly.String("bite-sized-favorite.com"),
            Ipv4: fastly.String("165.5.94.213"),
            Ipv6: fastly.String("fc2d:df7c:c78c:a1ba:928f:c816:742c:b739"),
            KeepaliveTime: fastly.Int64(135218),
            MaxConn: fastly.Int64(18789),
            MaxTLSVersion: fastly.String("ad"),
            MinTLSVersion: fastly.String("natus"),
            Name: fastly.String("test-backend"),
            OverrideHost: fastly.String("sed"),
            Port: fastly.Int64(612096),
            RequestCondition: fastly.String("dolor"),
            Shield: fastly.String("natus"),
            SslCaCert: fastly.String("laboriosam"),
            SslCertHostname: fastly.String("hic"),
            SslCheckCert: fastly.Bool(false),
            SslCiphers: fastly.String("saepe"),
            SslClientCert: fastly.String("fuga"),
            SslClientKey: fastly.String("in"),
            SslHostname: fastly.String("corporis"),
            SslSniHostname: fastly.String("iste"),
            UseSsl: fastly.Bool(false),
            Weight: fastly.Int64(437032),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BackendResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateBackendRequest](../../models/operations/createbackendrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateBackendResponse](../../models/operations/createbackendresponse.md), error**


## DeleteBackend

Delete the backend for a particular service and version.

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
    res, err := s.Backend.DeleteBackend(ctx, operations.DeleteBackendRequest{
        BackendName: "test-backend",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteBackend200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteBackendRequest](../../models/operations/deletebackendrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.DeleteBackendResponse](../../models/operations/deletebackendresponse.md), error**


## GetBackend

Get the backend for a particular service and version.

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
    res, err := s.Backend.GetBackend(ctx, operations.GetBackendRequest{
        BackendName: "test-backend",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BackendResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetBackendRequest](../../models/operations/getbackendrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetBackendResponse](../../models/operations/getbackendresponse.md), error**


## ListBackends

List all backends for a particular service and version.

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
    res, err := s.Backend.ListBackends(ctx, operations.ListBackendsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BackendResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListBackendsRequest](../../models/operations/listbackendsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListBackendsResponse](../../models/operations/listbackendsresponse.md), error**


## UpdateBackend

Update the backend for a particular service and version.

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
    res, err := s.Backend.UpdateBackend(ctx, operations.UpdateBackendRequest{
        Backend: &shared.Backend{
            Address: fastly.String("600 Orn Plain"),
            AutoLoadbalance: fastly.Bool(false),
            BetweenBytesTimeout: fastly.Int64(670638),
            ClientCert: fastly.String("dolores"),
            Comment: fastly.String("dolorem"),
            ConnectTimeout: fastly.Int64(358152),
            FirstByteTimeout: fastly.Int64(128926),
            Healthcheck: fastly.String("nobis"),
            Hostname: fastly.String("front-odyssey.info"),
            Ipv4: fastly.String("83.145.9.112"),
            Ipv6: fastly.String("aff1:a3a2:fa94:6773:9251:aa52:c3f5:ad01"),
            KeepaliveTime: fastly.Int64(622846),
            MaxConn: fastly.Int64(837945),
            MaxTLSVersion: fastly.String("laborum"),
            MinTLSVersion: fastly.String("quasi"),
            Name: fastly.String("test-backend"),
            OverrideHost: fastly.String("reiciendis"),
            Port: fastly.Int64(976460),
            RequestCondition: fastly.String("vero"),
            Shield: fastly.String("nihil"),
            SslCaCert: fastly.String("praesentium"),
            SslCertHostname: fastly.String("voluptatibus"),
            SslCheckCert: fastly.Bool(false),
            SslCiphers: fastly.String("ipsa"),
            SslClientCert: fastly.String("omnis"),
            SslClientKey: fastly.String("voluptate"),
            SslHostname: fastly.String("cum"),
            SslSniHostname: fastly.String("perferendis"),
            UseSsl: fastly.Bool(false),
            Weight: fastly.Int64(39187),
        },
        BackendName: "test-backend",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BackendResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateBackendRequest](../../models/operations/updatebackendrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateBackendResponse](../../models/operations/updatebackendresponse.md), error**

