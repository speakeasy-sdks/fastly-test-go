# LoggingScalyr

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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.CreateLogScalyrSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingScalyr.CreateLogScalyr(ctx, operations.CreateLogScalyrRequest{
        LoggingScalyr3: &shared.LoggingScalyr3{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingScalyrFormatVersionOne.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingScalyrPlacementLessThanNilGreaterThan.ToPointer(),
            ProjectID: sdk.String("esse"),
            Region: shared.LoggingScalyrRegionEu.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("aperiam"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingScalyrResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateLogScalyrRequest](../../models/operations/createlogscalyrrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.CreateLogScalyrSecurity](../../models/operations/createlogscalyrsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.DeleteLogScalyrSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingScalyr.DeleteLogScalyr(ctx, operations.DeleteLogScalyrRequest{
        LoggingScalyrName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogScalyr200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteLogScalyrRequest](../../models/operations/deletelogscalyrrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.DeleteLogScalyrSecurity](../../models/operations/deletelogscalyrsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.GetLogScalyrSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingScalyr.GetLogScalyr(ctx, operations.GetLogScalyrRequest{
        LoggingScalyrName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingScalyrResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetLogScalyrRequest](../../models/operations/getlogscalyrrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.GetLogScalyrSecurity](../../models/operations/getlogscalyrsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListLogScalyrSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingScalyr.ListLogScalyr(ctx, operations.ListLogScalyrRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingScalyrResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListLogScalyrRequest](../../models/operations/listlogscalyrrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListLogScalyrSecurity](../../models/operations/listlogscalyrsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.UpdateLogScalyrSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingScalyr.UpdateLogScalyr(ctx, operations.UpdateLogScalyrRequest{
        LoggingScalyr3: &shared.LoggingScalyr3{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingScalyrFormatVersionTwo.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingScalyrPlacementLessThanNilGreaterThan.ToPointer(),
            ProjectID: sdk.String("dignissimos"),
            Region: shared.LoggingScalyrRegionUs.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("nihil"),
        },
        LoggingScalyrName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingScalyrResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateLogScalyrRequest](../../models/operations/updatelogscalyrrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.UpdateLogScalyrSecurity](../../models/operations/updatelogscalyrsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.UpdateLogScalyrResponse](../../models/operations/updatelogscalyrresponse.md), error**

