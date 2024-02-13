# DictionaryItem
(*DictionaryItem*)

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
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.DictionaryItem.BulkUpdateDictionaryItem(ctx, operations.BulkUpdateDictionaryItemRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.BulkUpdateDictionaryItemRequest](../../models/operations/bulkupdatedictionaryitemrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.BulkUpdateDictionaryItemResponse](../../models/operations/bulkupdatedictionaryitemresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateDictionaryItem

Create DictionaryItem given service, dictionary ID, item key, and item value.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.DictionaryItem.CreateDictionaryItem(ctx, operations.CreateDictionaryItemRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateDictionaryItemRequest](../../models/operations/createdictionaryitemrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.CreateDictionaryItemResponse](../../models/operations/createdictionaryitemresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteDictionaryItem

Delete DictionaryItem given service, dictionary ID, and item key.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.DictionaryItem.DeleteDictionaryItem(ctx, operations.DeleteDictionaryItemRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        DictionaryItemKey: "test-key",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteDictionaryItemRequest](../../models/operations/deletedictionaryitemrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.DeleteDictionaryItemResponse](../../models/operations/deletedictionaryitemresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDictionaryItem

Retrieve a single DictionaryItem given service, dictionary ID and item key.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.DictionaryItem.GetDictionaryItem(ctx, operations.GetDictionaryItemRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        DictionaryItemKey: "test-key",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetDictionaryItemRequest](../../models/operations/getdictionaryitemrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetDictionaryItemResponse](../../models/operations/getdictionaryitemresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListDictionaryItems

List of DictionaryItems given service and dictionary ID.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.DictionaryItem.ListDictionaryItems(ctx, operations.ListDictionaryItemsRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        Direction: components.DirectionAscend.ToPointer(),
        Page: fastlytestgo.Int64(1),
        PerPage: fastlytestgo.Int64(20),
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        Sort: fastlytestgo.String("created"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListDictionaryItemsRequest](../../models/operations/listdictionaryitemsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListDictionaryItemsResponse](../../models/operations/listdictionaryitemsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateDictionaryItem

Update DictionaryItem given service, dictionary ID, item key, and item value.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.DictionaryItem.UpdateDictionaryItem(ctx, operations.UpdateDictionaryItemRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        DictionaryItemKey: "test-key",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateDictionaryItemRequest](../../models/operations/updatedictionaryitemrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.UpdateDictionaryItemResponse](../../models/operations/updatedictionaryitemresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpsertDictionaryItem

Upsert DictionaryItem given service, dictionary ID, item key, and item value.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.DictionaryItem.UpsertDictionaryItem(ctx, operations.UpsertDictionaryItemRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        DictionaryItemKey: "test-key",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpsertDictionaryItemRequest](../../models/operations/upsertdictionaryitemrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.UpsertDictionaryItemResponse](../../models/operations/upsertdictionaryitemresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
