# KvStoreItem
(*KvStoreItem*)

## Overview

An item in an kv store.

<https://developer.fastly.com/reference/api/services/resources/kv-store-item>
### Available Operations

* [DeleteKeyFromStore](#deletekeyfromstore) - Delete kv store item.
* [GetKeys](#getkeys) - List kv store keys.
* [GetValueForKey](#getvalueforkey) - Get the value of an kv store item
* [SetValueForKey](#setvalueforkey) - Insert an item into an kv store

## DeleteKeyFromStore

Delete an item from an kv store

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.KvStoreItem.DeleteKeyFromStore(ctx, operations.DeleteKeyFromStoreRequest{
        KeyName: "South Recumbent yuppify",
        StoreID: "Factors turning Electric",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteKeyFromStoreRequest](../../models/operations/deletekeyfromstorerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.DeleteKeyFromStoreResponse](../../models/operations/deletekeyfromstoreresponse.md), error**


## GetKeys

List the keys of all items within an kv store.

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.KvStoreItem.GetKeys(ctx, operations.GetKeysRequest{
        StoreID: "scoot",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetKeys200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetKeysRequest](../../models/operations/getkeysrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.GetKeysResponse](../../models/operations/getkeysresponse.md), error**


## GetValueForKey

Get the value associated with a key.

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.KvStoreItem.GetValueForKey(ctx, operations.GetValueForKeyRequest{
        KeyName: "rehome lumen Vista",
        StoreID: "turquoise incentivize joule",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetValueForKey200ApplicationOctetStreamByteString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetValueForKeyRequest](../../models/operations/getvalueforkeyrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetValueForKeyResponse](../../models/operations/getvalueforkeyresponse.md), error**


## SetValueForKey

Set a new value for a new or existing key in an kv store.

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.KvStoreItem.SetValueForKey(ctx, operations.SetValueForKeyRequest{
        KeyName: "Cotton Kia",
        StoreID: "Southwest twig harum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SetValueForKey200ApplicationOctetStreamByteString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.SetValueForKeyRequest](../../models/operations/setvalueforkeyrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.SetValueForKeyResponse](../../models/operations/setvalueforkeyresponse.md), error**

