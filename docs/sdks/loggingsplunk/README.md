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
    operationSecurity := operations.CreateLogSplunkSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSplunk.CreateLogSplunk(ctx, operations.CreateLogSplunkRequest{
        LoggingSplunk2: &shared.LoggingSplunk2{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSplunkFormatVersionTwo.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingSplunkPlacementWafDebug.ToPointer(),
            RequestMaxBytes: sdk.Int64(277596),
            RequestMaxEntries: sdk.Int64(539224),
            ResponseCondition: sdk.String("null"),
            TLSCaCert: sdk.String("explicabo"),
            TLSClientCert: sdk.String("minima"),
            TLSClientKey: sdk.String("nisi"),
            TLSHostname: sdk.String("fugit"),
            Token: sdk.String("sapiente"),
            URL: sdk.String("consequuntur"),
            UseTLS: shared.LoggingUseTLSZero.ToPointer(),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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
    operationSecurity := operations.DeleteLogSplunkSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSplunk.DeleteLogSplunk(ctx, operations.DeleteLogSplunkRequest{
        LoggingSplunkName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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
    operationSecurity := operations.GetLogSplunkSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSplunk.GetLogSplunk(ctx, operations.GetLogSplunkRequest{
        LoggingSplunkName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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
    operationSecurity := operations.ListLogSplunkSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSplunk.ListLogSplunk(ctx, operations.ListLogSplunkRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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
    operationSecurity := operations.UpdateLogSplunkSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSplunk.UpdateLogSplunk(ctx, operations.UpdateLogSplunkRequest{
        LoggingSplunk2: &shared.LoggingSplunk2{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSplunkFormatVersionOne.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingSplunkPlacementLessThanNilGreaterThan.ToPointer(),
            RequestMaxBytes: sdk.Int64(578922),
            RequestMaxEntries: sdk.Int64(543806),
            ResponseCondition: sdk.String("null"),
            TLSCaCert: sdk.String("et"),
            TLSClientCert: sdk.String("esse"),
            TLSClientKey: sdk.String("eveniet"),
            TLSHostname: sdk.String("accusamus"),
            Token: sdk.String("veritatis"),
            URL: sdk.String("esse"),
            UseTLS: shared.LoggingUseTLSOne.ToPointer(),
        },
        LoggingSplunkName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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

