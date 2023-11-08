# LoggingPubsub
(*.LoggingPubsub*)

## Overview

Fastly will publish log messages to a Google Cloud Pub/Sub topic in the format specified in the Pub/Sub logging object.

<https://developer.fastly.com/reference/api/logging/gcp-pubsub>
### Available Operations

* [CreateLogGcpPubsub](#createloggcppubsub) - Create a GCP Cloud Pub/Sub log endpoint
* [DeleteLogGcpPubsub](#deleteloggcppubsub) - Delete a GCP Cloud Pub/Sub log endpoint
* [GetLogGcpPubsub](#getloggcppubsub) - Get a GCP Cloud Pub/Sub log endpoint
* [ListLogGcpPubsub](#listloggcppubsub) - List GCP Cloud Pub/Sub log endpoints
* [UpdateLogGcpPubsub](#updateloggcppubsub) - Update a GCP Cloud Pub/Sub log endpoint

## CreateLogGcpPubsub

Create a Pub/Sub logging object for a particular service and version.

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
    res, err := s.LoggingPubsub.CreateLogGcpPubsub(ctx, operations.CreateLogGcpPubsubRequest{
        LoggingGooglePubsub: &components.LoggingGooglePubsub{
            AccountName: fastly.String("test-user@test-project-id.iam.gserviceaccount.com"),
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: components.LoggingGooglePubsubFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: components.LoggingGooglePubsubPlacementNone.ToPointer(),
            ProjectID: fastly.String("test-project-id"),
            ResponseCondition: fastly.String("null"),
            SecretKey: fastly.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            User: fastly.String("test-user@test-project-id.iam.gserviceaccount.com"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingGooglePubsubResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateLogGcpPubsubRequest](../../models/operations/createloggcppubsubrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CreateLogGcpPubsubResponse](../../models/operations/createloggcppubsubresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteLogGcpPubsub

Delete a Pub/Sub logging object for a particular service and version.

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
    res, err := s.LoggingPubsub.DeleteLogGcpPubsub(ctx, operations.DeleteLogGcpPubsubRequest{
        LoggingGooglePubsubName: "test-log-endpoint",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteLogGcpPubsubRequest](../../models/operations/deleteloggcppubsubrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.DeleteLogGcpPubsubResponse](../../models/operations/deleteloggcppubsubresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetLogGcpPubsub

Get the details for a Pub/Sub logging object for a particular service and version.

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
    res, err := s.LoggingPubsub.GetLogGcpPubsub(ctx, operations.GetLogGcpPubsubRequest{
        LoggingGooglePubsubName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingGooglePubsubResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetLogGcpPubsubRequest](../../models/operations/getloggcppubsubrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetLogGcpPubsubResponse](../../models/operations/getloggcppubsubresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListLogGcpPubsub

List all of the Pub/Sub logging objects for a particular service and version.

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
    res, err := s.LoggingPubsub.ListLogGcpPubsub(ctx, operations.ListLogGcpPubsubRequest{
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListLogGcpPubsubRequest](../../models/operations/listloggcppubsubrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListLogGcpPubsubResponse](../../models/operations/listloggcppubsubresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateLogGcpPubsub

Update a Pub/Sub logging object for a particular service and version.

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
    res, err := s.LoggingPubsub.UpdateLogGcpPubsub(ctx, operations.UpdateLogGcpPubsubRequest{
        LoggingGooglePubsub: &components.LoggingGooglePubsub{
            AccountName: fastly.String("test-user@test-project-id.iam.gserviceaccount.com"),
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: components.LoggingGooglePubsubFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: components.LoggingGooglePubsubPlacementNone.ToPointer(),
            ProjectID: fastly.String("test-project-id"),
            ResponseCondition: fastly.String("null"),
            SecretKey: fastly.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            User: fastly.String("test-user@test-project-id.iam.gserviceaccount.com"),
        },
        LoggingGooglePubsubName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingGooglePubsubResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateLogGcpPubsubRequest](../../models/operations/updateloggcppubsubrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateLogGcpPubsubResponse](../../models/operations/updateloggcppubsubresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
