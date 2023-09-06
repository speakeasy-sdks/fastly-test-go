# LoggingDigitalocean

## Overview

Fastly will upload log messages to the DigitalOcean Space in the format specified in the DigitalOcean Spaces object.

<https://developer.fastly.com/reference/api/logging/digitalocean>
### Available Operations

* [CreateLogDigocean](#createlogdigocean) - Create a DigitalOcean Spaces log endpoint
* [DeleteLogDigocean](#deletelogdigocean) - Delete a DigitalOcean Spaces log endpoint
* [GetLogDigocean](#getlogdigocean) - Get a DigitalOcean Spaces log endpoint
* [ListLogDigocean](#listlogdigocean) - List DigitalOcean Spaces log endpoints
* [UpdateLogDigocean](#updatelogdigocean) - Update a DigitalOcean Spaces log endpoint

## CreateLogDigocean

Create a DigitalOcean Space for a particular service and version.

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
    operationSecurity := operations.CreateLogDigoceanSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingDigitalocean.CreateLogDigocean(ctx, operations.CreateLogDigoceanRequest{
        LoggingDigitaloceanInput: &shared.LoggingDigitaloceanInput{
            AccessKey: sdk.String("id"),
            BucketName: sdk.String("saepe"),
            CompressionCodec: shared.LoggingDigitaloceanCompressionCodecZstd.ToPointer(),
            Domain: sdk.String("aspernatur"),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingDigitaloceanFormatVersionOne.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingDigitaloceanMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("amet"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingDigitaloceanPlacementLessThanNilGreaterThan.ToPointer(),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            SecretKey: sdk.String("accusamus"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDigitaloceanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateLogDigoceanRequest](../../models/operations/createlogdigoceanrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.CreateLogDigoceanSecurity](../../models/operations/createlogdigoceansecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.CreateLogDigoceanResponse](../../models/operations/createlogdigoceanresponse.md), error**


## DeleteLogDigocean

Delete the DigitalOcean Space for a particular service and version.

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
    operationSecurity := operations.DeleteLogDigoceanSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingDigitalocean.DeleteLogDigocean(ctx, operations.DeleteLogDigoceanRequest{
        LoggingDigitaloceanName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogDigocean200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteLogDigoceanRequest](../../models/operations/deletelogdigoceanrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.DeleteLogDigoceanSecurity](../../models/operations/deletelogdigoceansecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.DeleteLogDigoceanResponse](../../models/operations/deletelogdigoceanresponse.md), error**


## GetLogDigocean

Get the DigitalOcean Space for a particular service and version.

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
    operationSecurity := operations.GetLogDigoceanSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingDigitalocean.GetLogDigocean(ctx, operations.GetLogDigoceanRequest{
        LoggingDigitaloceanName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDigitaloceanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetLogDigoceanRequest](../../models/operations/getlogdigoceanrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetLogDigoceanSecurity](../../models/operations/getlogdigoceansecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetLogDigoceanResponse](../../models/operations/getlogdigoceanresponse.md), error**


## ListLogDigocean

List all of the DigitalOcean Spaces for a particular service and version.

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
    operationSecurity := operations.ListLogDigoceanSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingDigitalocean.ListLogDigocean(ctx, operations.ListLogDigoceanRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDigitaloceanResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListLogDigoceanRequest](../../models/operations/listlogdigoceanrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ListLogDigoceanSecurity](../../models/operations/listlogdigoceansecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ListLogDigoceanResponse](../../models/operations/listlogdigoceanresponse.md), error**


## UpdateLogDigocean

Update the DigitalOcean Space for a particular service and version.

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
    operationSecurity := operations.UpdateLogDigoceanSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingDigitalocean.UpdateLogDigocean(ctx, operations.UpdateLogDigoceanRequest{
        LoggingDigitaloceanInput: &shared.LoggingDigitaloceanInput{
            AccessKey: sdk.String("ad"),
            BucketName: sdk.String("saepe"),
            CompressionCodec: shared.LoggingDigitaloceanCompressionCodecSnappy.ToPointer(),
            Domain: sdk.String("deserunt"),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingDigitaloceanFormatVersionTwo.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingDigitaloceanMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("minima"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingDigitaloceanPlacementLessThanNilGreaterThan.ToPointer(),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            SecretKey: sdk.String("totam"),
        },
        LoggingDigitaloceanName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDigitaloceanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateLogDigoceanRequest](../../models/operations/updatelogdigoceanrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UpdateLogDigoceanSecurity](../../models/operations/updatelogdigoceansecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.UpdateLogDigoceanResponse](../../models/operations/updatelogdigoceanresponse.md), error**

