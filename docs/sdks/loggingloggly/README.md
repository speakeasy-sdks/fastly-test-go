# LoggingLoggly

## Overview

Fastly will stream log messages to the Loggly account in the format specified in the Loggly logging object.

<https://developer.fastly.com/reference/api/logging/loggly>
### Available Operations

* [CreateLogLoggly](#createlogloggly) - Create a Loggly log endpoint
* [DeleteLogLoggly](#deletelogloggly) - Delete a Loggly log endpoint
* [GetLogLoggly](#getlogloggly) - Get a Loggly log endpoint
* [ListLogLoggly](#listlogloggly) - List Loggly log endpoints
* [UpdateLogLoggly](#updatelogloggly) - Update a Loggly log endpoint

## CreateLogLoggly

Create a Loggly logging object for a particular service and version.

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
    res, err := s.LoggingLoggly.CreateLogLoggly(ctx, operations.CreateLogLogglyRequest{
        LoggingLoggly2: &shared.LoggingLoggly2{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingLogglyFormatVersionTwo.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingLogglyPlacementLessThanNilGreaterThan.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("possimus"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateLogLogglySecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogglyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateLogLogglyRequest](../../models/operations/createloglogglyrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.CreateLogLogglySecurity](../../models/operations/createloglogglysecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.CreateLogLogglyResponse](../../models/operations/createloglogglyresponse.md), error**


## DeleteLogLoggly

Delete the Loggly logging object for a particular service and version.

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
    res, err := s.LoggingLoggly.DeleteLogLoggly(ctx, operations.DeleteLogLogglyRequest{
        LoggingLogglyName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteLogLogglySecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogLoggly200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteLogLogglyRequest](../../models/operations/deleteloglogglyrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.DeleteLogLogglySecurity](../../models/operations/deleteloglogglysecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.DeleteLogLogglyResponse](../../models/operations/deleteloglogglyresponse.md), error**


## GetLogLoggly

Get the Loggly logging object for a particular service and version.

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
    res, err := s.LoggingLoggly.GetLogLoggly(ctx, operations.GetLogLogglyRequest{
        LoggingLogglyName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetLogLogglySecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogglyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetLogLogglyRequest](../../models/operations/getloglogglyrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.GetLogLogglySecurity](../../models/operations/getloglogglysecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.GetLogLogglyResponse](../../models/operations/getloglogglyresponse.md), error**


## ListLogLoggly

List all Loggly logging objects for a particular service and version.

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
    res, err := s.LoggingLoggly.ListLogLoggly(ctx, operations.ListLogLogglyRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListLogLogglySecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogglyResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListLogLogglyRequest](../../models/operations/listloglogglyrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListLogLogglySecurity](../../models/operations/listloglogglysecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.ListLogLogglyResponse](../../models/operations/listloglogglyresponse.md), error**


## UpdateLogLoggly

Update the Loggly logging object for a particular service and version.

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
    res, err := s.LoggingLoggly.UpdateLogLoggly(ctx, operations.UpdateLogLogglyRequest{
        LoggingLoggly2: &shared.LoggingLoggly2{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingLogglyFormatVersionOne.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingLogglyPlacementLessThanNilGreaterThan.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("asperiores"),
        },
        LoggingLogglyName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.UpdateLogLogglySecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogglyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateLogLogglyRequest](../../models/operations/updateloglogglyrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.UpdateLogLogglySecurity](../../models/operations/updateloglogglysecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.UpdateLogLogglyResponse](../../models/operations/updateloglogglyresponse.md), error**

