# Invitations

## Overview

Invitations allow superusers and engineers to invite users to set up accounts as collaborators under a main customer account. Superusers can invite collaborators and assign them any role or permission level on a per-service basis. Engineers with no per-service limitations on their role can only invite new collaborators but cannot modify their permissions.

<https://developer.fastly.com/reference/api/account/invitations>
### Available Operations

* [CreateInvitation](#createinvitation) - Create an invitation
* [DeleteInvitation](#deleteinvitation) - Delete an invitation
* [ListInvitations](#listinvitations) - List invitations

## CreateInvitation

Create an invitation.

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
    res, err := s.Invitations.CreateInvitation(ctx, shared.InvitationInput{
        Data: &shared.InvitationDataInput{
            Attributes: &shared.InvitationDataAttributes{
                Email: sdk.String("Kennedy20@yahoo.com"),
                LimitServices: sdk.Bool(false),
                Role: shared.RoleUserUser.ToPointer(),
                StatusCode: shared.InvitationDataAttributesStatusCodeZero.ToPointer(),
            },
            Relationships: &shared.RelationshipServiceInvitationsCreateInput{
                ServiceInvitations: &shared.RelationshipServiceInvitationsCreateServiceInvitationsInput{
                    Data: []shared.ServiceInvitationInput{
                        shared.ServiceInvitationInput{
                            Data: &shared.ServiceInvitationDataInput{
                                Attributes: &shared.ServiceInvitationDataAttributes{
                                    Permission: shared.ServiceInvitationDataAttributesPermissionReadOnly.ToPointer(),
                                },
                                Relationships: &shared.ServiceInvitationDataRelationshipsInput{
                                    Service: &shared.RelationshipMemberServiceInput{
                                        Type: shared.TypeServiceService.ToPointer(),
                                    },
                                },
                                Type: shared.TypeServiceInvitationServiceInvitation.ToPointer(),
                            },
                        },
                    },
                },
            },
            Type: shared.TypeInvitationInvitation.ToPointer(),
        },
    }, operations.CreateInvitationSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.InvitationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.InvitationInput](../../models/shared/invitationinput.md)                           | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.CreateInvitationSecurity](../../models/operations/createinvitationsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.CreateInvitationResponse](../../models/operations/createinvitationresponse.md), error**


## DeleteInvitation

Delete an invitation.

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
    res, err := s.Invitations.DeleteInvitation(ctx, operations.DeleteInvitationRequest{
        InvitationID: "3krg2uUGZzb2W9Euo4moOY",
    }, operations.DeleteInvitationSecurity{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteInvitationRequest](../../models/operations/deleteinvitationrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.DeleteInvitationSecurity](../../models/operations/deleteinvitationsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.DeleteInvitationResponse](../../models/operations/deleteinvitationresponse.md), error**


## ListInvitations

List all invitations.

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
    res, err := s.Invitations.ListInvitations(ctx, operations.ListInvitationsRequest{
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
    }, operations.ListInvitationsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.InvitationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListInvitationsRequest](../../models/operations/listinvitationsrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ListInvitationsSecurity](../../models/operations/listinvitationssecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ListInvitationsResponse](../../models/operations/listinvitationsresponse.md), error**
