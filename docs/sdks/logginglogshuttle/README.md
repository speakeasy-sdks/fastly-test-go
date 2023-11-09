# LoggingLogshuttle
(*LoggingLogshuttle*)

## Overview

Fastly will upload log messages to the Log Shuttle bucket in the format specified in the logshuttle object.

<https://developer.fastly.com/reference/api/logging/logshuttle>
### Available Operations

* [CreateLogLogshuttle](#createloglogshuttle) - Create a Log Shuttle log endpoint
* [DeleteLogLogshuttle](#deleteloglogshuttle) - Delete a Log Shuttle log endpoint
* [GetLogLogshuttle](#getloglogshuttle) - Get a Log Shuttle log endpoint
* [ListLogLogshuttle](#listloglogshuttle) - List Log Shuttle log endpoints
* [UpdateLogLogshuttle](#updateloglogshuttle) - Update a Log Shuttle log endpoint

## CreateLogLogshuttle

Create a Log Shuttle logging endpoint for a particular service and version.

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
    res, err := s.LoggingLogshuttle.CreateLogLogshuttle(ctx, operations.CreateLogLogshuttleRequest{
        LoggingLogshuttle: &components.LoggingLogshuttle{
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: components.LoggingLogshuttleFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: components.LoggingLogshuttlePlacementNone.ToPointer(),
            ResponseCondition: fastly.String("null"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogshuttleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateLogLogshuttleRequest](../../models/operations/createloglogshuttlerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.CreateLogLogshuttleResponse](../../models/operations/createloglogshuttleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteLogLogshuttle

Delete the Log Shuttle logging endpoint for a particular service and version.

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
    res, err := s.LoggingLogshuttle.DeleteLogLogshuttle(ctx, operations.DeleteLogLogshuttleRequest{
        LoggingLogshuttleName: "test-log-endpoint",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteLogLogshuttleRequest](../../models/operations/deleteloglogshuttlerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteLogLogshuttleResponse](../../models/operations/deleteloglogshuttleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetLogLogshuttle

Get the Log Shuttle logging endpoint for a particular service and version.

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
    res, err := s.LoggingLogshuttle.GetLogLogshuttle(ctx, operations.GetLogLogshuttleRequest{
        LoggingLogshuttleName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogshuttleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetLogLogshuttleRequest](../../models/operations/getloglogshuttlerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetLogLogshuttleResponse](../../models/operations/getloglogshuttleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListLogLogshuttle

List all of the Log Shuttle logging endpoints for a particular service and version.

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
    res, err := s.LoggingLogshuttle.ListLogLogshuttle(ctx, operations.ListLogLogshuttleRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListLogLogshuttleRequest](../../models/operations/listloglogshuttlerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListLogLogshuttleResponse](../../models/operations/listloglogshuttleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateLogLogshuttle

Update the Log Shuttle logging endpoint for a particular service and version.

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
    res, err := s.LoggingLogshuttle.UpdateLogLogshuttle(ctx, operations.UpdateLogLogshuttleRequest{
        LoggingLogshuttle: &components.LoggingLogshuttle{
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: components.LoggingLogshuttleFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: components.LoggingLogshuttlePlacementNone.ToPointer(),
            ResponseCondition: fastly.String("null"),
        },
        LoggingLogshuttleName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogshuttleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateLogLogshuttleRequest](../../models/operations/updateloglogshuttlerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdateLogLogshuttleResponse](../../models/operations/updateloglogshuttleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
