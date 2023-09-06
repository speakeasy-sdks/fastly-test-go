# LoggingHoneycomb

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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.CreateLogHoneycombSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHoneycomb.CreateLogHoneycomb(ctx, operations.CreateLogHoneycombRequest{
        LoggingHoneycomb2: &shared.LoggingHoneycomb2{
            Dataset: sdk.String("doloribus"),
            Format: sdk.String("ut"),
            FormatVersion: shared.LoggingHoneycombFormatVersionTwo.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingHoneycombPlacementWafDebug.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("qui"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHoneycomb != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateLogHoneycombRequest](../../models/operations/createloghoneycombrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.CreateLogHoneycombSecurity](../../models/operations/createloghoneycombsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.DeleteLogHoneycombSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHoneycomb.DeleteLogHoneycomb(ctx, operations.DeleteLogHoneycombRequest{
        LoggingHoneycombName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogHoneycomb200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteLogHoneycombRequest](../../models/operations/deleteloghoneycombrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.DeleteLogHoneycombSecurity](../../models/operations/deleteloghoneycombsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.GetLogHoneycombSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHoneycomb.GetLogHoneycomb(ctx, operations.GetLogHoneycombRequest{
        LoggingHoneycombName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHoneycomb != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetLogHoneycombRequest](../../models/operations/getloghoneycombrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.GetLogHoneycombSecurity](../../models/operations/getloghoneycombsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListLogHoneycombSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHoneycomb.ListLogHoneycomb(ctx, operations.ListLogHoneycombRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHoneycombResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListLogHoneycombRequest](../../models/operations/listloghoneycombrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ListLogHoneycombSecurity](../../models/operations/listloghoneycombsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.UpdateLogHoneycombSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHoneycomb.UpdateLogHoneycomb(ctx, operations.UpdateLogHoneycombRequest{
        LoggingHoneycomb2: &shared.LoggingHoneycomb2{
            Dataset: sdk.String("quae"),
            Format: sdk.String("laudantium"),
            FormatVersion: shared.LoggingHoneycombFormatVersionOne.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingHoneycombPlacementWafDebug.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("voluptatibus"),
        },
        LoggingHoneycombName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHoneycombResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateLogHoneycombRequest](../../models/operations/updateloghoneycombrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.UpdateLogHoneycombSecurity](../../models/operations/updateloghoneycombsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.UpdateLogHoneycombResponse](../../models/operations/updateloghoneycombresponse.md), error**

