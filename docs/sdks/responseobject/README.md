# ResponseObject

## Overview

Allows you to create synthetic responses that exist entirely on the varnish machine. Useful for creating error or maintenance pages that exists outside the scope of your backend architecture. Best when used with [Condition](#condition) objects.

<https://developer.fastly.com/reference/api/vcl-services/response-object>
### Available Operations

* [DeleteResponseObject](#deleteresponseobject) - Delete a Response Object
* [GetResponseObject](#getresponseobject) - Get a Response object
* [ListResponseObjects](#listresponseobjects) - List Response objects

## DeleteResponseObject

Deletes the specified Response Object.

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
    res, err := s.ResponseObject.DeleteResponseObject(ctx, operations.DeleteResponseObjectRequest{
        ResponseObjectName: "test-response",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteResponseObjectSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteResponseObject200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.DeleteResponseObjectRequest](../../models/operations/deleteresponseobjectrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.DeleteResponseObjectSecurity](../../models/operations/deleteresponseobjectsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.DeleteResponseObjectResponse](../../models/operations/deleteresponseobjectresponse.md), error**


## GetResponseObject

Gets the specified Response Object.

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
    res, err := s.ResponseObject.GetResponseObject(ctx, operations.GetResponseObjectRequest{
        ResponseObjectName: "test-response",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetResponseObjectSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponseObjectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetResponseObjectRequest](../../models/operations/getresponseobjectrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.GetResponseObjectSecurity](../../models/operations/getresponseobjectsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.GetResponseObjectResponse](../../models/operations/getresponseobjectresponse.md), error**


## ListResponseObjects

Returns all Response Objects for the specified service and version.

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
    res, err := s.ResponseObject.ListResponseObjects(ctx, operations.ListResponseObjectsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListResponseObjectsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponseObjectResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListResponseObjectsRequest](../../models/operations/listresponseobjectsrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.ListResponseObjectsSecurity](../../models/operations/listresponseobjectssecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.ListResponseObjectsResponse](../../models/operations/listresponseobjectsresponse.md), error**

