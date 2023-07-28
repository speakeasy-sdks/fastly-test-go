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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.CreateBackendSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Backend.CreateBackend(ctx, operations.CreateBackendRequest{
        Backend: &shared.Backend{
            Address: sdk.String("3980 Ashlee Place"),
            AutoLoadbalance: sdk.Bool(false),
            BetweenBytesTimeout: sdk.Int64(20218),
            ClientCert: sdk.String("ipsam"),
            Comment: sdk.String("repellendus"),
            ConnectTimeout: sdk.Int64(957156),
            FirstByteTimeout: sdk.Int64(778157),
            Healthcheck: sdk.String("odit"),
            Hostname: sdk.String("trifling-sunday.org"),
            Ipv4: sdk.String("121.204.205.118"),
            Ipv6: sdk.String("8ca1:ba92:8fc8:1674:2cb7:3920:5929:396f"),
            KeepaliveTime: sdk.Int64(902599),
            MaxConn: sdk.Int64(681820),
            MaxTLSVersion: sdk.String("in"),
            MinTLSVersion: sdk.String("corporis"),
            Name: sdk.String("test-backend"),
            OverrideHost: sdk.String("iste"),
            Port: sdk.Int64(437032),
            RequestCondition: sdk.String("saepe"),
            Shield: sdk.String("quidem"),
            SslCaCert: sdk.String("architecto"),
            SslCertHostname: sdk.String("ipsa"),
            SslCheckCert: sdk.Bool(false),
            SslCiphers: sdk.String("reiciendis"),
            SslClientCert: sdk.String("est"),
            SslClientKey: sdk.String("mollitia"),
            SslHostname: sdk.String("laborum"),
            SslSniHostname: sdk.String("dolores"),
            UseSsl: sdk.Bool(false),
            Weight: sdk.Int64(210382),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.BackendResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateBackendRequest](../../models/operations/createbackendrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.CreateBackendSecurity](../../models/operations/createbackendsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.DeleteBackendSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Backend.DeleteBackend(ctx, operations.DeleteBackendRequest{
        BackendName: "test-backend",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteBackend200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteBackendRequest](../../models/operations/deletebackendrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.DeleteBackendSecurity](../../models/operations/deletebackendsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.GetBackendSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Backend.GetBackend(ctx, operations.GetBackendRequest{
        BackendName: "test-backend",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.BackendResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetBackendRequest](../../models/operations/getbackendrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.GetBackendSecurity](../../models/operations/getbackendsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListBackendsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Backend.ListBackends(ctx, operations.ListBackendsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.BackendResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListBackendsRequest](../../models/operations/listbackendsrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.ListBackendsSecurity](../../models/operations/listbackendssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.UpdateBackendSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Backend.UpdateBackend(ctx, operations.UpdateBackendRequest{
        Backend: &shared.Backend{
            Address: sdk.String("1736 Era Mountains"),
            AutoLoadbalance: sdk.Bool(false),
            BetweenBytesTimeout: sdk.Int64(38425),
            ClientCert: sdk.String("iure"),
            Comment: sdk.String("culpa"),
            ConnectTimeout: sdk.Int64(988374),
            FirstByteTimeout: sdk.Int64(958950),
            Healthcheck: sdk.String("architecto"),
            Hostname: sdk.String("prize-cornmeal.name"),
            Ipv4: sdk.String("41.254.167.148"),
            Ipv6: sdk.String("4677:3925:1aa5:2c3f:5ad0:19da:1ffe:78f0"),
            KeepaliveTime: sdk.Int64(604846),
            MaxConn: sdk.Int64(451159),
            MaxTLSVersion: sdk.String("cum"),
            MinTLSVersion: sdk.String("perferendis"),
            Name: sdk.String("test-backend"),
            OverrideHost: sdk.String("doloremque"),
            Port: sdk.Int64(441711),
            RequestCondition: sdk.String("ut"),
            Shield: sdk.String("maiores"),
            SslCaCert: sdk.String("dicta"),
            SslCertHostname: sdk.String("corporis"),
            SslCheckCert: sdk.Bool(false),
            SslCiphers: sdk.String("dolore"),
            SslClientCert: sdk.String("iusto"),
            SslClientKey: sdk.String("dicta"),
            SslHostname: sdk.String("harum"),
            SslSniHostname: sdk.String("enim"),
            UseSsl: sdk.Bool(false),
            Weight: sdk.Int64(880476),
        },
        BackendName: "test-backend",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.BackendResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateBackendRequest](../../models/operations/updatebackendrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.UpdateBackendSecurity](../../models/operations/updatebackendsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.UpdateBackendResponse](../../models/operations/updatebackendresponse.md), error**

