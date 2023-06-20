# LoggingPapertrail

## Overview

Fastly will stream log messages to the Papertrail account in the format specified in the Papertrail object.

<https://developer.fastly.com/reference/api/logging/papertrail>
### Available Operations

* [CreateLogPapertrail](#createlogpapertrail) - Create a Papertrail log endpoint
* [DeleteLogPapertrail](#deletelogpapertrail) - Delete a Papertrail log endpoint
* [GetLogPapertrail](#getlogpapertrail) - Get a Papertrail log endpoint
* [ListLogPapertrail](#listlogpapertrail) - List Papertrail log endpoints
* [UpdateLogPapertrail](#updatelogpapertrail) - Update a Papertrail log endpoint

## CreateLogPapertrail

Create a Papertrail for a particular service and version.

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
    res, err := s.LoggingPapertrail.CreateLogPapertrail(ctx, operations.CreateLogPapertrailRequest{
        LoggingPapertrail2: &shared.LoggingPapertrail2{
            Address: sdk.String("example.com"),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingPapertrailFormatVersionTwo.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingPapertrailPlacementNone.ToPointer(),
            Port: sdk.Int64(869489),
            ResponseCondition: sdk.String("null"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateLogPapertrailSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingPapertrailResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateLogPapertrailRequest](../../models/operations/createlogpapertrailrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.CreateLogPapertrailSecurity](../../models/operations/createlogpapertrailsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.CreateLogPapertrailResponse](../../models/operations/createlogpapertrailresponse.md), error**


## DeleteLogPapertrail

Delete the Papertrail for a particular service and version.

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
    res, err := s.LoggingPapertrail.DeleteLogPapertrail(ctx, operations.DeleteLogPapertrailRequest{
        LoggingPapertrailName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteLogPapertrailSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogPapertrail200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteLogPapertrailRequest](../../models/operations/deletelogpapertrailrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.DeleteLogPapertrailSecurity](../../models/operations/deletelogpapertrailsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.DeleteLogPapertrailResponse](../../models/operations/deletelogpapertrailresponse.md), error**


## GetLogPapertrail

Get the Papertrail for a particular service and version.

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
    res, err := s.LoggingPapertrail.GetLogPapertrail(ctx, operations.GetLogPapertrailRequest{
        LoggingPapertrailName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetLogPapertrailSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingPapertrailResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetLogPapertrailRequest](../../models/operations/getlogpapertrailrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.GetLogPapertrailSecurity](../../models/operations/getlogpapertrailsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.GetLogPapertrailResponse](../../models/operations/getlogpapertrailresponse.md), error**


## ListLogPapertrail

List all of the Papertrails for a particular service and version.

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
    res, err := s.LoggingPapertrail.ListLogPapertrail(ctx, operations.ListLogPapertrailRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListLogPapertrailSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingPapertrailResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListLogPapertrailRequest](../../models/operations/listlogpapertrailrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ListLogPapertrailSecurity](../../models/operations/listlogpapertrailsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.ListLogPapertrailResponse](../../models/operations/listlogpapertrailresponse.md), error**


## UpdateLogPapertrail

Update the Papertrail for a particular service and version.

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
    res, err := s.LoggingPapertrail.UpdateLogPapertrail(ctx, operations.UpdateLogPapertrailRequest{
        LoggingPapertrail2: &shared.LoggingPapertrail2{
            Address: sdk.String("example.com"),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingPapertrailFormatVersionOne.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingPapertrailPlacementWafDebug.ToPointer(),
            Port: sdk.Int64(55965),
            ResponseCondition: sdk.String("null"),
        },
        LoggingPapertrailName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.UpdateLogPapertrailSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingPapertrailResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateLogPapertrailRequest](../../models/operations/updatelogpapertrailrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.UpdateLogPapertrailSecurity](../../models/operations/updatelogpapertrailsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.UpdateLogPapertrailResponse](../../models/operations/updatelogpapertrailresponse.md), error**
