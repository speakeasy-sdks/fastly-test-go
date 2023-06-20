# DictionaryItem

## Overview

A Dictionary Item is a single key-value pair that makes up an entry in a Dictionary. Dictionary Items can be added, removed and modified without activating a new version of the associated service.

<https://developer.fastly.com/reference/api/dictionaries/dictionary-item>
### Available Operations

* [BulkUpdateDictionaryItem](#bulkupdatedictionaryitem) - Update multiple entries in an edge dictionary
* [CreateDictionaryItem](#createdictionaryitem) - Create an entry in an edge dictionary
* [DeleteDictionaryItem](#deletedictionaryitem) - Delete an item from an edge dictionary
* [GetDictionaryItem](#getdictionaryitem) - Get an item from an edge dictionary
* [ListDictionaryItems](#listdictionaryitems) - List items in an edge dictionary
* [UpdateDictionaryItem](#updatedictionaryitem) - Update an entry in an edge dictionary
* [UpsertDictionaryItem](#upsertdictionaryitem) - Insert or update an entry in an edge dictionary

## BulkUpdateDictionaryItem

Update multiple items in the same dictionary. For faster updates to your service, group your changes into large batches. The maximum batch size is 1000 items. [Contact support](https://support.fastly.com/) to discuss raising this limit.

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
    res, err := s.DictionaryItem.BulkUpdateDictionaryItem(ctx, operations.BulkUpdateDictionaryItemRequest{
        BulkUpdateDictionaryListRequest: &shared.BulkUpdateDictionaryListRequest{
            Items: []shared.BulkUpdateDictionaryItem{
                shared.BulkUpdateDictionaryItem{
                    ItemKey: sdk.String("test-key"),
                    ItemValue: sdk.String("test-value"),
                    Op: shared.BulkUpdateDictionaryItemOpCreate.ToPointer(),
                },
                shared.BulkUpdateDictionaryItem{
                    ItemKey: sdk.String("test-key"),
                    ItemValue: sdk.String("test-value"),
                    Op: shared.BulkUpdateDictionaryItemOpCreate.ToPointer(),
                },
            },
        },
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.BulkUpdateDictionaryItemSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BulkUpdateDictionaryItem200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.BulkUpdateDictionaryItemRequest](../../models/operations/bulkupdatedictionaryitemrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.BulkUpdateDictionaryItemSecurity](../../models/operations/bulkupdatedictionaryitemsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.BulkUpdateDictionaryItemResponse](../../models/operations/bulkupdatedictionaryitemresponse.md), error**


## CreateDictionaryItem

Create DictionaryItem given service, dictionary ID, item key, and item value.

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
    res, err := s.DictionaryItem.CreateDictionaryItem(ctx, operations.CreateDictionaryItemRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        DictionaryItem: &shared.DictionaryItem{
            ItemKey: sdk.String("test-key"),
            ItemValue: sdk.String("test-value"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.CreateDictionaryItemSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DictionaryItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateDictionaryItemRequest](../../models/operations/createdictionaryitemrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.CreateDictionaryItemSecurity](../../models/operations/createdictionaryitemsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.CreateDictionaryItemResponse](../../models/operations/createdictionaryitemresponse.md), error**


## DeleteDictionaryItem

Delete DictionaryItem given service, dictionary ID, and item key.

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
    res, err := s.DictionaryItem.DeleteDictionaryItem(ctx, operations.DeleteDictionaryItemRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        DictionaryItemKey: "test-key",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.DeleteDictionaryItemSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDictionaryItem200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.DeleteDictionaryItemRequest](../../models/operations/deletedictionaryitemrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.DeleteDictionaryItemSecurity](../../models/operations/deletedictionaryitemsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.DeleteDictionaryItemResponse](../../models/operations/deletedictionaryitemresponse.md), error**


## GetDictionaryItem

Retrieve a single DictionaryItem given service, dictionary ID and item key.

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
    res, err := s.DictionaryItem.GetDictionaryItem(ctx, operations.GetDictionaryItemRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        DictionaryItemKey: "test-key",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.GetDictionaryItemSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DictionaryItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetDictionaryItemRequest](../../models/operations/getdictionaryitemrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.GetDictionaryItemSecurity](../../models/operations/getdictionaryitemsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.GetDictionaryItemResponse](../../models/operations/getdictionaryitemresponse.md), error**


## ListDictionaryItems

List of DictionaryItems given service and dictionary ID.

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
    res, err := s.DictionaryItem.ListDictionaryItems(ctx, operations.ListDictionaryItemsRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        Direction: shared.DirectionAscend.ToPointer(),
        Page: sdk.Int64(1),
        PerPage: sdk.Int64(20),
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        Sort: sdk.String("created"),
    }, operations.ListDictionaryItemsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DictionaryItemResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListDictionaryItemsRequest](../../models/operations/listdictionaryitemsrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.ListDictionaryItemsSecurity](../../models/operations/listdictionaryitemssecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.ListDictionaryItemsResponse](../../models/operations/listdictionaryitemsresponse.md), error**


## UpdateDictionaryItem

Update DictionaryItem given service, dictionary ID, item key, and item value.

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
    res, err := s.DictionaryItem.UpdateDictionaryItem(ctx, operations.UpdateDictionaryItemRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        DictionaryItem: &shared.DictionaryItem{
            ItemKey: sdk.String("test-key"),
            ItemValue: sdk.String("test-value"),
        },
        DictionaryItemKey: "test-key",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.UpdateDictionaryItemSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DictionaryItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateDictionaryItemRequest](../../models/operations/updatedictionaryitemrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.UpdateDictionaryItemSecurity](../../models/operations/updatedictionaryitemsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.UpdateDictionaryItemResponse](../../models/operations/updatedictionaryitemresponse.md), error**


## UpsertDictionaryItem

Upsert DictionaryItem given service, dictionary ID, item key, and item value.

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
    res, err := s.DictionaryItem.UpsertDictionaryItem(ctx, operations.UpsertDictionaryItemRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        DictionaryItem: &shared.DictionaryItem{
            ItemKey: sdk.String("test-key"),
            ItemValue: sdk.String("test-value"),
        },
        DictionaryItemKey: "test-key",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.UpsertDictionaryItemSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DictionaryItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpsertDictionaryItemRequest](../../models/operations/upsertdictionaryitemrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.UpsertDictionaryItemSecurity](../../models/operations/upsertdictionaryitemsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.UpsertDictionaryItemResponse](../../models/operations/upsertdictionaryitemresponse.md), error**

