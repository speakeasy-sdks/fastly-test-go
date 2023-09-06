# LoggingCloudfiles

## Overview

Fastly will upload log messages to your Rackspace Cloud Files account.

<https://developer.fastly.com/reference/api/logging/cloudfiles>
### Available Operations

* [CreateLogCloudfiles](#createlogcloudfiles) - Create a Cloud Files log endpoint
* [DeleteLogCloudfiles](#deletelogcloudfiles) - Delete the Cloud Files log endpoint
* [GetLogCloudfiles](#getlogcloudfiles) - Get a Cloud Files log endpoint
* [ListLogCloudfiles](#listlogcloudfiles) - List Cloud Files log endpoints
* [UpdateLogCloudfiles](#updatelogcloudfiles) - Update the Cloud Files log endpoint

## CreateLogCloudfiles

Create a Cloud Files log endpoint for a particular service and version.

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
    operationSecurity := operations.CreateLogCloudfilesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingCloudfiles.CreateLogCloudfiles(ctx, operations.CreateLogCloudfilesRequest{
        LoggingCloudfilesInput: &shared.LoggingCloudfilesInput{
            AccessKey: sdk.String("iure"),
            BucketName: sdk.String("odio"),
            CompressionCodec: shared.LoggingCloudfilesCompressionCodecZstd.ToPointer(),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingCloudfilesFormatVersionTwo.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingCloudfilesMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("quidem"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingCloudfilesPlacementLessThanNilGreaterThan.ToPointer(),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            Region: shared.LoggingCloudfilesRegionIad.ToPointer(),
            ResponseCondition: sdk.String("null"),
            User: sdk.String("natus"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingCloudfilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateLogCloudfilesRequest](../../models/operations/createlogcloudfilesrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.CreateLogCloudfilesSecurity](../../models/operations/createlogcloudfilessecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.CreateLogCloudfilesResponse](../../models/operations/createlogcloudfilesresponse.md), error**


## DeleteLogCloudfiles

Delete the Cloud Files log endpoint for a particular service and version.

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
    operationSecurity := operations.DeleteLogCloudfilesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingCloudfiles.DeleteLogCloudfiles(ctx, operations.DeleteLogCloudfilesRequest{
        LoggingCloudfilesName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogCloudfiles200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteLogCloudfilesRequest](../../models/operations/deletelogcloudfilesrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.DeleteLogCloudfilesSecurity](../../models/operations/deletelogcloudfilessecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.DeleteLogCloudfilesResponse](../../models/operations/deletelogcloudfilesresponse.md), error**


## GetLogCloudfiles

Get the Cloud Files log endpoint for a particular service and version.

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
    operationSecurity := operations.GetLogCloudfilesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingCloudfiles.GetLogCloudfiles(ctx, operations.GetLogCloudfilesRequest{
        LoggingCloudfilesName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingCloudfilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetLogCloudfilesRequest](../../models/operations/getlogcloudfilesrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.GetLogCloudfilesSecurity](../../models/operations/getlogcloudfilessecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.GetLogCloudfilesResponse](../../models/operations/getlogcloudfilesresponse.md), error**


## ListLogCloudfiles

List all of the Cloud Files log endpoints for a particular service and version.

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
    operationSecurity := operations.ListLogCloudfilesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingCloudfiles.ListLogCloudfiles(ctx, operations.ListLogCloudfilesRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingCloudfilesResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListLogCloudfilesRequest](../../models/operations/listlogcloudfilesrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ListLogCloudfilesSecurity](../../models/operations/listlogcloudfilessecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.ListLogCloudfilesResponse](../../models/operations/listlogcloudfilesresponse.md), error**


## UpdateLogCloudfiles

Update the Cloud Files log endpoint for a particular service and version.

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
    operationSecurity := operations.UpdateLogCloudfilesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingCloudfiles.UpdateLogCloudfiles(ctx, operations.UpdateLogCloudfilesRequest{
        LoggingCloudfilesInput: &shared.LoggingCloudfilesInput{
            AccessKey: sdk.String("eos"),
            BucketName: sdk.String("atque"),
            CompressionCodec: shared.LoggingCloudfilesCompressionCodecZstd.ToPointer(),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingCloudfilesFormatVersionTwo.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingCloudfilesMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("ab"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingCloudfilesPlacementLessThanNilGreaterThan.ToPointer(),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            Region: shared.LoggingCloudfilesRegionSyd.ToPointer(),
            ResponseCondition: sdk.String("null"),
            User: sdk.String("iusto"),
        },
        LoggingCloudfilesName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingCloudfilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateLogCloudfilesRequest](../../models/operations/updatelogcloudfilesrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.UpdateLogCloudfilesSecurity](../../models/operations/updatelogcloudfilessecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.UpdateLogCloudfilesResponse](../../models/operations/updatelogcloudfilesresponse.md), error**

