# Pool
(*Pool*)

## Overview

A pool is responsible for balancing requests among a group of servers. In addition to balancing, pools can be configured to attempt retrying failed requests. Pools have a quorum setting that can be used to determine when the pool as a whole is considered up, in order to prevent problems following an outage as servers come back up.

<https://developer.fastly.com/reference/api/load-balancing/pools/pool>
### Available Operations

* [CreateServerPool](#createserverpool) - Create a server pool
* [DeleteServerPool](#deleteserverpool) - Delete a server pool
* [GetServerPool](#getserverpool) - Get a server pool
* [ListServerPools](#listserverpools) - List server pools
* [UpdateServerPool](#updateserverpool) - Update a server pool

## CreateServerPool

Creates a pool for a particular service and version.

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
    res, err := s.Pool.CreateServerPool(ctx, operations.CreateServerPoolRequest{
        Pool: &shared.Pool{
            Comment: fastly.String("The Nagasaki Lander is the trademarked name of several series of Nagasaki sport bikes, that started with the 1984 ABC800J"),
            ConnectTimeout: fastly.Int64(811080),
            FirstByteTimeout: fastly.Int64(590807),
            Healthcheck: fastly.String("Producer ROI"),
            MaxConnDefault: fastly.Int64(866009),
            MaxTLSVersion: fastly.Int64(533393),
            MinTLSVersion: fastly.Int64(447912),
            Name: fastly.String("my-pool"),
            OverrideHost: fastly.String("empower Darmstadtium"),
            Quorum: fastly.Int64(662201),
            RequestCondition: fastly.String("null"),
            Shield: fastly.String("Paradigm"),
            TLSCaCert: fastly.String("Hybrid Electronic Nihonium"),
            TLSCertHostname: fastly.String("Dinar spool disintermediate"),
            TLSCheckCert: fastly.Int64(135881),
            TLSCiphers: fastly.String("Cotton useful"),
            TLSClientCert: fastly.String("instructive drive Touring"),
            TLSClientKey: fastly.String("Buckinghamshire hack"),
            TLSSniHostname: fastly.String("harum bifurcated"),
            Type: shared.PoolTypeRandom.ToPointer(),
            UseTLS: shared.PoolUseTLSOne.ToPointer(),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateServerPoolRequest](../../models/operations/createserverpoolrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateServerPoolResponse](../../models/operations/createserverpoolresponse.md), error**


## DeleteServerPool

Deletes a specific pool for a particular service and version.

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
    res, err := s.Pool.DeleteServerPool(ctx, operations.DeleteServerPoolRequest{
        PoolName: "my-pool",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteServerPool200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteServerPoolRequest](../../models/operations/deleteserverpoolrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeleteServerPoolResponse](../../models/operations/deleteserverpoolresponse.md), error**


## GetServerPool

Gets a single pool for a particular service and version.

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
    res, err := s.Pool.GetServerPool(ctx, operations.GetServerPoolRequest{
        PoolName: "my-pool",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetServerPoolRequest](../../models/operations/getserverpoolrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetServerPoolResponse](../../models/operations/getserverpoolresponse.md), error**


## ListServerPools

Lists all pools for a particular service and pool.

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
    res, err := s.Pool.ListServerPools(ctx, operations.ListServerPoolsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PoolResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListServerPoolsRequest](../../models/operations/listserverpoolsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListServerPoolsResponse](../../models/operations/listserverpoolsresponse.md), error**


## UpdateServerPool

Updates a specific pool for a particular service and version.

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
    res, err := s.Pool.UpdateServerPool(ctx, operations.UpdateServerPoolRequest{
        Pool: &shared.Pool{
            Comment: fastly.String("Carbonite web goalkeeper gloves are ergonomically designed to give easy fit"),
            ConnectTimeout: fastly.Int64(870420),
            FirstByteTimeout: fastly.Int64(29969),
            Healthcheck: fastly.String("gold Rancho Hybrid"),
            MaxConnDefault: fastly.Int64(422957),
            MaxTLSVersion: fastly.Int64(357933),
            MinTLSVersion: fastly.Int64(955467),
            Name: fastly.String("my-pool"),
            OverrideHost: fastly.String("Cedi"),
            Quorum: fastly.Int64(770007),
            RequestCondition: fastly.String("null"),
            Shield: fastly.String("experiences Tactics"),
            TLSCaCert: fastly.String("terribly quantifying MTF"),
            TLSCertHostname: fastly.String("services BCEAO requite"),
            TLSCheckCert: fastly.Int64(600539),
            TLSCiphers: fastly.String("male"),
            TLSClientCert: fastly.String("Northwest blue digital"),
            TLSClientKey: fastly.String("navigating"),
            TLSSniHostname: fastly.String("Arab"),
            Type: shared.PoolTypeClient.ToPointer(),
            UseTLS: shared.PoolUseTLSOne.ToPointer(),
        },
        PoolName: "my-pool",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateServerPoolRequest](../../models/operations/updateserverpoolrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateServerPoolResponse](../../models/operations/updateserverpoolresponse.md), error**

