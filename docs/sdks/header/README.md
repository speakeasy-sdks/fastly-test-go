# Header
(*Header*)

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
    res, err := s.Header.CreateHeaderObject(ctx, operations.CreateHeaderObjectRequest{
        Header: &shared.Header{
            Action: shared.HeaderActionAppend.ToPointer(),
            CacheCondition: fastly.String("null"),
            Dst: fastly.String("anti gold"),
            IgnoreIfSet: fastly.Int64(802152),
            Name: fastly.String("test-header"),
            Priority: fastly.Int64(352185),
            Regex: fastly.String("TCP infomediaries"),
            RequestCondition: fastly.String("null"),
            ResponseCondition: fastly.String("monitor Northeast"),
            Src: fastly.String("briefly Account Secured"),
            Substitution: fastly.String("Egypt gee"),
            Type: shared.HeaderTypeCache.ToPointer(),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HeaderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateHeaderObjectRequest](../../models/operations/createheaderobjectrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


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
    res, err := s.Header.DeleteHeaderObject(ctx, operations.DeleteHeaderObjectRequest{
        HeaderName: "test-header",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteHeaderObject200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteHeaderObjectRequest](../../models/operations/deleteheaderobjectrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


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
    res, err := s.Header.GetHeaderObject(ctx, operations.GetHeaderObjectRequest{
        HeaderName: "test-header",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HeaderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetHeaderObjectRequest](../../models/operations/getheaderobjectrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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
    res, err := s.Header.ListHeaderObjects(ctx, operations.ListHeaderObjectsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HeaderResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListHeaderObjectsRequest](../../models/operations/listheaderobjectsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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
    res, err := s.Header.UpdateHeaderObject(ctx, operations.UpdateHeaderObjectRequest{
        Header: &shared.Header{
            Action: shared.HeaderActionRegexRepeat.ToPointer(),
            CacheCondition: fastly.String("null"),
            Dst: fastly.String("male"),
            IgnoreIfSet: fastly.Int64(926675),
            Name: fastly.String("test-header"),
            Priority: fastly.Int64(444514),
            Regex: fastly.String("North North"),
            RequestCondition: fastly.String("null"),
            ResponseCondition: fastly.String("candela program 24"),
            Src: fastly.String("Salad"),
            Substitution: fastly.String("tempora Buckinghamshire alarm"),
            Type: shared.HeaderTypeRequest.ToPointer(),
        },
        HeaderName: "test-header",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HeaderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateHeaderObjectRequest](../../models/operations/updateheaderobjectrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateHeaderObjectResponse](../../models/operations/updateheaderobjectresponse.md), error**

