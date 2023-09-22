# Director

## Overview

A Director is responsible for balancing requests among a group of Backends. In addition to simply balancing, Directors can be configured to attempt retrying failed requests. Additionally, Directors have a quorum setting which can be used to determine when the Director as a whole is considered "up", in order to prevent "server whack-a-mole" following an outage as servers come back up. Only directors created via the API can be modified via the API. Directors known as "autodirectors" that are created automatically when load balancing groups of servers together cannot be modified or retrieved via the API.

<https://developer.fastly.com/reference/api/load-balancing/directors/director>
### Available Operations

* [CreateDirector](#createdirector) - Create a director
* [DeleteDirector](#deletedirector) - Delete a director
* [GetDirector](#getdirector) - Get a director
* [ListDirectors](#listdirectors) - List directors

## CreateDirector

Create a director for a particular service and version.

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
    res, err := s.Director.CreateDirector(ctx, operations.CreateDirectorRequest{
        Director: &shared.Director{
            Backends: []shared.Backend{
                shared.Backend{
                    Address: fastly.String("21355 Casimir Station"),
                    AutoLoadbalance: fastly.Bool(false),
                    BetweenBytesTimeout: fastly.Int64(369808),
                    ClientCert: fastly.String("alias"),
                    Comment: fastly.String("fugit"),
                    ConnectTimeout: fastly.Int64(677817),
                    FirstByteTimeout: fastly.Int64(569618),
                    Healthcheck: fastly.String("tempora"),
                    Hostname: fastly.String("rotating-relationship.biz"),
                    Ipv4: fastly.String("246.110.63.193"),
                    Ipv6: fastly.String("969e:9a3e:fa77:dfb1:4cd6:6ae3:95ef:b9ba"),
                    KeepaliveTime: fastly.Int64(501324),
                    MaxConn: fastly.Int64(533206),
                    MaxTLSVersion: fastly.String("sapiente"),
                    MinTLSVersion: fastly.String("amet"),
                    Name: fastly.String("test-backend"),
                    OverrideHost: fastly.String("deserunt"),
                    Port: fastly.Int64(394869),
                    RequestCondition: fastly.String("vel"),
                    Shield: fastly.String("natus"),
                    SslCaCert: fastly.String("omnis"),
                    SslCertHostname: fastly.String("molestiae"),
                    SslCheckCert: fastly.Bool(false),
                    SslCiphers: fastly.String("perferendis"),
                    SslClientCert: fastly.String("nihil"),
                    SslClientKey: fastly.String("magnam"),
                    SslHostname: fastly.String("distinctio"),
                    SslSniHostname: fastly.String("id"),
                    UseSsl: fastly.Bool(false),
                    Weight: fastly.Int64(287991),
                },
            },
            Capacity: fastly.Int64(290077),
            Comment: fastly.String("suscipit"),
            Name: fastly.String("test-director"),
            Quorum: fastly.Int64(618016),
            Retries: fastly.Int64(749170),
            Shield: fastly.String("eum"),
            Type: shared.DirectorTypeFour.ToPointer(),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateDirectorRequest](../../models/operations/createdirectorrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateDirectorResponse](../../models/operations/createdirectorresponse.md), error**


## DeleteDirector

Delete the director for a particular service and version.

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
    res, err := s.Director.DeleteDirector(ctx, operations.DeleteDirectorRequest{
        DirectorName: "test-director",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDirector200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteDirectorRequest](../../models/operations/deletedirectorrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteDirectorResponse](../../models/operations/deletedirectorresponse.md), error**


## GetDirector

Get the director for a particular service and version.

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
    res, err := s.Director.GetDirector(ctx, operations.GetDirectorRequest{
        DirectorName: "test-director",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetDirectorRequest](../../models/operations/getdirectorrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetDirectorResponse](../../models/operations/getdirectorresponse.md), error**


## ListDirectors

List the directors for a particular service and version.

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
    res, err := s.Director.ListDirectors(ctx, operations.ListDirectorsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectorResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListDirectorsRequest](../../models/operations/listdirectorsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListDirectorsResponse](../../models/operations/listdirectorsresponse.md), error**

