# Header

## Overview

Header objects are used to add, modify, or delete headers from requests and responses. The header content can be simple strings or be derived from variables inside Varnish. Regular expressions can be used to customize the headers even further.

<https://developer.fastly.com/reference/api/vcl-services/header>
### Available Operations

* [CreateHeaderObject](#createheaderobject) - Create a Header object
* [DeleteHeaderObject](#deleteheaderobject) - Delete a Header object
* [GetHeaderObject](#getheaderobject) - Get a Header object
* [ListHeaderObjects](#listheaderobjects) - List Header objects
* [UpdateHeaderObject](#updateheaderobject) - Update a Header object

## CreateHeaderObject

Creates a new Header object.

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
    operationSecurity := operations.CreateHeaderObjectSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Header.CreateHeaderObject(ctx, operations.CreateHeaderObjectRequest{
        Header2: &shared.Header2{
            Action: shared.HeaderActionDelete.ToPointer(),
            CacheCondition: sdk.String("null"),
            Dst: sdk.String("accusantium"),
            IgnoreIfSet: sdk.Int64(653201),
            Name: sdk.String("test-header"),
            Priority: sdk.Int64(968962),
            Regex: sdk.String("mollitia"),
            RequestCondition: sdk.String("null"),
            ResponseCondition: sdk.String("ad"),
            Src: sdk.String("eum"),
            Substitution: sdk.String("dolor"),
            Type: shared.HeaderTypeResponse.ToPointer(),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HeaderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateHeaderObjectRequest](../../models/operations/createheaderobjectrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.CreateHeaderObjectSecurity](../../models/operations/createheaderobjectsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.CreateHeaderObjectResponse](../../models/operations/createheaderobjectresponse.md), error**


## DeleteHeaderObject

Deletes a Header object by name.

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
    operationSecurity := operations.DeleteHeaderObjectSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Header.DeleteHeaderObject(ctx, operations.DeleteHeaderObjectRequest{
        HeaderName: "test-header",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteHeaderObject200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteHeaderObjectRequest](../../models/operations/deleteheaderobjectrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.DeleteHeaderObjectSecurity](../../models/operations/deleteheaderobjectsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.DeleteHeaderObjectResponse](../../models/operations/deleteheaderobjectresponse.md), error**


## GetHeaderObject

Retrieves a Header object by name.

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
    operationSecurity := operations.GetHeaderObjectSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Header.GetHeaderObject(ctx, operations.GetHeaderObjectRequest{
        HeaderName: "test-header",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HeaderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHeaderObjectRequest](../../models/operations/getheaderobjectrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.GetHeaderObjectSecurity](../../models/operations/getheaderobjectsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.GetHeaderObjectResponse](../../models/operations/getheaderobjectresponse.md), error**


## ListHeaderObjects

Retrieves all Header objects for a particular Version of a Service.

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
    operationSecurity := operations.ListHeaderObjectsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Header.ListHeaderObjects(ctx, operations.ListHeaderObjectsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HeaderResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHeaderObjectsRequest](../../models/operations/listheaderobjectsrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ListHeaderObjectsSecurity](../../models/operations/listheaderobjectssecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.ListHeaderObjectsResponse](../../models/operations/listheaderobjectsresponse.md), error**


## UpdateHeaderObject

Modifies an existing Header object by name.

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
    operationSecurity := operations.UpdateHeaderObjectSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Header.UpdateHeaderObject(ctx, operations.UpdateHeaderObjectRequest{
        Header2: &shared.Header2{
            Action: shared.HeaderActionSet.ToPointer(),
            CacheCondition: sdk.String("null"),
            Dst: sdk.String("nemo"),
            IgnoreIfSet: sdk.Int64(97260),
            Name: sdk.String("test-header"),
            Priority: sdk.Int64(435865),
            Regex: sdk.String("doloribus"),
            RequestCondition: sdk.String("null"),
            ResponseCondition: sdk.String("debitis"),
            Src: sdk.String("eius"),
            Substitution: sdk.String("maxime"),
            Type: shared.HeaderTypeCache.ToPointer(),
        },
        HeaderName: "test-header",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HeaderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateHeaderObjectRequest](../../models/operations/updateheaderobjectrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.UpdateHeaderObjectSecurity](../../models/operations/updateheaderobjectsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.UpdateHeaderObjectResponse](../../models/operations/updateheaderobjectresponse.md), error**

