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
    operationSecurity := operations.CreateDirectorSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Director.CreateDirector(ctx, operations.CreateDirectorRequest{
        Director: &shared.Director{
            Backends: []shared.Backend{
                shared.Backend{
                    Address: sdk.String("822 Grant Oval"),
                    AutoLoadbalance: sdk.Bool(false),
                    BetweenBytesTimeout: sdk.Int64(552822),
                    ClientCert: sdk.String("perferendis"),
                    Comment: sdk.String("magni"),
                    ConnectTimeout: sdk.Int64(828940),
                    FirstByteTimeout: sdk.Int64(369808),
                    Healthcheck: sdk.String("alias"),
                    Hostname: sdk.String("costly-poncho.name"),
                    Ipv4: sdk.String("69.180.188.73"),
                    Ipv6: sdk.String("f63c:969e:9a3e:fa77:dfb1:4cd6:6ae3:95ef"),
                    KeepaliveTime: sdk.Int64(692532),
                    MaxConn: sdk.Int64(588465),
                    MaxTLSVersion: sdk.String("nam"),
                    MinTLSVersion: sdk.String("id"),
                    Name: sdk.String("test-backend"),
                    OverrideHost: sdk.String("blanditiis"),
                    Port: sdk.Int64(533206),
                    RequestCondition: sdk.String("sapiente"),
                    Shield: sdk.String("amet"),
                    SslCaCert: sdk.String("deserunt"),
                    SslCertHostname: sdk.String("nisi"),
                    SslCheckCert: sdk.Bool(false),
                    SslCiphers: sdk.String("vel"),
                    SslClientCert: sdk.String("natus"),
                    SslClientKey: sdk.String("omnis"),
                    SslHostname: sdk.String("molestiae"),
                    SslSniHostname: sdk.String("perferendis"),
                    UseSsl: sdk.Bool(false),
                    Weight: sdk.Int64(470132),
                },
            },
            Capacity: sdk.Int64(301575),
            Comment: sdk.String("distinctio"),
            Name: sdk.String("test-director"),
            Quorum: sdk.Int64(660174),
            Retries: sdk.Int64(287991),
            Shield: sdk.String("labore"),
            Type: shared.DirectorTypeThree.ToPointer(),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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
    operationSecurity := operations.DeleteDirectorSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Director.DeleteDirector(ctx, operations.DeleteDirectorRequest{
        DirectorName: "test-director",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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
    operationSecurity := operations.GetDirectorSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Director.GetDirector(ctx, operations.GetDirectorRequest{
        DirectorName: "test-director",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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
    operationSecurity := operations.ListDirectorsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Director.ListDirectors(ctx, operations.ListDirectorsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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

