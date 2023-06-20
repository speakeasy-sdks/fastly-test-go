# IamServiceGroups

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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.IamServiceGroups.DeleteAServiceGroup(ctx, operations.DeleteAServiceGroupRequest{
        ServiceGroupID: "t4Gg2uUGZzb2W9Euo4mo0R",
    }, operations.DeleteAServiceGroupSecurity{
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteAServiceGroupRequest](../../models/operations/deleteaservicegrouprequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.DeleteAServiceGroupSecurity](../../models/operations/deleteaservicegroupsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.IamServiceGroups.GetAServiceGroup(ctx, operations.GetAServiceGroupRequest{
        ServiceGroupID: "t4Gg2uUGZzb2W9Euo4mo0R",
    }, operations.GetAServiceGroupSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAServiceGroup200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetAServiceGroupRequest](../../models/operations/getaservicegrouprequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.GetAServiceGroupSecurity](../../models/operations/getaservicegroupsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.IamServiceGroups.ListServiceGroupServices(ctx, operations.ListServiceGroupServicesRequest{
        Page: sdk.Int64(1),
        PerPage: sdk.Int64(20),
        ServiceGroupID: "t4Gg2uUGZzb2W9Euo4mo0R",
    }, operations.ListServiceGroupServicesSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListServiceGroupServices200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListServiceGroupServicesRequest](../../models/operations/listservicegroupservicesrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.ListServiceGroupServicesSecurity](../../models/operations/listservicegroupservicessecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.IamServiceGroups.ListServiceGroups(ctx, operations.ListServiceGroupsRequest{
        Page: sdk.Int64(1),
        PerPage: sdk.Int64(20),
    }, operations.ListServiceGroupsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListServiceGroups200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListServiceGroupsRequest](../../models/operations/listservicegroupsrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ListServiceGroupsSecurity](../../models/operations/listservicegroupssecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.ListServiceGroupsResponse](../../models/operations/listservicegroupsresponse.md), error**

