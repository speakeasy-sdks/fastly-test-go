# LoggingNewrelic

## Overview

Fastly will upload log messages to New Relic Logs in the format specified in the New Relic configuration object.

<https://developer.fastly.com/reference/api/logging/newrelic>
### Available Operations

* [CreateLogNewrelic](#createlognewrelic) - Create a New Relic log endpoint
* [DeleteLogNewrelic](#deletelognewrelic) - Delete a New Relic log endpoint
* [GetLogNewrelic](#getlognewrelic) - Get a New Relic log endpoint
* [ListLogNewrelic](#listlognewrelic) - List New Relic log endpoints
* [UpdateLogNewrelic](#updatelognewrelic) - Update a New Relic log endpoint

## CreateLogNewrelic

Create a New Relic Logs logging object for a particular service and version.

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
    res, err := s.LoggingNewrelic.CreateLogNewrelic(ctx, operations.CreateLogNewrelicRequest{
        LoggingNewrelic3: &shared.LoggingNewrelic3{
            Format: fastly.String("dolorum"),
            FormatVersion: shared.LoggingNewrelicFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingNewrelicPlacementLessThanNilGreaterThan.ToPointer(),
            Region: shared.LoggingNewrelicRegionUs.ToPointer(),
            ResponseCondition: fastly.String("null"),
            Token: fastly.String("eum"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingNewrelicResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateLogNewrelicRequest](../../models/operations/createlognewrelicrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateLogNewrelicResponse](../../models/operations/createlognewrelicresponse.md), error**


## DeleteLogNewrelic

Delete the New Relic Logs logging object for a particular service and version.

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
    res, err := s.LoggingNewrelic.DeleteLogNewrelic(ctx, operations.DeleteLogNewrelicRequest{
        LoggingNewrelicName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogNewrelic200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteLogNewrelicRequest](../../models/operations/deletelognewrelicrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteLogNewrelicResponse](../../models/operations/deletelognewrelicresponse.md), error**


## GetLogNewrelic

Get the details of a New Relic Logs logging object for a particular service and version.

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
    res, err := s.LoggingNewrelic.GetLogNewrelic(ctx, operations.GetLogNewrelicRequest{
        LoggingNewrelicName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingNewrelicResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetLogNewrelicRequest](../../models/operations/getlognewrelicrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetLogNewrelicResponse](../../models/operations/getlognewrelicresponse.md), error**


## ListLogNewrelic

List all of the New Relic Logs logging objects for a particular service and version.

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
    res, err := s.LoggingNewrelic.ListLogNewrelic(ctx, operations.ListLogNewrelicRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingNewrelicResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListLogNewrelicRequest](../../models/operations/listlognewrelicrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListLogNewrelicResponse](../../models/operations/listlognewrelicresponse.md), error**


## UpdateLogNewrelic

Update a New Relic Logs logging object for a particular service and version.

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
    res, err := s.LoggingNewrelic.UpdateLogNewrelic(ctx, operations.UpdateLogNewrelicRequest{
        LoggingNewrelic3: &shared.LoggingNewrelic3{
            Format: fastly.String("autem"),
            FormatVersion: shared.LoggingNewrelicFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingNewrelicPlacementWafDebug.ToPointer(),
            Region: shared.LoggingNewrelicRegionEu.ToPointer(),
            ResponseCondition: fastly.String("null"),
            Token: fastly.String("nulla"),
        },
        LoggingNewrelicName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingNewrelicResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateLogNewrelicRequest](../../models/operations/updatelognewrelicrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateLogNewrelicResponse](../../models/operations/updatelognewrelicresponse.md), error**

