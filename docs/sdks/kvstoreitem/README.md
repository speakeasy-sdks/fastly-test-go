# KvStoreItem

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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.KvStoreItem.DeleteKeyFromStore(ctx, operations.DeleteKeyFromStoreRequest{
        Force: sdk.Bool(false),
        KeyName: "veritatis",
        StoreID: "nobis",
    }, operations.DeleteKeyFromStoreSecurity{
        Token: "",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteKeyFromStoreRequest](../../models/operations/deletekeyfromstorerequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.DeleteKeyFromStoreSecurity](../../models/operations/deletekeyfromstoresecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.KvStoreItem.GetKeys(ctx, operations.GetKeysRequest{
        Cursor: sdk.String("quos"),
        Limit: sdk.Int64(731694),
        Prefix: sdk.String("cupiditate"),
        StoreID: "aperiam",
    }, operations.GetKeysSecurity{
        Token: "",
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetKeysRequest](../../models/operations/getkeysrequest.md)   | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `security`                                                               | [operations.GetKeysSecurity](../../models/operations/getkeyssecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.KvStoreItem.GetValueForKey(ctx, operations.GetValueForKeyRequest{
        KeyName: "delectus",
        StoreID: "dolorem",
    }, operations.GetValueForKeySecurity{
        Token: "",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetValueForKeyRequest](../../models/operations/getvalueforkeyrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetValueForKeySecurity](../../models/operations/getvalueforkeysecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.KvStoreItem.SetValueForKey(ctx, operations.SetValueForKeyRequest{
        RequestBody: sdk.String("dolore"),
        Add: sdk.Bool(false),
        Append: sdk.Bool(false),
        BackgroundFetch: sdk.Bool(false),
        IfGenerationMatch: sdk.Int64(286915),
        KeyName: "adipisci",
        Metadata: sdk.String("dolorum"),
        Prepend: sdk.Bool(false),
        StoreID: "architecto",
        TimeToLiveSec: sdk.Int64(63038),
    }, operations.SetValueForKeySecurity{
        Token: "",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.SetValueForKeyRequest](../../models/operations/setvalueforkeyrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.SetValueForKeySecurity](../../models/operations/setvalueforkeysecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.SetValueForKeyResponse](../../models/operations/setvalueforkeyresponse.md), error**
