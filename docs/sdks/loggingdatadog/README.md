# LoggingDatadog
(*LoggingDatadog*)

## Overview

Fastly will upload log messages to Datadog in the format specified in the Datadog configuration object.

<https://developer.fastly.com/reference/api/logging/datadog>
### Available Operations

* [CreateLogDatadog](#createlogdatadog) - Create a Datadog log endpoint
* [DeleteLogDatadog](#deletelogdatadog) - Delete a Datadog log endpoint
* [GetLogDatadog](#getlogdatadog) - Get a Datadog log endpoint
* [ListLogDatadog](#listlogdatadog) - List Datadog log endpoints
* [UpdateLogDatadog](#updatelogdatadog) - Update a Datadog log endpoint

## CreateLogDatadog

Create a Datadog logging object for a particular service and version.

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
    res, err := s.LoggingDatadog.CreateLogDatadog(ctx, operations.CreateLogDatadogRequest{
        LoggingDatadog: &shared.LoggingDatadog{
            FormatVersion: shared.LoggingDatadogFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingDatadogPlacementNone.ToPointer(),
            ResponseCondition: fastly.String("null"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDatadogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateLogDatadogRequest](../../models/operations/createlogdatadogrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateLogDatadogResponse](../../models/operations/createlogdatadogresponse.md), error**


## DeleteLogDatadog

Delete the Datadog logging object for a particular service and version.

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
    res, err := s.LoggingDatadog.DeleteLogDatadog(ctx, operations.DeleteLogDatadogRequest{
        LoggingDatadogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogDatadog200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteLogDatadogRequest](../../models/operations/deletelogdatadogrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeleteLogDatadogResponse](../../models/operations/deletelogdatadogresponse.md), error**


## GetLogDatadog

Get the details for a Datadog logging object for a particular service and version.

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
    res, err := s.LoggingDatadog.GetLogDatadog(ctx, operations.GetLogDatadogRequest{
        LoggingDatadogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDatadogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetLogDatadogRequest](../../models/operations/getlogdatadogrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetLogDatadogResponse](../../models/operations/getlogdatadogresponse.md), error**


## ListLogDatadog

List all of the Datadog logging objects for a particular service and version.

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
    res, err := s.LoggingDatadog.ListLogDatadog(ctx, operations.ListLogDatadogRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDatadogResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListLogDatadogRequest](../../models/operations/listlogdatadogrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListLogDatadogResponse](../../models/operations/listlogdatadogresponse.md), error**


## UpdateLogDatadog

Update the Datadog logging object for a particular service and version.

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
    res, err := s.LoggingDatadog.UpdateLogDatadog(ctx, operations.UpdateLogDatadogRequest{
        LoggingDatadog: &shared.LoggingDatadog{
            FormatVersion: shared.LoggingDatadogFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingDatadogPlacementLessThanNilGreaterThan.ToPointer(),
            ResponseCondition: fastly.String("null"),
        },
        LoggingDatadogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDatadogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateLogDatadogRequest](../../models/operations/updatelogdatadogrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateLogDatadogResponse](../../models/operations/updatelogdatadogresponse.md), error**

