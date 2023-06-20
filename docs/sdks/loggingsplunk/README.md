# LoggingSplunk

## Overview

Fastly will POST messages to your Splunk account in the format specified in the Splunk object.

<https://developer.fastly.com/reference/api/logging/splunk>
### Available Operations

* [CreateLogSplunk](#createlogsplunk) - Create a Splunk log endpoint
* [DeleteLogSplunk](#deletelogsplunk) - Delete a Splunk log endpoint
* [GetLogSplunk](#getlogsplunk) - Get a Splunk log endpoint
* [ListLogSplunk](#listlogsplunk) - List Splunk log endpoints
* [UpdateLogSplunk](#updatelogsplunk) - Update a Splunk log endpoint

## CreateLogSplunk

Create a Splunk logging object for a particular service and version.

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
    res, err := s.LoggingSplunk.CreateLogSplunk(ctx, operations.CreateLogSplunkRequest{
        LoggingSplunk2: &shared.LoggingSplunk2{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSplunkFormatVersionTwo.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingSplunkPlacementWafDebug.ToPointer(),
            RequestMaxBytes: sdk.Int64(287051),
            RequestMaxEntries: sdk.Int64(822560),
            ResponseCondition: sdk.String("null"),
            TLSCaCert: sdk.String("facilis"),
            TLSClientCert: sdk.String("cum"),
            TLSClientKey: sdk.String("commodi"),
            TLSHostname: sdk.String("in"),
            Token: sdk.String("corporis"),
            URL: sdk.String("reiciendis"),
            UseTLS: shared.LoggingUseTLSOne.ToPointer(),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateLogSplunkSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSplunkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateLogSplunkRequest](../../models/operations/createlogsplunkrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.CreateLogSplunkSecurity](../../models/operations/createlogsplunksecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.CreateLogSplunkResponse](../../models/operations/createlogsplunkresponse.md), error**


## DeleteLogSplunk

Delete the Splunk logging object for a particular service and version.

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
    res, err := s.LoggingSplunk.DeleteLogSplunk(ctx, operations.DeleteLogSplunkRequest{
        LoggingSplunkName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteLogSplunkSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogSplunk200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteLogSplunkRequest](../../models/operations/deletelogsplunkrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.DeleteLogSplunkSecurity](../../models/operations/deletelogsplunksecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.DeleteLogSplunkResponse](../../models/operations/deletelogsplunkresponse.md), error**


## GetLogSplunk

Get the details for a Splunk logging object for a particular service and version.

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
    res, err := s.LoggingSplunk.GetLogSplunk(ctx, operations.GetLogSplunkRequest{
        LoggingSplunkName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetLogSplunkSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSplunkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetLogSplunkRequest](../../models/operations/getlogsplunkrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.GetLogSplunkSecurity](../../models/operations/getlogsplunksecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.GetLogSplunkResponse](../../models/operations/getlogsplunkresponse.md), error**


## ListLogSplunk

List all of the Splunk logging objects for a particular service and version.

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
    res, err := s.LoggingSplunk.ListLogSplunk(ctx, operations.ListLogSplunkRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListLogSplunkSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSplunkResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListLogSplunkRequest](../../models/operations/listlogsplunkrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListLogSplunkSecurity](../../models/operations/listlogsplunksecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.ListLogSplunkResponse](../../models/operations/listlogsplunkresponse.md), error**


## UpdateLogSplunk

Update the Splunk logging object for a particular service and version.

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
    res, err := s.LoggingSplunk.UpdateLogSplunk(ctx, operations.UpdateLogSplunkRequest{
        LoggingSplunk2: &shared.LoggingSplunk2{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSplunkFormatVersionOne.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingSplunkPlacementLessThanNilGreaterThan.ToPointer(),
            RequestMaxBytes: sdk.Int64(397533),
            RequestMaxEntries: sdk.Int64(46007),
            ResponseCondition: sdk.String("null"),
            TLSCaCert: sdk.String("cum"),
            TLSClientCert: sdk.String("consectetur"),
            TLSClientKey: sdk.String("in"),
            TLSHostname: sdk.String("exercitationem"),
            Token: sdk.String("earum"),
            URL: sdk.String("facere"),
            UseTLS: shared.LoggingUseTLSZero.ToPointer(),
        },
        LoggingSplunkName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.UpdateLogSplunkSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSplunkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateLogSplunkRequest](../../models/operations/updatelogsplunkrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.UpdateLogSplunkSecurity](../../models/operations/updatelogsplunksecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.UpdateLogSplunkResponse](../../models/operations/updatelogsplunkresponse.md), error**

