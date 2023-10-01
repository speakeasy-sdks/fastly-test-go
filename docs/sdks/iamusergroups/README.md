# IamUserGroups
(*IamUserGroups*)

## Overview

A user group is a collection of users with service groups and roles.

<https://developer.fastly.com/reference/api/account/user-groups>
### Available Operations

* [DeleteAUserGroup](#deleteausergroup) - Delete a user group
* [GetAUserGroup](#getausergroup) - Get a user group
* [ListUserGroupMembers](#listusergroupmembers) - List members of a user group
* [ListUserGroupRoles](#listusergrouproles) - List roles in a user group
* [ListUserGroupServiceGroups](#listusergroupservicegroups) - List service groups in a user group
* [ListUserGroups](#listusergroups) - List user groups

## DeleteAUserGroup

Delete a user group.

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
    res, err := s.IamUserGroups.DeleteAUserGroup(ctx, operations.DeleteAUserGroupRequest{
        UserGroupID: "t4Gg2uUGZzb2W9Euo4mo0R",
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
| `request`                                                                                | [operations.DeleteAUserGroupRequest](../../models/operations/deleteausergrouprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeleteAUserGroupResponse](../../models/operations/deleteausergroupresponse.md), error**


## GetAUserGroup

Get a user group.

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
    res, err := s.IamUserGroups.GetAUserGroup(ctx, operations.GetAUserGroupRequest{
        UserGroupID: "t4Gg2uUGZzb2W9Euo4mo0R",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAUserGroup200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetAUserGroupRequest](../../models/operations/getausergrouprequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetAUserGroupResponse](../../models/operations/getausergroupresponse.md), error**


## ListUserGroupMembers

List members of a user group.

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
    res, err := s.IamUserGroups.ListUserGroupMembers(ctx, operations.ListUserGroupMembersRequest{
        Page: fastly.Int64(1),
        PerPage: fastly.Int64(20),
        UserGroupID: "t4Gg2uUGZzb2W9Euo4mo0R",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListUserGroupMembers200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListUserGroupMembersRequest](../../models/operations/listusergroupmembersrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ListUserGroupMembersResponse](../../models/operations/listusergroupmembersresponse.md), error**


## ListUserGroupRoles

List roles in a user group.

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
    res, err := s.IamUserGroups.ListUserGroupRoles(ctx, operations.ListUserGroupRolesRequest{
        Page: fastly.Int64(1),
        PerPage: fastly.Int64(20),
        UserGroupID: "t4Gg2uUGZzb2W9Euo4mo0R",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListUserGroupRoles200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListUserGroupRolesRequest](../../models/operations/listusergrouprolesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListUserGroupRolesResponse](../../models/operations/listusergrouprolesresponse.md), error**


## ListUserGroupServiceGroups

List service groups in a user group.

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
    res, err := s.IamUserGroups.ListUserGroupServiceGroups(ctx, operations.ListUserGroupServiceGroupsRequest{
        Page: fastly.Int64(1),
        PerPage: fastly.Int64(20),
        UserGroupID: "t4Gg2uUGZzb2W9Euo4mo0R",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListUserGroupServiceGroups200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListUserGroupServiceGroupsRequest](../../models/operations/listusergroupservicegroupsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.ListUserGroupServiceGroupsResponse](../../models/operations/listusergroupservicegroupsresponse.md), error**


## ListUserGroups

List all user groups.

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
    res, err := s.IamUserGroups.ListUserGroups(ctx, operations.ListUserGroupsRequest{
        Page: fastly.Int64(1),
        PerPage: fastly.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListUserGroups200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListUserGroupsRequest](../../models/operations/listusergroupsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListUserGroupsResponse](../../models/operations/listusergroupsresponse.md), error**

