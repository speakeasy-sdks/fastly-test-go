# Server

## Overview

A server is an address (IP address or hostname) to which the Fastly Load Balancer service can forward requests. This service can define multiple servers and assign it to a pool. Fastly can then select any one of these servers based on a selection policy defined for the pool.

<https://developer.fastly.com/reference/api/load-balancing/pools/server>
### Available Operations

* [CreatePoolServer](#createpoolserver) - Add a server to a pool
* [DeletePoolServer](#deletepoolserver) - Delete a server from a pool
* [GetPoolServer](#getpoolserver) - Get a pool server
* [ListPoolServers](#listpoolservers) - List servers in a pool
* [UpdatePoolServer](#updatepoolserver) - Update a server

## CreatePoolServer

Creates a single server for a particular service and pool.

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
    res, err := s.Server.CreatePoolServer(ctx, operations.CreatePoolServerRequest{
        PoolID: "2Yd1WfiCBPENLloXfXmlO",
        Server: &shared.Server{
            Address: fastly.String("829 Vivienne Radial"),
            Comment: fastly.String("saepe"),
            Disabled: fastly.Bool(false),
            MaxConn: fastly.Int64(897071),
            OverrideHost: fastly.String("dolore"),
            Port: fastly.Int64(121059),
            Weight: fastly.Int64(992012),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreatePoolServerRequest](../../models/operations/createpoolserverrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreatePoolServerResponse](../../models/operations/createpoolserverresponse.md), error**


## DeletePoolServer

Deletes a single server for a particular service and pool.

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
    res, err := s.Server.DeletePoolServer(ctx, operations.DeletePoolServerRequest{
        PoolID: "2Yd1WfiCBPENLloXfXmlO",
        ServerID: "6kEuoknxiaDBCLiAjKqyXq",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletePoolServer200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeletePoolServerRequest](../../models/operations/deletepoolserverrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeletePoolServerResponse](../../models/operations/deletepoolserverresponse.md), error**


## GetPoolServer

Gets a single server for a particular service and pool.

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
    res, err := s.Server.GetPoolServer(ctx, operations.GetPoolServerRequest{
        PoolID: "2Yd1WfiCBPENLloXfXmlO",
        ServerID: "6kEuoknxiaDBCLiAjKqyXq",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetPoolServerRequest](../../models/operations/getpoolserverrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetPoolServerResponse](../../models/operations/getpoolserverresponse.md), error**


## ListPoolServers

Lists all servers for a particular service and pool.

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
    res, err := s.Server.ListPoolServers(ctx, operations.ListPoolServersRequest{
        PoolID: "2Yd1WfiCBPENLloXfXmlO",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListPoolServersRequest](../../models/operations/listpoolserversrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListPoolServersResponse](../../models/operations/listpoolserversresponse.md), error**


## UpdatePoolServer

Updates a single server for a particular service and pool.

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
    res, err := s.Server.UpdatePoolServer(ctx, operations.UpdatePoolServerRequest{
        PoolID: "2Yd1WfiCBPENLloXfXmlO",
        Server: &shared.Server{
            Address: fastly.String("22149 Feest Harbor"),
            Comment: fastly.String("harum"),
            Disabled: fastly.Bool(false),
            MaxConn: fastly.Int64(385237),
            OverrideHost: fastly.String("ipsa"),
            Port: fastly.Int64(916727),
            Weight: fastly.Int64(730709),
        },
        ServerID: "6kEuoknxiaDBCLiAjKqyXq",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdatePoolServerRequest](../../models/operations/updatepoolserverrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdatePoolServerResponse](../../models/operations/updatepoolserverresponse.md), error**

