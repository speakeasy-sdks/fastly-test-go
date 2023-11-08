# LoggingSyslog
(*.LoggingSyslog*)

## Overview

Fastly will stream log messages to the location in the format specified in the Syslog object.

<https://developer.fastly.com/reference/api/logging/syslog>
### Available Operations

* [CreateLogSyslog](#createlogsyslog) - Create a syslog log endpoint
* [DeleteLogSyslog](#deletelogsyslog) - Delete a syslog log endpoint
* [GetLogSyslog](#getlogsyslog) - Get a syslog log endpoint
* [ListLogSyslog](#listlogsyslog) - List Syslog log endpoints
* [UpdateLogSyslog](#updatelogsyslog) - Update a syslog log endpoint

## CreateLogSyslog

Create a Syslog for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingSyslog.CreateLogSyslog(ctx, operations.CreateLogSyslogRequest{
        LoggingSyslog: &components.LoggingSyslog{
            Address: fastly.String("example.com"),
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: components.LoggingSyslogFormatVersionTwo.ToPointer(),
            MessageType: components.LoggingMessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: components.LoggingSyslogPlacementWafDebug.ToPointer(),
            ResponseCondition: fastly.String("null"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSyslogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateLogSyslogRequest](../../models/operations/createlogsyslogrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateLogSyslogResponse](../../models/operations/createlogsyslogresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteLogSyslog

Delete the Syslog for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingSyslog.DeleteLogSyslog(ctx, operations.DeleteLogSyslogRequest{
        LoggingSyslogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteLogSyslogRequest](../../models/operations/deletelogsyslogrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DeleteLogSyslogResponse](../../models/operations/deletelogsyslogresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetLogSyslog

Get the Syslog for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingSyslog.GetLogSyslog(ctx, operations.GetLogSyslogRequest{
        LoggingSyslogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSyslogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetLogSyslogRequest](../../models/operations/getlogsyslogrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetLogSyslogResponse](../../models/operations/getlogsyslogresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListLogSyslog

List all of the Syslogs for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingSyslog.ListLogSyslog(ctx, operations.ListLogSyslogRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListLogSyslogRequest](../../models/operations/listlogsyslogrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListLogSyslogResponse](../../models/operations/listlogsyslogresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateLogSyslog

Update the Syslog for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingSyslog.UpdateLogSyslog(ctx, operations.UpdateLogSyslogRequest{
        LoggingSyslog: &components.LoggingSyslog{
            Address: fastly.String("example.com"),
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: components.LoggingSyslogFormatVersionTwo.ToPointer(),
            MessageType: components.LoggingMessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: components.LoggingSyslogPlacementWafDebug.ToPointer(),
            ResponseCondition: fastly.String("null"),
        },
        LoggingSyslogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSyslogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateLogSyslogRequest](../../models/operations/updatelogsyslogrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateLogSyslogResponse](../../models/operations/updatelogsyslogresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
