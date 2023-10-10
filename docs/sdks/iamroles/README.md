# IamRoles
(*IamRoles*)

## Overview

A role is a collection of permissions that can be assigned to a user group.

<https://developer.fastly.com/reference/api/account/roles>
### Available Operations

* [DeleteARole](#deletearole) - Delete a role
* [GetARole](#getarole) - Get a role
* [ListRolePermissions](#listrolepermissions) - List permissions in a role
* [ListRoles](#listroles) - List roles

## DeleteARole

Delete a role.

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
    res, err := s.IamRoles.DeleteARole(ctx, operations.DeleteARoleRequest{
        RoleID: "t4Gg2uUGZzb2W9Euo4mo0R",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.DeleteARoleRequest](../../models/operations/deletearolerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.DeleteARoleResponse](../../models/operations/deletearoleresponse.md), error**


## GetARole

Get a role.

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
    res, err := s.IamRoles.GetARole(ctx, operations.GetARoleRequest{
        RoleID: "t4Gg2uUGZzb2W9Euo4mo0R",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetARole200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetARoleRequest](../../models/operations/getarolerequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetARoleResponse](../../models/operations/getaroleresponse.md), error**


## ListRolePermissions

List all permissions in a role.

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
    res, err := s.IamRoles.ListRolePermissions(ctx, operations.ListRolePermissionsRequest{
        RoleID: "t4Gg2uUGZzb2W9Euo4mo0R",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListRolePermissions200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListRolePermissionsRequest](../../models/operations/listrolepermissionsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListRolePermissionsResponse](../../models/operations/listrolepermissionsresponse.md), error**


## ListRoles

List all roles.

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
    res, err := s.IamRoles.ListRoles(ctx, operations.ListRolesRequest{
        Page: fastly.Int64(1),
        PerPage: fastly.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListRoles200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListRolesRequest](../../models/operations/listrolesrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.ListRolesResponse](../../models/operations/listrolesresponse.md), error**

