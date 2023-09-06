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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.CreateLogSyslogSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSyslog.CreateLogSyslog(ctx, operations.CreateLogSyslogRequest{
        LoggingSyslog2: &shared.LoggingSyslog2{
            Address: sdk.String("example.com"),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSyslogFormatVersionTwo.ToPointer(),
            Hostname: sdk.String("kooky-propaganda.name"),
            Ipv4: sdk.String("83.183.193.6"),
            MessageType: shared.LoggingMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingSyslogPlacementWafDebug.ToPointer(),
            Port: sdk.Int64(731398),
            ResponseCondition: sdk.String("null"),
            TLSCaCert: sdk.String("adipisci"),
            TLSClientCert: sdk.String("cumque"),
            TLSClientKey: sdk.String("consequuntur"),
            TLSHostname: sdk.String("consequatur"),
            Token: sdk.String("minus"),
            UseTLS: shared.LoggingUseTLSZero.ToPointer(),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSyslogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateLogSyslogRequest](../../models/operations/createlogsyslogrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.CreateLogSyslogSecurity](../../models/operations/createlogsyslogsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.DeleteLogSyslogSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSyslog.DeleteLogSyslog(ctx, operations.DeleteLogSyslogRequest{
        LoggingSyslogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogSyslog200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteLogSyslogRequest](../../models/operations/deletelogsyslogrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.DeleteLogSyslogSecurity](../../models/operations/deletelogsyslogsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.GetLogSyslogSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSyslog.GetLogSyslog(ctx, operations.GetLogSyslogRequest{
        LoggingSyslogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSyslogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetLogSyslogRequest](../../models/operations/getlogsyslogrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.GetLogSyslogSecurity](../../models/operations/getlogsyslogsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListLogSyslogSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSyslog.ListLogSyslog(ctx, operations.ListLogSyslogRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSyslogResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListLogSyslogRequest](../../models/operations/listlogsyslogrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListLogSyslogSecurity](../../models/operations/listlogsyslogsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.UpdateLogSyslogSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSyslog.UpdateLogSyslog(ctx, operations.UpdateLogSyslogRequest{
        LoggingSyslog2: &shared.LoggingSyslog2{
            Address: sdk.String("example.com"),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSyslogFormatVersionTwo.ToPointer(),
            Hostname: sdk.String("emotional-infancy.name"),
            Ipv4: sdk.String("151.244.219.142"),
            MessageType: shared.LoggingMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingSyslogPlacementWafDebug.ToPointer(),
            Port: sdk.Int64(97468),
            ResponseCondition: sdk.String("null"),
            TLSCaCert: sdk.String("a"),
            TLSClientCert: sdk.String("error"),
            TLSClientKey: sdk.String("sint"),
            TLSHostname: sdk.String("pariatur"),
            Token: sdk.String("possimus"),
            UseTLS: shared.LoggingUseTLSZero.ToPointer(),
        },
        LoggingSyslogName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSyslogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateLogSyslogRequest](../../models/operations/updatelogsyslogrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.UpdateLogSyslogSecurity](../../models/operations/updatelogsyslogsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.UpdateLogSyslogResponse](../../models/operations/updatelogsyslogresponse.md), error**

