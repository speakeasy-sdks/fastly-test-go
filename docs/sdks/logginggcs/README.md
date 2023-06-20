# LoggingGcs

## Overview

Fastly will upload log messages to the GCS bucket in the format specified in the GCS object.

<https://developer.fastly.com/reference/api/logging/gcs>
### Available Operations

* [CreateLogGcs](#createloggcs) - Create a GCS log endpoint
* [DeleteLogGcs](#deleteloggcs) - Delete a GCS log endpoint
* [GetLogGcs](#getloggcs) - Get a GCS log endpoint
* [ListLogGcs](#listloggcs) - List GCS log endpoints
* [UpdateLogGcs](#updateloggcs) - Update a GCS log endpoint

## CreateLogGcs

Create GCS logging for a particular service and version.

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
    res, err := s.LoggingGcs.CreateLogGcs(ctx, operations.CreateLogGcsRequest{
        LoggingGcsInput: &shared.LoggingGcsInput{
            AccountName: sdk.String("test-user@test-project-id.iam.gserviceaccount.com"),
            BucketName: sdk.String("ab"),
            CompressionCodec: shared.LoggingGcsCompressionCodecSnappy.ToPointer(),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingGcsFormatVersionOne.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingGcsMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("tempora"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingGcsPlacementLessThanNilGreaterThan.ToPointer(),
            ProjectID: sdk.String("test-project-id"),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            SecretKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            User: sdk.String("test-user@test-project-id.iam.gserviceaccount.com"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateLogGcsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingGcsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateLogGcsRequest](../../models/operations/createloggcsrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.CreateLogGcsSecurity](../../models/operations/createloggcssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.CreateLogGcsResponse](../../models/operations/createloggcsresponse.md), error**


## DeleteLogGcs

Delete the GCS Logging for a particular service and version.

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
    res, err := s.LoggingGcs.DeleteLogGcs(ctx, operations.DeleteLogGcsRequest{
        LoggingGcsName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteLogGcsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogGcs200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteLogGcsRequest](../../models/operations/deleteloggcsrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.DeleteLogGcsSecurity](../../models/operations/deleteloggcssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.DeleteLogGcsResponse](../../models/operations/deleteloggcsresponse.md), error**


## GetLogGcs

Get the GCS Logging for a particular service and version.

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
    res, err := s.LoggingGcs.GetLogGcs(ctx, operations.GetLogGcsRequest{
        LoggingGcsName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetLogGcsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingGcsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetLogGcsRequest](../../models/operations/getloggcsrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.GetLogGcsSecurity](../../models/operations/getloggcssecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.GetLogGcsResponse](../../models/operations/getloggcsresponse.md), error**


## ListLogGcs

List all of the GCS log endpoints for a particular service and version.

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
    res, err := s.LoggingGcs.ListLogGcs(ctx, operations.ListLogGcsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListLogGcsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingGcsResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListLogGcsRequest](../../models/operations/listloggcsrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.ListLogGcsSecurity](../../models/operations/listloggcssecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.ListLogGcsResponse](../../models/operations/listloggcsresponse.md), error**


## UpdateLogGcs

Update the GCS for a particular service and version.

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
    res, err := s.LoggingGcs.UpdateLogGcs(ctx, operations.UpdateLogGcsRequest{
        LoggingGcsInput: &shared.LoggingGcsInput{
            AccountName: sdk.String("test-user@test-project-id.iam.gserviceaccount.com"),
            BucketName: sdk.String("ipsam"),
            CompressionCodec: shared.LoggingGcsCompressionCodecZstd.ToPointer(),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingGcsFormatVersionOne.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingGcsMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("quo"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingGcsPlacementWafDebug.ToPointer(),
            ProjectID: sdk.String("test-project-id"),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            SecretKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            User: sdk.String("test-user@test-project-id.iam.gserviceaccount.com"),
        },
        LoggingGcsName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.UpdateLogGcsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingGcsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateLogGcsRequest](../../models/operations/updateloggcsrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.UpdateLogGcsSecurity](../../models/operations/updateloggcssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.UpdateLogGcsResponse](../../models/operations/updateloggcsresponse.md), error**

