# Star
(*Star*)

## Overview

A star allows users to mark services of interest.

<https://developer.fastly.com/reference/api/account/star>
### Available Operations

* [CreateServiceStar](#createservicestar) - Create a star
* [DeleteServiceStar](#deleteservicestar) - Delete a star
* [GetServiceStar](#getservicestar) - Get a star
* [ListServiceStars](#listservicestars) - List stars

## CreateServiceStar

Create star.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Star.CreateServiceStar(ctx, &components.Star{
        Data: &components.StarData{
            Relationships: &components.RelationshipsForStarInput{
                Service: &components.RelationshipMemberServiceInput{},
                User: &components.RelationshipsForStarUserInput{
                    Data: &components.RelationshipsForStarDataInput{},
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StarResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [components.Star](../../models/components/star.md)    | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateServiceStarResponse](../../models/operations/createservicestarresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteServiceStar

Delete star.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
	"net/http"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Star.DeleteServiceStar(ctx, operations.DeleteServiceStarRequest{
        StarID: "3krg2uUGZzb2W9Euo4moOY",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteServiceStarRequest](../../models/operations/deleteservicestarrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteServiceStarResponse](../../models/operations/deleteservicestarresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetServiceStar

Show star.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Star.GetServiceStar(ctx, operations.GetServiceStarRequest{
        StarID: "3krg2uUGZzb2W9Euo4moOY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StarResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetServiceStarRequest](../../models/operations/getservicestarrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetServiceStarResponse](../../models/operations/getservicestarresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListServiceStars

List stars.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Star.ListServiceStars(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListServiceStarsResponse](../../models/operations/listservicestarsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
