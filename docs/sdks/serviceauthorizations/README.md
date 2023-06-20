# ServiceAuthorizations

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
	"Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.ServiceAuthorizations.CreateServiceAuthorization(ctx, shared.ServiceAuthorizationInput{
        Data: &shared.ServiceAuthorizationDataInput{
            Attributes: &shared.ServiceAuthorizationDataAttributes{
                Permission: shared.PermissionFull.ToPointer(),
            },
            Relationships: &shared.ServiceAuthorizationDataRelationshipsInput{
                Service: &shared.RelationshipMemberServiceInput{
                    Type: shared.TypeServiceService.ToPointer(),
                },
                User: &shared.ServiceAuthorizationDataRelationshipsUserInput{
                    Data: &shared.ServiceAuthorizationDataRelationshipsUserDataInput{
                        Type: shared.TypeUserUser.ToPointer(),
                    },
                },
            },
            Type: shared.TypeServiceAuthorizationServiceAuthorization.ToPointer(),
        },
    }, operations.CreateServiceAuthorizationSecurity{
        Token: "",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [shared.ServiceAuthorizationInput](../../models/shared/serviceauthorizationinput.md)                           | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.CreateServiceAuthorizationSecurity](../../models/operations/createserviceauthorizationsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.CreateServiceAuthorizationResponse](../../models/operations/createserviceauthorizationresponse.md), error**


## DeleteServiceAuthorization

Delete service authorization.

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
    res, err := s.ServiceAuthorizations.DeleteServiceAuthorization(ctx, operations.DeleteServiceAuthorizationRequest{
        ServiceAuthorizationID: "3krg2uUGZzb2W9Euo4moOY",
    }, operations.DeleteServiceAuthorizationSecurity{
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteServiceAuthorizationRequest](../../models/operations/deleteserviceauthorizationrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.DeleteServiceAuthorizationSecurity](../../models/operations/deleteserviceauthorizationsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.DeleteServiceAuthorizationResponse](../../models/operations/deleteserviceauthorizationresponse.md), error**


## ListServiceAuthorization

List service authorizations.

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
    res, err := s.ServiceAuthorizations.ListServiceAuthorization(ctx, operations.ListServiceAuthorizationRequest{
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
    }, operations.ListServiceAuthorizationSecurity{
        Token: "",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListServiceAuthorizationRequest](../../models/operations/listserviceauthorizationrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.ListServiceAuthorizationSecurity](../../models/operations/listserviceauthorizationsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.ListServiceAuthorizationResponse](../../models/operations/listserviceauthorizationresponse.md), error**


## ShowServiceAuthorization

Show service authorization.

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
    res, err := s.ServiceAuthorizations.ShowServiceAuthorization(ctx, operations.ShowServiceAuthorizationRequest{
        ServiceAuthorizationID: "3krg2uUGZzb2W9Euo4moOY",
    }, operations.ShowServiceAuthorizationSecurity{
        Token: "",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ShowServiceAuthorizationRequest](../../models/operations/showserviceauthorizationrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.ShowServiceAuthorizationSecurity](../../models/operations/showserviceauthorizationsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.ShowServiceAuthorizationResponse](../../models/operations/showserviceauthorizationresponse.md), error**


## UpdateServiceAuthorization

Update service authorization.

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
    res, err := s.ServiceAuthorizations.UpdateServiceAuthorization(ctx, operations.UpdateServiceAuthorizationRequest{
        ServiceAuthorizationInput: &shared.ServiceAuthorizationInput{
            Data: &shared.ServiceAuthorizationDataInput{
                Attributes: &shared.ServiceAuthorizationDataAttributes{
                    Permission: shared.PermissionFull.ToPointer(),
                },
                Relationships: &shared.ServiceAuthorizationDataRelationshipsInput{
                    Service: &shared.RelationshipMemberServiceInput{
                        Type: shared.TypeServiceService.ToPointer(),
                    },
                    User: &shared.ServiceAuthorizationDataRelationshipsUserInput{
                        Data: &shared.ServiceAuthorizationDataRelationshipsUserDataInput{
                            Type: shared.TypeUserUser.ToPointer(),
                        },
                    },
                },
                Type: shared.TypeServiceAuthorizationServiceAuthorization.ToPointer(),
            },
        },
        ServiceAuthorizationID: "3krg2uUGZzb2W9Euo4moOY",
    }, operations.UpdateServiceAuthorizationSecurity{
        Token: "",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.UpdateServiceAuthorizationRequest](../../models/operations/updateserviceauthorizationrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.UpdateServiceAuthorizationSecurity](../../models/operations/updateserviceauthorizationsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.UpdateServiceAuthorizationResponse](../../models/operations/updateserviceauthorizationresponse.md), error**

