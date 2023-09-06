# LoggingHeroku

## Overview

Fastly will stream log messages to the Heroku account in the format specified in the Heroku object.

<https://developer.fastly.com/reference/api/logging/heroku>
### Available Operations

* [CreateLogHeroku](#createlogheroku) - Create a Heroku log endpoint
* [DeleteLogHeroku](#deletelogheroku) - Delete the Heroku log endpoint
* [GetLogHeroku](#getlogheroku) - Get a Heroku log endpoint
* [ListLogHeroku](#listlogheroku) - List Heroku log endpoints
* [UpdateLogHeroku](#updatelogheroku) - Update the Heroku log endpoint

## CreateLogHeroku

Create a Heroku for a particular service and version.

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
    operationSecurity := operations.CreateLogHerokuSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHeroku.CreateLogHeroku(ctx, operations.CreateLogHerokuRequest{
        LoggingHeroku2: &shared.LoggingHeroku2{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingHerokuFormatVersionOne.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingHerokuPlacementNone.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("quas"),
            URL: sdk.String("itaque"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHerokuResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateLogHerokuRequest](../../models/operations/createlogherokurequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.CreateLogHerokuSecurity](../../models/operations/createlogherokusecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.CreateLogHerokuResponse](../../models/operations/createlogherokuresponse.md), error**


## DeleteLogHeroku

Delete the Heroku for a particular service and version.

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
    operationSecurity := operations.DeleteLogHerokuSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHeroku.DeleteLogHeroku(ctx, operations.DeleteLogHerokuRequest{
        LoggingHerokuName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogHeroku200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteLogHerokuRequest](../../models/operations/deletelogherokurequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.DeleteLogHerokuSecurity](../../models/operations/deletelogherokusecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.DeleteLogHerokuResponse](../../models/operations/deletelogherokuresponse.md), error**


## GetLogHeroku

Get the Heroku for a particular service and version.

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
    operationSecurity := operations.GetLogHerokuSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHeroku.GetLogHeroku(ctx, operations.GetLogHerokuRequest{
        LoggingHerokuName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHerokuResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetLogHerokuRequest](../../models/operations/getlogherokurequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.GetLogHerokuSecurity](../../models/operations/getlogherokusecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.GetLogHerokuResponse](../../models/operations/getlogherokuresponse.md), error**


## ListLogHeroku

List all of the Herokus for a particular service and version.

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
    operationSecurity := operations.ListLogHerokuSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHeroku.ListLogHeroku(ctx, operations.ListLogHerokuRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHerokuResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListLogHerokuRequest](../../models/operations/listlogherokurequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListLogHerokuSecurity](../../models/operations/listlogherokusecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.ListLogHerokuResponse](../../models/operations/listlogherokuresponse.md), error**


## UpdateLogHeroku

Update the Heroku for a particular service and version.

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
    operationSecurity := operations.UpdateLogHerokuSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHeroku.UpdateLogHeroku(ctx, operations.UpdateLogHerokuRequest{
        LoggingHeroku2: &shared.LoggingHeroku2{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingHerokuFormatVersionOne.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingHerokuPlacementLessThanNilGreaterThan.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("repellendus"),
            URL: sdk.String("porro"),
        },
        LoggingHerokuName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHerokuResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateLogHerokuRequest](../../models/operations/updatelogherokurequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.UpdateLogHerokuSecurity](../../models/operations/updatelogherokusecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.UpdateLogHerokuResponse](../../models/operations/updatelogherokuresponse.md), error**

