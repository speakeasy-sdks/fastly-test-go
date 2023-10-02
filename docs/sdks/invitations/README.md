# Invitations
(*Invitations*)

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
    res, err := s.Invitations.CreateInvitation(ctx, shared.InvitationInput{
        Data: &shared.InvitationDataInput{
            Attributes: &shared.InvitationDataAttributes{
                Email: fastly.String("Kenneth.Ullrich36@gmail.com"),
                LimitServices: fastly.Bool(false),
                Role: shared.RoleUserUser.ToPointer(),
                StatusCode: shared.InvitationDataAttributesStatusCodeOne.ToPointer(),
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
                                        Type: fastly.String("Developer"),
                                    },
                                },
                                Type: fastly.String("migration"),
                            },
                        },
                    },
                },
            },
            Type: fastly.String("matrix"),
        },
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

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.InvitationInput](../../models/shared/invitationinput.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


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
    res, err := s.Invitations.DeleteInvitation(ctx, operations.DeleteInvitationRequest{
        InvitationID: "3krg2uUGZzb2W9Euo4moOY",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteInvitationRequest](../../models/operations/deleteinvitationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
    res, err := s.Invitations.ListInvitations(ctx, operations.ListInvitationsRequest{
        PageNumber: fastly.Int64(1),
        PageSize: fastly.Int64(20),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListInvitationsRequest](../../models/operations/listinvitationsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListInvitationsResponse](../../models/operations/listinvitationsresponse.md), error**

