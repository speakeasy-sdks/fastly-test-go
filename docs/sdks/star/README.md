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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.CreateServiceStarSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Star.CreateServiceStar(ctx, shared.StarInput{
        Data: &shared.StarDataInput{
            Relationships: &shared.RelationshipsForStarInput{
                Service: &shared.RelationshipMemberServiceInput{
                    Type: shared.TypeServiceService.ToPointer(),
                },
                User: &shared.RelationshipsForStarUserInput{
                    Data: &shared.RelationshipsForStarUserDataInput{
                        Type: shared.TypeUserUser.ToPointer(),
                    },
                },
            },
            Type: shared.TypeStarStar.ToPointer(),
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StarResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.StarInput](../../models/shared/starinput.md)                                         | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.CreateServiceStarSecurity](../../models/operations/createservicestarsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.DeleteServiceStarSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Star.DeleteServiceStar(ctx, operations.DeleteServiceStarRequest{
        StarID: "3krg2uUGZzb2W9Euo4moOY",
    }, operationSecurity)
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
| `request`                                                                                    | [operations.DeleteServiceStarRequest](../../models/operations/deleteservicestarrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.DeleteServiceStarSecurity](../../models/operations/deleteservicestarsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.GetServiceStarSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Star.GetServiceStar(ctx, operations.GetServiceStarRequest{
        StarID: "3krg2uUGZzb2W9Euo4moOY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StarResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetServiceStarRequest](../../models/operations/getservicestarrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetServiceStarSecurity](../../models/operations/getservicestarsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListServiceStarsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Star.ListServiceStars(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListServiceStars200ApplicationVndAPIPlusJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `security`                                                                                 | [operations.ListServiceStarsSecurity](../../models/operations/listservicestarssecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ListServiceStarsResponse](../../models/operations/listservicestarsresponse.md), error**

