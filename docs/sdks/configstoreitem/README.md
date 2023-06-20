# ConfigStoreItem

## Overview

A key-value pair within a config store.

<https://developer.fastly.com/reference/api/services/resources/config-store-item>
### Available Operations

* [BulkUpdateConfigStoreItem](#bulkupdateconfigstoreitem) - Update multiple entries in a config store
* [CreateConfigStoreItem](#createconfigstoreitem) - Create an entry in a config store
* [DeleteConfigStoreItem](#deleteconfigstoreitem) - Delete an item from a config store
* [GetConfigStoreItem](#getconfigstoreitem) - Get an item from a config store
* [ListConfigStoreItems](#listconfigstoreitems) - List items in a config store
* [UpdateConfigStoreItem](#updateconfigstoreitem) - Update an entry in a config store
* [UpsertConfigStoreItem](#upsertconfigstoreitem) - Insert or update an entry in a config store

## BulkUpdateConfigStoreItem

Add multiple key-value pairs to an individual config store, specified by ID.

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
    res, err := s.ConfigStoreItem.BulkUpdateConfigStoreItem(ctx, operations.BulkUpdateConfigStoreItemRequest{
        BulkUpdateConfigStoreListRequest: &shared.BulkUpdateConfigStoreListRequest{
            Items: []shared.BulkUpdateConfigStoreItem{
                shared.BulkUpdateConfigStoreItem{
                    ItemKey: sdk.String("test-key"),
                    ItemValue: sdk.String("test-value"),
                    Op: shared.BulkUpdateConfigStoreItemOpCreate.ToPointer(),
                },
                shared.BulkUpdateConfigStoreItem{
                    ItemKey: sdk.String("test-key"),
                    ItemValue: sdk.String("test-value"),
                    Op: shared.BulkUpdateConfigStoreItemOpUpsert.ToPointer(),
                },
                shared.BulkUpdateConfigStoreItem{
                    ItemKey: sdk.String("test-key"),
                    ItemValue: sdk.String("test-value"),
                    Op: shared.BulkUpdateConfigStoreItemOpUpdate.ToPointer(),
                },
            },
        },
        ConfigStoreID: "7Lsb7Y76rChV9hSrv3KgFl",
    }, operations.BulkUpdateConfigStoreItemSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BulkUpdateConfigStoreItem200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.BulkUpdateConfigStoreItemRequest](../../models/operations/bulkupdateconfigstoreitemrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.BulkUpdateConfigStoreItemSecurity](../../models/operations/bulkupdateconfigstoreitemsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.BulkUpdateConfigStoreItemResponse](../../models/operations/bulkupdateconfigstoreitemresponse.md), error**


## CreateConfigStoreItem

Add a single key-value pair to an individual config store, specified by ID.

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
    res, err := s.ConfigStoreItem.CreateConfigStoreItem(ctx, operations.CreateConfigStoreItemRequest{
        ConfigStoreID: "7Lsb7Y76rChV9hSrv3KgFl",
        ConfigStoreItem: &shared.ConfigStoreItem{
            ItemKey: sdk.String("test-key"),
            ItemValue: sdk.String("test-value"),
        },
    }, operations.CreateConfigStoreItemSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigStoreItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateConfigStoreItemRequest](../../models/operations/createconfigstoreitemrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.CreateConfigStoreItemSecurity](../../models/operations/createconfigstoreitemsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.CreateConfigStoreItemResponse](../../models/operations/createconfigstoreitemresponse.md), error**


## DeleteConfigStoreItem

Delete an entry in a config store given a config store ID, and item key.

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
    res, err := s.ConfigStoreItem.DeleteConfigStoreItem(ctx, operations.DeleteConfigStoreItemRequest{
        ConfigStoreID: "7Lsb7Y76rChV9hSrv3KgFl",
        ConfigStoreItemKey: "test-key",
    }, operations.DeleteConfigStoreItemSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteConfigStoreItem200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteConfigStoreItemRequest](../../models/operations/deleteconfigstoreitemrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.DeleteConfigStoreItemSecurity](../../models/operations/deleteconfigstoreitemsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.DeleteConfigStoreItemResponse](../../models/operations/deleteconfigstoreitemresponse.md), error**


## GetConfigStoreItem

Retrieve a config store entry given a config store ID and item key.

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
    res, err := s.ConfigStoreItem.GetConfigStoreItem(ctx, operations.GetConfigStoreItemRequest{
        ConfigStoreID: "7Lsb7Y76rChV9hSrv3KgFl",
        ConfigStoreItemKey: "test-key",
    }, operations.GetConfigStoreItemSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigStoreItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetConfigStoreItemRequest](../../models/operations/getconfigstoreitemrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.GetConfigStoreItemSecurity](../../models/operations/getconfigstoreitemsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.GetConfigStoreItemResponse](../../models/operations/getconfigstoreitemresponse.md), error**


## ListConfigStoreItems

List the key-value pairs associated with a given config store ID.

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
    res, err := s.ConfigStoreItem.ListConfigStoreItems(ctx, operations.ListConfigStoreItemsRequest{
        ConfigStoreID: "7Lsb7Y76rChV9hSrv3KgFl",
    }, operations.ListConfigStoreItemsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigStoreItemResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListConfigStoreItemsRequest](../../models/operations/listconfigstoreitemsrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.ListConfigStoreItemsSecurity](../../models/operations/listconfigstoreitemssecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.ListConfigStoreItemsResponse](../../models/operations/listconfigstoreitemsresponse.md), error**


## UpdateConfigStoreItem

Update an entry in a config store given a config store ID, item key, and item value.

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
    res, err := s.ConfigStoreItem.UpdateConfigStoreItem(ctx, operations.UpdateConfigStoreItemRequest{
        ConfigStoreID: "7Lsb7Y76rChV9hSrv3KgFl",
        ConfigStoreItem: &shared.ConfigStoreItem{
            ItemKey: sdk.String("test-key"),
            ItemValue: sdk.String("test-value"),
        },
        ConfigStoreItemKey: "test-key",
    }, operations.UpdateConfigStoreItemSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigStoreItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateConfigStoreItemRequest](../../models/operations/updateconfigstoreitemrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.UpdateConfigStoreItemSecurity](../../models/operations/updateconfigstoreitemsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.UpdateConfigStoreItemResponse](../../models/operations/updateconfigstoreitemresponse.md), error**


## UpsertConfigStoreItem

Insert or update an entry in a config store given a config store ID, item key, and item value.

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
    res, err := s.ConfigStoreItem.UpsertConfigStoreItem(ctx, operations.UpsertConfigStoreItemRequest{
        ConfigStoreID: "7Lsb7Y76rChV9hSrv3KgFl",
        ConfigStoreItem: &shared.ConfigStoreItem{
            ItemKey: sdk.String("test-key"),
            ItemValue: sdk.String("test-value"),
        },
        ConfigStoreItemKey: "test-key",
    }, operations.UpsertConfigStoreItemSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigStoreItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpsertConfigStoreItemRequest](../../models/operations/upsertconfigstoreitemrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.UpsertConfigStoreItemSecurity](../../models/operations/upsertconfigstoreitemsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.UpsertConfigStoreItemResponse](../../models/operations/upsertconfigstoreitemresponse.md), error**

