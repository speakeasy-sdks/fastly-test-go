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
            Address: sdk.String("80923 Paxton Spurs"),
            AutoLoadbalance: sdk.Bool(false),
            BetweenBytesTimeout: sdk.Int64(528895),
            ClientCert: sdk.String("iusto"),
            Comment: sdk.String("excepturi"),
            ConnectTimeout: sdk.Int64(392785),
            FirstByteTimeout: sdk.Int64(925597),
            Healthcheck: sdk.String("temporibus"),
            Hostname: sdk.String("bite-sized-favorite.com"),
            Ipv4: sdk.String("165.5.94.213"),
            Ipv6: sdk.String("fc2d:df7c:c78c:a1ba:928f:c816:742c:b739"),
            KeepaliveTime: sdk.Int64(135218),
            MaxConn: sdk.Int64(18789),
            MaxTLSVersion: sdk.String("ad"),
            MinTLSVersion: sdk.String("natus"),
            Name: sdk.String("test-backend"),
            OverrideHost: sdk.String("sed"),
            Port: sdk.Int64(612096),
            RequestCondition: sdk.String("dolor"),
            Shield: sdk.String("natus"),
            SslCaCert: sdk.String("laboriosam"),
            SslCertHostname: sdk.String("hic"),
            SslCheckCert: sdk.Bool(false),
            SslCiphers: sdk.String("saepe"),
            SslClientCert: sdk.String("fuga"),
            SslClientKey: sdk.String("in"),
            SslHostname: sdk.String("corporis"),
            SslSniHostname: sdk.String("iste"),
            UseSsl: sdk.Bool(false),
            Weight: sdk.Int64(437032),
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
            Address: sdk.String("600 Orn Plain"),
            AutoLoadbalance: sdk.Bool(false),
            BetweenBytesTimeout: sdk.Int64(670638),
            ClientCert: sdk.String("dolores"),
            Comment: sdk.String("dolorem"),
            ConnectTimeout: sdk.Int64(358152),
            FirstByteTimeout: sdk.Int64(128926),
            Healthcheck: sdk.String("nobis"),
            Hostname: sdk.String("front-odyssey.info"),
            Ipv4: sdk.String("83.145.9.112"),
            Ipv6: sdk.String("aff1:a3a2:fa94:6773:9251:aa52:c3f5:ad01"),
            KeepaliveTime: sdk.Int64(622846),
            MaxConn: sdk.Int64(837945),
            MaxTLSVersion: sdk.String("laborum"),
            MinTLSVersion: sdk.String("quasi"),
            Name: sdk.String("test-backend"),
            OverrideHost: sdk.String("reiciendis"),
            Port: sdk.Int64(976460),
            RequestCondition: sdk.String("vero"),
            Shield: sdk.String("nihil"),
            SslCaCert: sdk.String("praesentium"),
            SslCertHostname: sdk.String("voluptatibus"),
            SslCheckCert: sdk.Bool(false),
            SslCiphers: sdk.String("ipsa"),
            SslClientCert: sdk.String("omnis"),
            SslClientKey: sdk.String("voluptate"),
            SslHostname: sdk.String("cum"),
            SslSniHostname: sdk.String("perferendis"),
            UseSsl: sdk.Bool(false),
            Weight: sdk.Int64(39187),
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

