# ResponseObject
(*ResponseObject*)

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
    res, err := s.ResponseObject.DeleteResponseObject(ctx, operations.DeleteResponseObjectRequest{
        ResponseObjectName: "test-response",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteResponseObjectRequest](../../models/operations/deleteresponseobjectrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.DeleteResponseObjectResponse](../../models/operations/deleteresponseobjectresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetResponseObject

Gets the specified Response Object.

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
    res, err := s.ResponseObject.GetResponseObject(ctx, operations.GetResponseObjectRequest{
        ResponseObjectName: "test-response",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetResponseObjectRequest](../../models/operations/getresponseobjectrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetResponseObjectResponse](../../models/operations/getresponseobjectresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListResponseObjects

Returns all Response Objects for the specified service and version.

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
    res, err := s.ResponseObject.ListResponseObjects(ctx, operations.ListResponseObjectsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListResponseObjectsRequest](../../models/operations/listresponseobjectsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListResponseObjectsResponse](../../models/operations/listresponseobjectsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
