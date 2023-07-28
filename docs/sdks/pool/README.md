# Pool

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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.CreateServerPoolSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Pool.CreateServerPool(ctx, operations.CreateServerPoolRequest{
        Pool2: &shared.Pool2{
            Comment: sdk.String("sit"),
            ConnectTimeout: sdk.Int64(699575),
            FirstByteTimeout: sdk.Int64(148829),
            Healthcheck: sdk.String("reiciendis"),
            MaxConnDefault: sdk.Int64(131852),
            MaxTLSVersion: sdk.Int64(994401),
            MinTLSVersion: sdk.Int64(707918),
            Name: sdk.String("my-pool"),
            OverrideHost: sdk.String("voluptate"),
            Quorum: sdk.Int64(709072),
            RequestCondition: sdk.String("null"),
            Shield: sdk.String("ab"),
            TLSCaCert: sdk.String("iste"),
            TLSCertHostname: sdk.String("dolore"),
            TLSCheckCert: sdk.Int64(671907),
            TLSCiphers: sdk.String("sed"),
            TLSClientCert: sdk.String("in"),
            TLSClientKey: sdk.String("commodi"),
            TLSSniHostname: sdk.String("quidem"),
            Type: shared.PoolTypeRandom.ToPointer(),
            UseTLS: shared.PoolUseTLSZero.ToPointer(),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateServerPoolRequest](../../models/operations/createserverpoolrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.CreateServerPoolSecurity](../../models/operations/createserverpoolsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.DeleteServerPoolSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Pool.DeleteServerPool(ctx, operations.DeleteServerPoolRequest{
        PoolName: "my-pool",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteServerPool200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteServerPoolRequest](../../models/operations/deleteserverpoolrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.DeleteServerPoolSecurity](../../models/operations/deleteserverpoolsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.GetServerPoolSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Pool.GetServerPool(ctx, operations.GetServerPoolRequest{
        PoolName: "my-pool",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetServerPoolRequest](../../models/operations/getserverpoolrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.GetServerPoolSecurity](../../models/operations/getserverpoolsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListServerPoolsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Pool.ListServerPools(ctx, operations.ListServerPoolsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PoolResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListServerPoolsRequest](../../models/operations/listserverpoolsrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ListServerPoolsSecurity](../../models/operations/listserverpoolssecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.UpdateServerPoolSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Pool.UpdateServerPool(ctx, operations.UpdateServerPoolRequest{
        Pool2: &shared.Pool2{
            Comment: sdk.String("unde"),
            ConnectTimeout: sdk.Int64(100032),
            FirstByteTimeout: sdk.Int64(382808),
            Healthcheck: sdk.String("sapiente"),
            MaxConnDefault: sdk.Int64(895386),
            MaxTLSVersion: sdk.Int64(72434),
            MinTLSVersion: sdk.Int64(967795),
            Name: sdk.String("my-pool"),
            OverrideHost: sdk.String("perferendis"),
            Quorum: sdk.Int64(546885),
            RequestCondition: sdk.String("null"),
            Shield: sdk.String("maiores"),
            TLSCaCert: sdk.String("incidunt"),
            TLSCertHostname: sdk.String("sed"),
            TLSCheckCert: sdk.Int64(592231),
            TLSCiphers: sdk.String("eius"),
            TLSClientCert: sdk.String("necessitatibus"),
            TLSClientKey: sdk.String("ipsum"),
            TLSSniHostname: sdk.String("ea"),
            Type: shared.PoolTypeHash.ToPointer(),
            UseTLS: shared.PoolUseTLSOne.ToPointer(),
        },
        PoolName: "my-pool",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateServerPoolRequest](../../models/operations/updateserverpoolrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.UpdateServerPoolSecurity](../../models/operations/updateserverpoolsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.UpdateServerPoolResponse](../../models/operations/updateserverpoolresponse.md), error**

