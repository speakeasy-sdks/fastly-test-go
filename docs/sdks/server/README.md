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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.CreatePoolServerSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Server.CreatePoolServer(ctx, operations.CreatePoolServerRequest{
        PoolID: "2Yd1WfiCBPENLloXfXmlO",
        Server: &shared.Server{
            Address: sdk.String("3983 Jacobi Bypass"),
            Comment: sdk.String("cum"),
            Disabled: sdk.Bool(false),
            MaxConn: sdk.Int64(232627),
            OverrideHost: sdk.String("in"),
            Port: sdk.Int64(348519),
            Weight: sdk.Int64(937285),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreatePoolServerRequest](../../models/operations/createpoolserverrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.CreatePoolServerSecurity](../../models/operations/createpoolserversecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.DeletePoolServerSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Server.DeletePoolServer(ctx, operations.DeletePoolServerRequest{
        PoolID: "2Yd1WfiCBPENLloXfXmlO",
        ServerID: "6kEuoknxiaDBCLiAjKqyXq",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletePoolServer200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeletePoolServerRequest](../../models/operations/deletepoolserverrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.DeletePoolServerSecurity](../../models/operations/deletepoolserversecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.GetPoolServerSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Server.GetPoolServer(ctx, operations.GetPoolServerRequest{
        PoolID: "2Yd1WfiCBPENLloXfXmlO",
        ServerID: "6kEuoknxiaDBCLiAjKqyXq",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetPoolServerRequest](../../models/operations/getpoolserverrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.GetPoolServerSecurity](../../models/operations/getpoolserversecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListPoolServersSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Server.ListPoolServers(ctx, operations.ListPoolServersRequest{
        PoolID: "2Yd1WfiCBPENLloXfXmlO",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListPoolServersRequest](../../models/operations/listpoolserversrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ListPoolServersSecurity](../../models/operations/listpoolserverssecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.UpdatePoolServerSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Server.UpdatePoolServer(ctx, operations.UpdatePoolServerRequest{
        PoolID: "2Yd1WfiCBPENLloXfXmlO",
        Server: &shared.Server{
            Address: sdk.String("293 Predovic Union"),
            Comment: sdk.String("necessitatibus"),
            Disabled: sdk.Bool(false),
            MaxConn: sdk.Int64(296556),
            OverrideHost: sdk.String("sunt"),
            Port: sdk.Int64(992012),
            Weight: sdk.Int64(241545),
        },
        ServerID: "6kEuoknxiaDBCLiAjKqyXq",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdatePoolServerRequest](../../models/operations/updatepoolserverrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.UpdatePoolServerSecurity](../../models/operations/updatepoolserversecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.UpdatePoolServerResponse](../../models/operations/updatepoolserverresponse.md), error**

