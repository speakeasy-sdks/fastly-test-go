# LoggingSyslog

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
    res, err := s.LoggingSyslog.CreateLogSyslog(ctx, operations.CreateLogSyslogRequest{
        LoggingSyslog2: &shared.LoggingSyslog2{
            Address: fastly.String("example.com"),
            Format: fastly.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSyslogFormatVersionTwo.ToPointer(),
            Hostname: fastly.String("equal-salute.com"),
            Ipv4: fastly.String("2.203.78.245"),
            MessageType: shared.LoggingMessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingSyslogPlacementNone.ToPointer(),
            Port: fastly.Int64(458139),
            ResponseCondition: fastly.String("null"),
            TLSCaCert: fastly.String("blanditiis"),
            TLSClientCert: fastly.String("provident"),
            TLSClientKey: fastly.String("a"),
            TLSHostname: fastly.String("nulla"),
            Token: fastly.String("quas"),
            UseTLS: shared.LoggingUseTLSZero.ToPointer(),
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


## DeleteLogSyslog

Delete the Syslog for a particular service and version.

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
    res, err := s.LoggingSyslog.DeleteLogSyslog(ctx, operations.DeleteLogSyslogRequest{
        LoggingSyslogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogSyslog200ApplicationJSONObject != nil {
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


## GetLogSyslog

Get the Syslog for a particular service and version.

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


## ListLogSyslog

List all of the Syslogs for a particular service and version.

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
    res, err := s.LoggingSyslog.ListLogSyslog(ctx, operations.ListLogSyslogRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSyslogResponses != nil {
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


## UpdateLogSyslog

Update the Syslog for a particular service and version.

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
    res, err := s.LoggingSyslog.UpdateLogSyslog(ctx, operations.UpdateLogSyslogRequest{
        LoggingSyslog2: &shared.LoggingSyslog2{
            Address: fastly.String("example.com"),
            Format: fastly.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSyslogFormatVersionOne.ToPointer(),
            Hostname: fastly.String("weekly-overexertion.name"),
            Ipv4: fastly.String("220.210.40.232"),
            MessageType: shared.LoggingMessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingSyslogPlacementLessThanNilGreaterThan.ToPointer(),
            Port: fastly.Int64(815524),
            ResponseCondition: fastly.String("null"),
            TLSCaCert: fastly.String("veritatis"),
            TLSClientCert: fastly.String("consequuntur"),
            TLSClientKey: fastly.String("quasi"),
            TLSHostname: fastly.String("similique"),
            Token: fastly.String("culpa"),
            UseTLS: shared.LoggingUseTLSZero.ToPointer(),
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

