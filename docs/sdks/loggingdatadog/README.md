# LoggingDatadog

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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.CreateLogDatadogSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingDatadog.CreateLogDatadog(ctx, operations.CreateLogDatadogRequest{
        LoggingDatadog3: &shared.LoggingDatadog3{
            Format: sdk.String("voluptate"),
            FormatVersion: shared.LoggingDatadogFormatVersionTwo.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingDatadogPlacementWafDebug.ToPointer(),
            Region: shared.LoggingDatadogRegionEu.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("necessitatibus"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDatadogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateLogDatadogRequest](../../models/operations/createlogdatadogrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.CreateLogDatadogSecurity](../../models/operations/createlogdatadogsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.DeleteLogDatadogSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingDatadog.DeleteLogDatadog(ctx, operations.DeleteLogDatadogRequest{
        LoggingDatadogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogDatadog200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteLogDatadogRequest](../../models/operations/deletelogdatadogrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.DeleteLogDatadogSecurity](../../models/operations/deletelogdatadogsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.GetLogDatadogSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingDatadog.GetLogDatadog(ctx, operations.GetLogDatadogRequest{
        LoggingDatadogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDatadogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetLogDatadogRequest](../../models/operations/getlogdatadogrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.GetLogDatadogSecurity](../../models/operations/getlogdatadogsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListLogDatadogSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingDatadog.ListLogDatadog(ctx, operations.ListLogDatadogRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDatadogResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListLogDatadogRequest](../../models/operations/listlogdatadogrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.ListLogDatadogSecurity](../../models/operations/listlogdatadogsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.UpdateLogDatadogSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingDatadog.UpdateLogDatadog(ctx, operations.UpdateLogDatadogRequest{
        LoggingDatadog3: &shared.LoggingDatadog3{
            Format: sdk.String("distinctio"),
            FormatVersion: shared.LoggingDatadogFormatVersionTwo.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingDatadogPlacementWafDebug.ToPointer(),
            Region: shared.LoggingDatadogRegionUs.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("voluptate"),
        },
        LoggingDatadogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDatadogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateLogDatadogRequest](../../models/operations/updatelogdatadogrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.UpdateLogDatadogSecurity](../../models/operations/updatelogdatadogsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.UpdateLogDatadogResponse](../../models/operations/updatelogdatadogresponse.md), error**

