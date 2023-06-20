# Snippet

## Overview

VCL Snippets are blocks of VCL logic inserted into your service's configuration that don't require custom VCL.

<https://developer.fastly.com/reference/api/vcl-services/snippet>
### Available Operations

* [CreateSnippet](#createsnippet) - Create a snippet
* [DeleteSnippet](#deletesnippet) - Delete a snippet
* [GetSnippet](#getsnippet) - Get a versioned snippet
* [GetSnippetDynamic](#getsnippetdynamic) - Get a dynamic snippet
* [ListSnippets](#listsnippets) - List snippets
* [UpdateSnippetDynamic](#updatesnippetdynamic) - Update a dynamic snippet

## CreateSnippet

Create a snippet for a particular service and version.

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

    ctx := context.Background()
    res, err := s.Snippet.CreateSnippet(ctx, operations.CreateSnippetRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        Snippet2: &shared.Snippet2{
            Content: sdk.String("provident"),
            Dynamic: shared.SnippetDynamicZero.ToPointer(),
            Name: sdk.String("test-snippet"),
            Priority: sdk.String("10"),
            Type: shared.SnippetTypeHit.ToPointer(),
        },
        VersionID: 1,
    }, operations.CreateSnippetSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SnippetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateSnippetRequest](../../models/operations/createsnippetrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.CreateSnippetSecurity](../../models/operations/createsnippetsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.CreateSnippetResponse](../../models/operations/createsnippetresponse.md), error**


## DeleteSnippet

Delete a specific snippet for a particular service and version.

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
    res, err := s.Snippet.DeleteSnippet(ctx, operations.DeleteSnippetRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        SnippetName: "test-snippet",
        VersionID: 1,
    }, operations.DeleteSnippetSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteSnippet200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteSnippetRequest](../../models/operations/deletesnippetrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.DeleteSnippetSecurity](../../models/operations/deletesnippetsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.DeleteSnippetResponse](../../models/operations/deletesnippetresponse.md), error**


## GetSnippet

Get a single snippet for a particular service and version.

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
    res, err := s.Snippet.GetSnippet(ctx, operations.GetSnippetRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        SnippetName: "test-snippet",
        VersionID: 1,
    }, operations.GetSnippetSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SnippetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetSnippetRequest](../../models/operations/getsnippetrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.GetSnippetSecurity](../../models/operations/getsnippetsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.GetSnippetResponse](../../models/operations/getsnippetresponse.md), error**


## GetSnippetDynamic

Get a single dynamic snippet for a particular service.

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
    res, err := s.Snippet.GetSnippetDynamic(ctx, operations.GetSnippetDynamicRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        SnippetID: "62Yd1WfiCBPENLloXfXmlO",
    }, operations.GetSnippetDynamicSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SnippetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetSnippetDynamicRequest](../../models/operations/getsnippetdynamicrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.GetSnippetDynamicSecurity](../../models/operations/getsnippetdynamicsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.GetSnippetDynamicResponse](../../models/operations/getsnippetdynamicresponse.md), error**


## ListSnippets

List all snippets for a particular service and version.

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
    res, err := s.Snippet.ListSnippets(ctx, operations.ListSnippetsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListSnippetsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SnippetResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListSnippetsRequest](../../models/operations/listsnippetsrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.ListSnippetsSecurity](../../models/operations/listsnippetssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.ListSnippetsResponse](../../models/operations/listsnippetsresponse.md), error**


## UpdateSnippetDynamic

Update a dynamic snippet for a particular service.

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

    ctx := context.Background()
    res, err := s.Snippet.UpdateSnippetDynamic(ctx, operations.UpdateSnippetDynamicRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        Snippet2: &shared.Snippet2{
            Content: sdk.String("quasi"),
            Dynamic: shared.SnippetDynamicOne.ToPointer(),
            Name: sdk.String("test-snippet"),
            Priority: sdk.String("10"),
            Type: shared.SnippetTypeHit.ToPointer(),
        },
        SnippetID: "62Yd1WfiCBPENLloXfXmlO",
    }, operations.UpdateSnippetDynamicSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SnippetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateSnippetDynamicRequest](../../models/operations/updatesnippetdynamicrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.UpdateSnippetDynamicSecurity](../../models/operations/updatesnippetdynamicsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.UpdateSnippetDynamicResponse](../../models/operations/updatesnippetdynamicresponse.md), error**

