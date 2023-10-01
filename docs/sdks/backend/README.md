# Backend
(*Backend*)

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
            Address: fastly.String("8907 Jenkins River"),
            AutoLoadbalance: fastly.Bool(false),
            BetweenBytesTimeout: fastly.Int64(183210),
            ClientCert: fastly.String("impactful"),
            Comment: fastly.String("The Football Is Good For Training And Recreational Purposes"),
            ConnectTimeout: fastly.Int64(98007),
            FirstByteTimeout: fastly.Int64(257494),
            Healthcheck: fastly.String("dolor Celsius"),
            Hostname: fastly.String("partial-palm.com"),
            Ipv4: fastly.String("161.90.29.1"),
            Ipv6: fastly.String("d39b:d79a:5d04:c177:f7f9:6719:86e6:473c"),
            KeepaliveTime: fastly.Int64(885848),
            MaxConn: fastly.Int64(588682),
            MaxTLSVersion: fastly.String("gray Eritrea"),
            MinTLSVersion: fastly.String("withdrawal US Pop"),
            Name: fastly.String("test-backend"),
            OverrideHost: fastly.String("Transexual Southwest Buckinghamshire"),
            Port: fastly.Int64(572529),
            RequestCondition: fastly.String("productivity enhance"),
            Shield: fastly.String("Account"),
            SslCaCert: fastly.String("Officer realistic Bronze"),
            SslCertHostname: fastly.String("Bigender"),
            SslCheckCert: fastly.Bool(false),
            SslCiphers: fastly.String("Checking Alabama Representative"),
            SslClientCert: fastly.String("calculate Funk man"),
            SslClientKey: fastly.String("collaborative Configuration"),
            SslHostname: fastly.String("Southwest holistic"),
            SslSniHostname: fastly.String("Cotton"),
            UseSsl: fastly.Bool(false),
            Weight: fastly.Int64(161737),
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
            Address: fastly.String("64862 Kerluke Track"),
            AutoLoadbalance: fastly.Bool(false),
            BetweenBytesTimeout: fastly.Int64(56573),
            ClientCert: fastly.String("Response politicize Ashburn"),
            Comment: fastly.String("The Nagasaki Lander is the trademarked name of several series of Nagasaki sport bikes, that started with the 1984 ABC800J"),
            ConnectTimeout: fastly.Int64(378358),
            FirstByteTimeout: fastly.Int64(306955),
            Healthcheck: fastly.String("API"),
            Hostname: fastly.String("willing-emergent.name"),
            Ipv4: fastly.String("227.145.93.119"),
            Ipv6: fastly.String("bab4:866d:fe81:1d0a:13db:54fa:628f:0807"),
            KeepaliveTime: fastly.Int64(424116),
            MaxConn: fastly.Int64(391076),
            MaxTLSVersion: fastly.String("man Northwest"),
            MinTLSVersion: fastly.String("pfft watt generate"),
            Name: fastly.String("test-backend"),
            OverrideHost: fastly.String("Chrysler"),
            Port: fastly.Int64(882064),
            RequestCondition: fastly.String("Convertible"),
            Shield: fastly.String("turquoise"),
            SslCaCert: fastly.String("West Security Jeep"),
            SslCertHostname: fastly.String("Strategist"),
            SslCheckCert: fastly.Bool(false),
            SslCiphers: fastly.String("hacksaw supermarket battle"),
            SslClientCert: fastly.String("green Computer"),
            SslClientKey: fastly.String("Wooden Bicycle open"),
            SslHostname: fastly.String("bifurcated"),
            SslSniHostname: fastly.String("unto nor"),
            UseSsl: fastly.Bool(false),
            Weight: fastly.Int64(67164),
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

