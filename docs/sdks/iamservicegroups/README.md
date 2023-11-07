# IamServiceGroups
(*.IamServiceGroups*)

## Overview

A service group is a collection of services that can be assigned to a user group.

<https://developer.fastly.com/reference/api/account/service-groups>
### Available Operations

* [DeleteAServiceGroup](#deleteaservicegroup) - Delete a service group
* [GetAServiceGroup](#getaservicegroup) - Get a service group
* [ListServiceGroupServices](#listservicegroupservices) - List services to a service group
* [ListServiceGroups](#listservicegroups) - List service groups

## DeleteAServiceGroup

Delete a service group.

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
    res, err := s.IamServiceGroups.DeleteAServiceGroup(ctx, operations.DeleteAServiceGroupRequest{
        ServiceGroupID: "t4Gg2uUGZzb2W9Euo4mo0R",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteAServiceGroupRequest](../../models/operations/deleteaservicegrouprequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteAServiceGroupResponse](../../models/operations/deleteaservicegroupresponse.md), error**


## GetAServiceGroup

Get a service group.

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
    res, err := s.IamServiceGroups.GetAServiceGroup(ctx, operations.GetAServiceGroupRequest{
        ServiceGroupID: "t4Gg2uUGZzb2W9Euo4mo0R",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetAServiceGroupRequest](../../models/operations/getaservicegrouprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetAServiceGroupResponse](../../models/operations/getaservicegroupresponse.md), error**


## ListServiceGroupServices

List services to a service group.

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
    res, err := s.IamServiceGroups.ListServiceGroupServices(ctx, operations.ListServiceGroupServicesRequest{
        Page: fastly.Int64(1),
        PerPage: fastly.Int64(20),
        ServiceGroupID: "t4Gg2uUGZzb2W9Euo4mo0R",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListServiceGroupServicesRequest](../../models/operations/listservicegroupservicesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.ListServiceGroupServicesResponse](../../models/operations/listservicegroupservicesresponse.md), error**


## ListServiceGroups

List all service groups.

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
    res, err := s.IamServiceGroups.ListServiceGroups(ctx, operations.ListServiceGroupsRequest{
        Page: fastly.Int64(1),
        PerPage: fastly.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListServiceGroupsRequest](../../models/operations/listservicegroupsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListServiceGroupsResponse](../../models/operations/listservicegroupsresponse.md), error**

