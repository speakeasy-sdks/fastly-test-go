# LoggingScalyr
(*LoggingScalyr*)

## Overview

Fastly will stream log messages to the Scalyr account in the format specified in the Scalyr object.

<https://developer.fastly.com/reference/api/logging/scalyr>
### Available Operations

* [CreateLogScalyr](#createlogscalyr) - Create a Scalyr log endpoint
* [DeleteLogScalyr](#deletelogscalyr) - Delete the Scalyr log endpoint
* [GetLogScalyr](#getlogscalyr) - Get a Scalyr log endpoint
* [ListLogScalyr](#listlogscalyr) - List Scalyr log endpoints
* [UpdateLogScalyr](#updatelogscalyr) - Update the Scalyr log endpoint

## CreateLogScalyr

Create a Scalyr for a particular service and version.

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
    res, err := s.LoggingScalyr.CreateLogScalyr(ctx, operations.CreateLogScalyrRequest{
        LoggingScalyr3: &shared.LoggingScalyr3{
            Format: fastly.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingScalyrFormatVersionOne.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingScalyrPlacementWafDebug.ToPointer(),
            ProjectID: fastly.String("accusamus"),
            Region: shared.LoggingScalyrRegionUs.ToPointer(),
            ResponseCondition: fastly.String("null"),
            Token: fastly.String("odio"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingScalyrResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateLogScalyrRequest](../../models/operations/createlogscalyrrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateLogScalyrResponse](../../models/operations/createlogscalyrresponse.md), error**


## DeleteLogScalyr

Delete the Scalyr for a particular service and version.

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
    res, err := s.LoggingScalyr.DeleteLogScalyr(ctx, operations.DeleteLogScalyrRequest{
        LoggingScalyrName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogScalyr200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteLogScalyrRequest](../../models/operations/deletelogscalyrrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DeleteLogScalyrResponse](../../models/operations/deletelogscalyrresponse.md), error**


## GetLogScalyr

Get the Scalyr for a particular service and version.

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
    res, err := s.LoggingScalyr.GetLogScalyr(ctx, operations.GetLogScalyrRequest{
        LoggingScalyrName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingScalyrResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetLogScalyrRequest](../../models/operations/getlogscalyrrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetLogScalyrResponse](../../models/operations/getlogscalyrresponse.md), error**


## ListLogScalyr

List all of the Scalyrs for a particular service and version.

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
    res, err := s.LoggingScalyr.ListLogScalyr(ctx, operations.ListLogScalyrRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingScalyrResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListLogScalyrRequest](../../models/operations/listlogscalyrrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListLogScalyrResponse](../../models/operations/listlogscalyrresponse.md), error**


## UpdateLogScalyr

Update the Scalyr for a particular service and version.

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
    res, err := s.LoggingScalyr.UpdateLogScalyr(ctx, operations.UpdateLogScalyrRequest{
        LoggingScalyr3: &shared.LoggingScalyr3{
            Format: fastly.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingScalyrFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingScalyrPlacementWafDebug.ToPointer(),
            ProjectID: fastly.String("sapiente"),
            Region: shared.LoggingScalyrRegionUs.ToPointer(),
            ResponseCondition: fastly.String("null"),
            Token: fastly.String("deserunt"),
        },
        LoggingScalyrName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingScalyrResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateLogScalyrRequest](../../models/operations/updatelogscalyrrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateLogScalyrResponse](../../models/operations/updatelogscalyrresponse.md), error**

