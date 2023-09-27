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
        Pool2: &shared.Pool2{
            Comment: fastly.String("ex"),
            ConnectTimeout: fastly.Int64(536275),
            FirstByteTimeout: fastly.Int64(929292),
            Healthcheck: fastly.String("dolorum"),
            MaxConnDefault: fastly.Int64(99615),
            MaxTLSVersion: fastly.Int64(609178),
            MinTLSVersion: fastly.Int64(945302),
            Name: fastly.String("my-pool"),
            OverrideHost: fastly.String("quasi"),
            Quorum: fastly.Int64(869489),
            RequestCondition: fastly.String("null"),
            Shield: fastly.String("et"),
            TLSCaCert: fastly.String("voluptate"),
            TLSCertHostname: fastly.String("ipsa"),
            TLSCheckCert: fastly.Int64(326701),
            TLSCiphers: fastly.String("veritatis"),
            TLSClientCert: fastly.String("consectetur"),
            TLSClientKey: fastly.String("adipisci"),
            TLSSniHostname: fastly.String("iste"),
            Type: shared.PoolTypeClient.ToPointer(),
            UseTLS: shared.PoolUseTLSZero.ToPointer(),
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
        Pool2: &shared.Pool2{
            Comment: fastly.String("rem"),
            ConnectTimeout: fastly.Int64(15606),
            FirstByteTimeout: fastly.Int64(513075),
            Healthcheck: fastly.String("eum"),
            MaxConnDefault: fastly.Int64(649832),
            MaxTLSVersion: fastly.Int64(68074),
            MinTLSVersion: fastly.Int64(544591),
            Name: fastly.String("my-pool"),
            OverrideHost: fastly.String("non"),
            Quorum: fastly.Int64(32465),
            RequestCondition: fastly.String("null"),
            Shield: fastly.String("dolor"),
            TLSCaCert: fastly.String("occaecati"),
            TLSCertHostname: fastly.String("numquam"),
            TLSCheckCert: fastly.Int64(771089),
            TLSCiphers: fastly.String("explicabo"),
            TLSClientCert: fastly.String("voluptas"),
            TLSClientKey: fastly.String("aut"),
            TLSSniHostname: fastly.String("dignissimos"),
            Type: shared.PoolTypeRandom.ToPointer(),
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

