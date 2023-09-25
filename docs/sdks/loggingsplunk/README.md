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
    res, err := s.LoggingSplunk.CreateLogSplunk(ctx, operations.CreateLogSplunkRequest{
        LoggingSplunk2: &shared.LoggingSplunk2{
            Format: fastly.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSplunkFormatVersionOne.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingSplunkPlacementNone.ToPointer(),
            RequestMaxBytes: fastly.Int64(129412),
            RequestMaxEntries: fastly.Int64(903984),
            ResponseCondition: fastly.String("null"),
            TLSCaCert: fastly.String("occaecati"),
            TLSClientCert: fastly.String("atque"),
            TLSClientKey: fastly.String("et"),
            TLSHostname: fastly.String("esse"),
            Token: fastly.String("eveniet"),
            URL: fastly.String("accusamus"),
            UseTLS: shared.LoggingUseTLSZero.ToPointer(),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateLogSplunkRequest](../../models/operations/createlogsplunkrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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
    res, err := s.LoggingSplunk.DeleteLogSplunk(ctx, operations.DeleteLogSplunkRequest{
        LoggingSplunkName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteLogSplunkRequest](../../models/operations/deletelogsplunkrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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
    res, err := s.LoggingSplunk.GetLogSplunk(ctx, operations.GetLogSplunkRequest{
        LoggingSplunkName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetLogSplunkRequest](../../models/operations/getlogsplunkrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


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
    res, err := s.LoggingSplunk.ListLogSplunk(ctx, operations.ListLogSplunkRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListLogSplunkRequest](../../models/operations/listlogsplunkrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


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
    res, err := s.LoggingSplunk.UpdateLogSplunk(ctx, operations.UpdateLogSplunkRequest{
        LoggingSplunk2: &shared.LoggingSplunk2{
            Format: fastly.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSplunkFormatVersionOne.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingSplunkPlacementLessThanNilGreaterThan.ToPointer(),
            RequestMaxBytes: fastly.Int64(724168),
            RequestMaxEntries: fastly.Int64(877131),
            ResponseCondition: fastly.String("null"),
            TLSCaCert: fastly.String("aliquid"),
            TLSClientCert: fastly.String("quasi"),
            TLSClientKey: fastly.String("saepe"),
            TLSHostname: fastly.String("vel"),
            Token: fastly.String("harum"),
            URL: fastly.String("molestiae"),
            UseTLS: shared.LoggingUseTLSOne.ToPointer(),
        },
        LoggingSplunkName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateLogSplunkRequest](../../models/operations/updatelogsplunkrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateLogSplunkResponse](../../models/operations/updatelogsplunkresponse.md), error**

