# ServiceAuthorizations
(*.ServiceAuthorizations*)

## Overview

A service authorization allows limited users to access only specified services.

<https://developer.fastly.com/reference/api/account/service-authorization>
### Available Operations

* [CreateServiceAuthorization](#createserviceauthorization) - Create service authorization
* [DeleteServiceAuthorization](#deleteserviceauthorization) - Delete service authorization
* [ListServiceAuthorization](#listserviceauthorization) - List service authorizations
* [ShowServiceAuthorization](#showserviceauthorization) - Show service authorization
* [UpdateServiceAuthorization](#updateserviceauthorization) - Update service authorization

## CreateServiceAuthorization

Create service authorization.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.ServiceAuthorizations.CreateServiceAuthorization(ctx, &components.ServiceAuthorization{
        Data: &components.ServiceAuthorizationData{
            Attributes: &components.ServiceAuthorizationDataAttributes{
                Permission: components.PermissionFull.ToPointer(),
            },
            Relationships: &components.ServiceAuthorizationDataRelationships{
                Service: &components.RelationshipMemberServiceInput{},
                User: &components.ServiceAuthorizationDataUser{
                    Data: &components.ServiceAuthorizationDataData{},
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServiceAuthorizationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.ServiceAuthorization](../../models/shared/serviceauthorization.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateServiceAuthorizationResponse](../../models/operations/createserviceauthorizationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteServiceAuthorization

Delete service authorization.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.ServiceAuthorizations.DeleteServiceAuthorization(ctx, operations.DeleteServiceAuthorizationRequest{
        ServiceAuthorizationID: "3krg2uUGZzb2W9Euo4moOY",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.DeleteServiceAuthorizationRequest](../../models/operations/deleteserviceauthorizationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.DeleteServiceAuthorizationResponse](../../models/operations/deleteserviceauthorizationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListServiceAuthorization

List service authorizations.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.ServiceAuthorizations.ListServiceAuthorization(ctx, operations.ListServiceAuthorizationRequest{
        PageNumber: fastly.Int64(1),
        PageSize: fastly.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServiceAuthorizationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListServiceAuthorizationRequest](../../models/operations/listserviceauthorizationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.ListServiceAuthorizationResponse](../../models/operations/listserviceauthorizationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ShowServiceAuthorization

Show service authorization.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.ServiceAuthorizations.ShowServiceAuthorization(ctx, operations.ShowServiceAuthorizationRequest{
        ServiceAuthorizationID: "3krg2uUGZzb2W9Euo4moOY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServiceAuthorizationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ShowServiceAuthorizationRequest](../../models/operations/showserviceauthorizationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.ShowServiceAuthorizationResponse](../../models/operations/showserviceauthorizationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateServiceAuthorization

Update service authorization.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.ServiceAuthorizations.UpdateServiceAuthorization(ctx, operations.UpdateServiceAuthorizationRequest{
        ServiceAuthorization: &components.ServiceAuthorization{
            Data: &components.ServiceAuthorizationData{
                Attributes: &components.ServiceAuthorizationDataAttributes{
                    Permission: components.PermissionFull.ToPointer(),
                },
                Relationships: &components.ServiceAuthorizationDataRelationships{
                    Service: &components.RelationshipMemberServiceInput{},
                    User: &components.ServiceAuthorizationDataUser{
                        Data: &components.ServiceAuthorizationDataData{},
                    },
                },
            },
        },
        ServiceAuthorizationID: "3krg2uUGZzb2W9Euo4moOY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServiceAuthorizationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.UpdateServiceAuthorizationRequest](../../models/operations/updateserviceauthorizationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.UpdateServiceAuthorizationResponse](../../models/operations/updateserviceauthorizationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
