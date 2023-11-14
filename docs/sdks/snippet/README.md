# Snippet
(*Snippet*)

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
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Snippet.CreateSnippet(ctx, operations.CreateSnippetRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        Snippet: &components.Snippet{
            Name: fastlytestgo.String("test-snippet"),
            Priority: fastlytestgo.String("10"),
        },
        VersionID: 1,
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateSnippetRequest](../../models/operations/createsnippetrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateSnippetResponse](../../models/operations/createsnippetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteSnippet

Delete a specific snippet for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Snippet.DeleteSnippet(ctx, operations.DeleteSnippetRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        SnippetName: "test-snippet",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteSnippetRequest](../../models/operations/deletesnippetrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.DeleteSnippetResponse](../../models/operations/deletesnippetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSnippet

Get a single snippet for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Snippet.GetSnippet(ctx, operations.GetSnippetRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        SnippetName: "test-snippet",
        VersionID: 1,
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetSnippetRequest](../../models/operations/getsnippetrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetSnippetResponse](../../models/operations/getsnippetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSnippetDynamic

Get a single dynamic snippet for a particular service.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Snippet.GetSnippetDynamic(ctx, operations.GetSnippetDynamicRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        SnippetID: "62Yd1WfiCBPENLloXfXmlO",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetSnippetDynamicRequest](../../models/operations/getsnippetdynamicrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetSnippetDynamicResponse](../../models/operations/getsnippetdynamicresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListSnippets

List all snippets for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Snippet.ListSnippets(ctx, operations.ListSnippetsRequest{
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListSnippetsRequest](../../models/operations/listsnippetsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListSnippetsResponse](../../models/operations/listsnippetsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateSnippetDynamic

Update a dynamic snippet for a particular service.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Snippet.UpdateSnippetDynamic(ctx, operations.UpdateSnippetDynamicRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        Snippet: &components.Snippet{
            Name: fastlytestgo.String("test-snippet"),
            Priority: fastlytestgo.String("10"),
        },
        SnippetID: "62Yd1WfiCBPENLloXfXmlO",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateSnippetDynamicRequest](../../models/operations/updatesnippetdynamicrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.UpdateSnippetDynamicResponse](../../models/operations/updatesnippetdynamicresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
