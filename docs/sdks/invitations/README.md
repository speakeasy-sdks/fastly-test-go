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
    res, err := s.Invitations.CreateInvitation(ctx, &components.Invitation{
        Data: &components.InvitationData{
            Attributes: &components.InvitationDataAttributes{
                Role: components.RoleUserUser.ToPointer(),
            },
            Relationships: &components.RelationshipServiceInvitationsCreate{
                ServiceInvitations: &components.RelationshipServiceInvitationsCreateServiceInvitations{
                    Data: []components.ServiceInvitation{
                        components.ServiceInvitation{
                            Data: &components.ServiceInvitationData{
                                Attributes: &components.ServiceInvitationDataAttributes{},
                                Relationships: &components.ServiceInvitationDataRelationships{
                                    Service: &components.RelationshipMemberServiceInput{},
                                },
                            },
                        },
                    },
                },
            },
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

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.Invitation](../../models/components/invitation.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |


### Response

**[*operations.CreateInvitationResponse](../../models/operations/createinvitationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteInvitation

Delete an invitation.

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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListInvitations

List all invitations.

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
    res, err := s.Invitations.ListInvitations(ctx, operations.ListInvitationsRequest{
        PageNumber: fastlytestgo.Int64(1),
        PageSize: fastlytestgo.Int64(20),
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
