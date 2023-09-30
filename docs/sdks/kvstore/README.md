# KvStore
(*KvStore*)

## Overview

An kv store is a persistent, globally consistent key-value store accessible to Compute@Edge services during request processing.

<https://developer.fastly.com/reference/api/services/resources/kv-store>
### Available Operations

* [CreateStore](#createstore) - Create an kv store.
* [DeleteStore](#deletestore) - Delete an kv store.
* [GetStore](#getstore) - Describe an kv store.
* [GetStores](#getstores) - List kv stores.

## CreateStore

Create a new kv store.

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
    res, err := s.KvStore.CreateStore(ctx, operations.CreateStoreRequest{
        Location: fastly.String("Digitized"),
        Store: &shared.Store{
            Name: fastly.String("Ferry"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StoreResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CreateStoreRequest](../../models/operations/createstorerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateStoreResponse](../../models/operations/createstoreresponse.md), error**


## DeleteStore

An kv store must be empty before it can be deleted.  Deleting an kv store that still contains keys will result in a `409` (Conflict).

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
    res, err := s.KvStore.DeleteStore(ctx, operations.DeleteStoreRequest{
        Force: fastly.Bool(false),
        StoreID: "male",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.DeleteStoreRequest](../../models/operations/deletestorerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.DeleteStoreResponse](../../models/operations/deletestoreresponse.md), error**


## GetStore

Get an kv store by ID.

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
    res, err := s.KvStore.GetStore(ctx, operations.GetStoreRequest{
        StoreID: "SSD Technician East",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StoreResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetStoreRequest](../../models/operations/getstorerequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetStoreResponse](../../models/operations/getstoreresponse.md), error**


## GetStores

Get all stores for a given customer.

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
    res, err := s.KvStore.GetStores(ctx, operations.GetStoresRequest{
        Cursor: fastly.String("card Plastic"),
        Limit: fastly.Int64(829081),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetStores200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetStoresRequest](../../models/operations/getstoresrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetStoresResponse](../../models/operations/getstoresresponse.md), error**

