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
    res, err := s.LoggingHoneycomb.CreateLogHoneycomb(ctx, operations.CreateLogHoneycombRequest{
        LoggingHoneycomb: &shared.LoggingHoneycomb{
            Dataset: fastly.String("septicaemia calculate"),
            Format: fastly.String("complexity Lock"),
            FormatVersion: shared.LoggingHoneycombFormatVersionOne.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingHoneycombPlacementLessThanNilGreaterThan.ToPointer(),
            ResponseCondition: fastly.String("null"),
            Token: fastly.String("Northeast Northwest"),
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


## DeleteLogHoneycomb

Delete the Honeycomb logging object for a particular service and version.

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
    res, err := s.LoggingHoneycomb.DeleteLogHoneycomb(ctx, operations.DeleteLogHoneycombRequest{
        LoggingHoneycombName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogHoneycomb200ApplicationJSONObject != nil {
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


## GetLogHoneycomb

Get the details of a Honeycomb logging object for a particular service and version.

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


## ListLogHoneycomb

List all of the Honeycomb logging objects for a particular service and version.

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
    res, err := s.LoggingHoneycomb.ListLogHoneycomb(ctx, operations.ListLogHoneycombRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHoneycombResponses != nil {
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


## UpdateLogHoneycomb

Update a Honeycomb logging object for a particular service and version.

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
    res, err := s.LoggingHoneycomb.UpdateLogHoneycomb(ctx, operations.UpdateLogHoneycombRequest{
        LoggingHoneycomb: &shared.LoggingHoneycomb{
            Dataset: fastly.String("Towels"),
            Format: fastly.String("mint Jamaican array"),
            FormatVersion: shared.LoggingHoneycombFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingHoneycombPlacementWafDebug.ToPointer(),
            ResponseCondition: fastly.String("null"),
            Token: fastly.String("probable male copying"),
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

