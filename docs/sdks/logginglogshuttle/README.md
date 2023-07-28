# LoggingLogshuttle

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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.CreateLogLogshuttleSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingLogshuttle.CreateLogLogshuttle(ctx, operations.CreateLogLogshuttleRequest{
        LoggingLogshuttle2: &shared.LoggingLogshuttle2{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingLogshuttleFormatVersionTwo.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingLogshuttlePlacementNone.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("consequuntur"),
            URL: sdk.String("quasi"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogshuttleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateLogLogshuttleRequest](../../models/operations/createloglogshuttlerequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.CreateLogLogshuttleSecurity](../../models/operations/createloglogshuttlesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.CreateLogLogshuttleResponse](../../models/operations/createloglogshuttleresponse.md), error**


## DeleteLogLogshuttle

Delete the Log Shuttle logging endpoint for a particular service and version.

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
    operationSecurity := operations.DeleteLogLogshuttleSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingLogshuttle.DeleteLogLogshuttle(ctx, operations.DeleteLogLogshuttleRequest{
        LoggingLogshuttleName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogLogshuttle200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteLogLogshuttleRequest](../../models/operations/deleteloglogshuttlerequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.DeleteLogLogshuttleSecurity](../../models/operations/deleteloglogshuttlesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.DeleteLogLogshuttleResponse](../../models/operations/deleteloglogshuttleresponse.md), error**


## GetLogLogshuttle

Get the Log Shuttle logging endpoint for a particular service and version.

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
    operationSecurity := operations.GetLogLogshuttleSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingLogshuttle.GetLogLogshuttle(ctx, operations.GetLogLogshuttleRequest{
        LoggingLogshuttleName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogshuttleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetLogLogshuttleRequest](../../models/operations/getloglogshuttlerequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.GetLogLogshuttleSecurity](../../models/operations/getloglogshuttlesecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.GetLogLogshuttleResponse](../../models/operations/getloglogshuttleresponse.md), error**


## ListLogLogshuttle

List all of the Log Shuttle logging endpoints for a particular service and version.

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
    operationSecurity := operations.ListLogLogshuttleSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingLogshuttle.ListLogLogshuttle(ctx, operations.ListLogLogshuttleRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogshuttleResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListLogLogshuttleRequest](../../models/operations/listloglogshuttlerequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ListLogLogshuttleSecurity](../../models/operations/listloglogshuttlesecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.ListLogLogshuttleResponse](../../models/operations/listloglogshuttleresponse.md), error**


## UpdateLogLogshuttle

Update the Log Shuttle logging endpoint for a particular service and version.

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
    operationSecurity := operations.UpdateLogLogshuttleSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingLogshuttle.UpdateLogLogshuttle(ctx, operations.UpdateLogLogshuttleRequest{
        LoggingLogshuttle2: &shared.LoggingLogshuttle2{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingLogshuttleFormatVersionTwo.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingLogshuttlePlacementWafDebug.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("aliquid"),
            URL: sdk.String("tenetur"),
        },
        LoggingLogshuttleName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogshuttleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateLogLogshuttleRequest](../../models/operations/updateloglogshuttlerequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.UpdateLogLogshuttleSecurity](../../models/operations/updateloglogshuttlesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.UpdateLogLogshuttleResponse](../../models/operations/updateloglogshuttleresponse.md), error**

