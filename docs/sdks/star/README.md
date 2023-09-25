# Star

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
	"context"
	"log"
	fastly "Fastly"
	"Fastly/pkg/models/shared"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Star.CreateServiceStar(ctx, shared.StarInput{
        Data: &shared.StarDataInput{
            Relationships: &shared.RelationshipsForStarInput{
                Service: &shared.RelationshipMemberServiceInput{
                    Type: fastly.String("officiis"),
                },
                User: &shared.RelationshipsForStarUserInput{
                    Data: &shared.RelationshipsForStarUserDataInput{
                        Type: fastly.String("temporibus"),
                    },
                },
            },
            Type: fastly.String("ullam"),
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
| `request`                                             | [shared.StarInput](../../models/shared/starinput.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateServiceStarResponse](../../models/operations/createservicestarresponse.md), error**


## DeleteServiceStar

Delete star.

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


## GetServiceStar

Show star.

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


## ListServiceStars

List stars.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/pkg/models/shared"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Star.ListServiceStars(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListServiceStars200ApplicationVndAPIPlusJSONObject != nil {
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

