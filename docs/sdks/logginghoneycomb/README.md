# LoggingHoneycomb
(*LoggingHoneycomb*)

## Overview

Fastly will upload log messages to Honeycomb.io in the format specified in the Honeycomb object.

<https://developer.fastly.com/reference/api/logging/honeycomb>
### Available Operations

* [CreateLogHoneycomb](#createloghoneycomb) - Create a Honeycomb log endpoint
* [DeleteLogHoneycomb](#deleteloghoneycomb) - Delete the Honeycomb log endpoint
* [GetLogHoneycomb](#getloghoneycomb) - Get a Honeycomb log endpoint
* [ListLogHoneycomb](#listloghoneycomb) - List Honeycomb log endpoints
* [UpdateLogHoneycomb](#updateloghoneycomb) - Update a Honeycomb log endpoint

## CreateLogHoneycomb

Create a Honeycomb logging object for a particular service and version.

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
    res, err := s.LoggingHoneycomb.CreateLogHoneycomb(ctx, operations.CreateLogHoneycombRequest{
        LoggingHoneycomb: &components.LoggingHoneycomb{
            FormatVersion: components.LoggingHoneycombFormatVersionTwo.ToPointer(),
            Name: fastlytestgo.String("test-log-endpoint"),
            Placement: components.LoggingHoneycombPlacementWafDebug.ToPointer(),
            ResponseCondition: fastlytestgo.String("null"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHoneycomb != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateLogHoneycombRequest](../../models/operations/createloghoneycombrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CreateLogHoneycombResponse](../../models/operations/createloghoneycombresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteLogHoneycomb

Delete the Honeycomb logging object for a particular service and version.

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
    res, err := s.LoggingHoneycomb.DeleteLogHoneycomb(ctx, operations.DeleteLogHoneycombRequest{
        LoggingHoneycombName: "test-log-endpoint",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteLogHoneycombRequest](../../models/operations/deleteloghoneycombrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.DeleteLogHoneycombResponse](../../models/operations/deleteloghoneycombresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetLogHoneycomb

Get the details of a Honeycomb logging object for a particular service and version.

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
    res, err := s.LoggingHoneycomb.GetLogHoneycomb(ctx, operations.GetLogHoneycombRequest{
        LoggingHoneycombName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHoneycomb != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetLogHoneycombRequest](../../models/operations/getloghoneycombrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetLogHoneycombResponse](../../models/operations/getloghoneycombresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListLogHoneycomb

List all of the Honeycomb logging objects for a particular service and version.

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
    res, err := s.LoggingHoneycomb.ListLogHoneycomb(ctx, operations.ListLogHoneycombRequest{
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListLogHoneycombRequest](../../models/operations/listloghoneycombrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListLogHoneycombResponse](../../models/operations/listloghoneycombresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateLogHoneycomb

Update a Honeycomb logging object for a particular service and version.

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
    res, err := s.LoggingHoneycomb.UpdateLogHoneycomb(ctx, operations.UpdateLogHoneycombRequest{
        LoggingHoneycomb: &components.LoggingHoneycomb{
            FormatVersion: components.LoggingHoneycombFormatVersionTwo.ToPointer(),
            Name: fastlytestgo.String("test-log-endpoint"),
            Placement: components.LoggingHoneycombPlacementNone.ToPointer(),
            ResponseCondition: fastlytestgo.String("null"),
        },
        LoggingHoneycombName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHoneycombResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateLogHoneycombRequest](../../models/operations/updateloghoneycombrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateLogHoneycombResponse](../../models/operations/updateloghoneycombresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
