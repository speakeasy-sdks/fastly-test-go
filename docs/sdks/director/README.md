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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Director.CreateDirector(ctx, operations.CreateDirectorRequest{
        Director: &shared.Director{
            Backends: []shared.Backend{
                shared.Backend{
                    Address: sdk.String("77294 Nigel Neck"),
                    AutoLoadbalance: sdk.Bool(false),
                    BetweenBytesTimeout: sdk.Int64(396098),
                    ClientCert: sdk.String("provident"),
                    Comment: sdk.String("necessitatibus"),
                    ConnectTimeout: sdk.Int64(572252),
                    FirstByteTimeout: sdk.Int64(638921),
                    Healthcheck: sdk.String("dolor"),
                    Hostname: sdk.String("unimportant-venture.net"),
                    Ipv4: sdk.String("114.114.216.250"),
                    Ipv6: sdk.String("b14c:d66a:e395:efb9:ba88:f3a6:6997:074b"),
                    KeepaliveTime: sdk.Int64(660174),
                    MaxConn: sdk.Int64(287991),
                    MaxTLSVersion: sdk.String("labore"),
                    MinTLSVersion: sdk.String("suscipit"),
                    Name: sdk.String("test-backend"),
                    OverrideHost: sdk.String("natus"),
                    Port: sdk.Int64(749170),
                    RequestCondition: sdk.String("eum"),
                    Shield: sdk.String("vero"),
                    SslCaCert: sdk.String("aspernatur"),
                    SslCertHostname: sdk.String("architecto"),
                    SslCheckCert: sdk.Bool(false),
                    SslCiphers: sdk.String("magnam"),
                    SslClientCert: sdk.String("et"),
                    SslClientKey: sdk.String("excepturi"),
                    SslHostname: sdk.String("ullam"),
                    SslSniHostname: sdk.String("provident"),
                    UseSsl: sdk.Bool(false),
                    Weight: sdk.Int64(551816),
                },
                shared.Backend{
                    Address: sdk.String("0696 Holden Extensions"),
                    AutoLoadbalance: sdk.Bool(false),
                    BetweenBytesTimeout: sdk.Int64(896547),
                    ClientCert: sdk.String("odit"),
                    Comment: sdk.String("nemo"),
                    ConnectTimeout: sdk.Int64(97260),
                    FirstByteTimeout: sdk.Int64(435865),
                    Healthcheck: sdk.String("doloribus"),
                    Hostname: sdk.String("unique-diesel.net"),
                    Ipv4: sdk.String("137.180.114.25"),
                    Ipv6: sdk.String("1e5b:7fd2:ed02:8921:cddc:6926:01fb:576b"),
                    KeepaliveTime: sdk.Int64(50588),
                    MaxConn: sdk.Int64(866383),
                    MaxTLSVersion: sdk.String("nemo"),
                    MinTLSVersion: sdk.String("voluptatibus"),
                    Name: sdk.String("test-backend"),
                    OverrideHost: sdk.String("perferendis"),
                    Port: sdk.Int64(855804),
                    RequestCondition: sdk.String("amet"),
                    Shield: sdk.String("aut"),
                    SslCaCert: sdk.String("cumque"),
                    SslCertHostname: sdk.String("corporis"),
                    SslCheckCert: sdk.Bool(false),
                    SslCiphers: sdk.String("hic"),
                    SslClientCert: sdk.String("libero"),
                    SslClientKey: sdk.String("nobis"),
                    SslHostname: sdk.String("dolores"),
                    SslSniHostname: sdk.String("quis"),
                    UseSsl: sdk.Bool(false),
                    Weight: sdk.Int64(521037),
                },
                shared.Backend{
                    Address: sdk.String("0311 Cecilia Skyway"),
                    AutoLoadbalance: sdk.Bool(false),
                    BetweenBytesTimeout: sdk.Int64(463451),
                    ClientCert: sdk.String("dolor"),
                    Comment: sdk.String("vero"),
                    ConnectTimeout: sdk.Int64(345352),
                    FirstByteTimeout: sdk.Int64(944120),
                    Healthcheck: sdk.String("recusandae"),
                    Hostname: sdk.String("pale-psychoanalyst.name"),
                    Ipv4: sdk.String("8.200.42.128"),
                    Ipv6: sdk.String("909b:3fe4:9a8d:9cbf:4863:3323:f9b7:7f3a"),
                    KeepaliveTime: sdk.Int64(254356),
                    MaxConn: sdk.Int64(85295),
                    MaxTLSVersion: sdk.String("ipsa"),
                    MinTLSVersion: sdk.String("ipsa"),
                    Name: sdk.String("test-backend"),
                    OverrideHost: sdk.String("iure"),
                    Port: sdk.Int64(487838),
                    RequestCondition: sdk.String("quaerat"),
                    Shield: sdk.String("accusamus"),
                    SslCaCert: sdk.String("quidem"),
                    SslCertHostname: sdk.String("voluptatibus"),
                    SslCheckCert: sdk.Bool(false),
                    SslCiphers: sdk.String("voluptas"),
                    SslClientCert: sdk.String("natus"),
                    SslClientKey: sdk.String("eos"),
                    SslHostname: sdk.String("atque"),
                    SslSniHostname: sdk.String("sit"),
                    UseSsl: sdk.Bool(false),
                    Weight: sdk.Int64(854614),
                },
            },
            Capacity: sdk.Int64(67249),
            Comment: sdk.String("soluta"),
            Name: sdk.String("test-director"),
            Quorum: sdk.Int64(679393),
            Retries: sdk.Int64(478596),
            Shield: sdk.String("voluptate"),
            Type: shared.DirectorTypeFour.ToPointer(),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateDirectorSecurity{
        Token: "",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateDirectorRequest](../../models/operations/createdirectorrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.CreateDirectorSecurity](../../models/operations/createdirectorsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Director.DeleteDirector(ctx, operations.DeleteDirectorRequest{
        DirectorName: "test-director",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteDirectorSecurity{
        Token: "",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteDirectorRequest](../../models/operations/deletedirectorrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.DeleteDirectorSecurity](../../models/operations/deletedirectorsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Director.GetDirector(ctx, operations.GetDirectorRequest{
        DirectorName: "test-director",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetDirectorSecurity{
        Token: "",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetDirectorRequest](../../models/operations/getdirectorrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.GetDirectorSecurity](../../models/operations/getdirectorsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Director.ListDirectors(ctx, operations.ListDirectorsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListDirectorsSecurity{
        Token: "",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListDirectorsRequest](../../models/operations/listdirectorsrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListDirectorsSecurity](../../models/operations/listdirectorssecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.ListDirectorsResponse](../../models/operations/listdirectorsresponse.md), error**

