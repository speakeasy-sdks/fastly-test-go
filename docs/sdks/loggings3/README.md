# LoggingS3

## Overview

Fastly will upload log messages to the S3 bucket in the format specified in the S3 object.

<https://developer.fastly.com/reference/api/logging/s3>
### Available Operations

* [CreateLogAwsS3](#createlogawss3) - Create an AWS S3 log endpoint
* [DeleteLogAwsS3](#deletelogawss3) - Delete an AWS S3 log endpoint
* [GetLogAwsS3](#getlogawss3) - Get an AWS S3 log endpoint
* [ListLogAwsS3](#listlogawss3) - List AWS S3 log endpoints
* [UpdateLogAwsS3](#updatelogawss3) - Update an AWS S3 log endpoint

## CreateLogAwsS3

Create a S3 for a particular service and version.

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
    res, err := s.LoggingS3.CreateLogAwsS3(ctx, operations.CreateLogAwsS3Request{
        LoggingS3Input: &shared.LoggingS3Input{
            AccessKey: sdk.String("accusantium"),
            ACL: sdk.String("rem"),
            BucketName: sdk.String("aut"),
            CompressionCodec: shared.LoggingS3CompressionCodecSnappy.ToPointer(),
            Domain: sdk.String("eum"),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingS3FormatVersionTwo.ToPointer(),
            GzipLevel: sdk.Int64(0),
            IamRole: sdk.String("ab"),
            MessageType: shared.LoggingS3MessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("corrupti"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingS3PlacementNone.ToPointer(),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            Redundancy: sdk.String("voluptatem"),
            ResponseCondition: sdk.String("null"),
            SecretKey: sdk.String("dolor"),
            ServerSideEncryption: sdk.String("occaecati"),
            ServerSideEncryptionKmsKeyID: sdk.String("numquam"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateLogAwsS3Security{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingS3Response != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateLogAwsS3Request](../../models/operations/createlogawss3request.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.CreateLogAwsS3Security](../../models/operations/createlogawss3security.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.CreateLogAwsS3Response](../../models/operations/createlogawss3response.md), error**


## DeleteLogAwsS3

Delete the S3 for a particular service and version.

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
    res, err := s.LoggingS3.DeleteLogAwsS3(ctx, operations.DeleteLogAwsS3Request{
        LoggingS3Name: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteLogAwsS3Security{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogAwsS3200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteLogAwsS3Request](../../models/operations/deletelogawss3request.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.DeleteLogAwsS3Security](../../models/operations/deletelogawss3security.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.DeleteLogAwsS3Response](../../models/operations/deletelogawss3response.md), error**


## GetLogAwsS3

Get the S3 for a particular service and version.

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
    res, err := s.LoggingS3.GetLogAwsS3(ctx, operations.GetLogAwsS3Request{
        LoggingS3Name: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetLogAwsS3Security{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingS3Response != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetLogAwsS3Request](../../models/operations/getlogawss3request.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.GetLogAwsS3Security](../../models/operations/getlogawss3security.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.GetLogAwsS3Response](../../models/operations/getlogawss3response.md), error**


## ListLogAwsS3

List all of the S3s for a particular service and version.

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
    res, err := s.LoggingS3.ListLogAwsS3(ctx, operations.ListLogAwsS3Request{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListLogAwsS3Security{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingS3Responses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListLogAwsS3Request](../../models/operations/listlogawss3request.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.ListLogAwsS3Security](../../models/operations/listlogawss3security.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.ListLogAwsS3Response](../../models/operations/listlogawss3response.md), error**


## UpdateLogAwsS3

Update the S3 for a particular service and version.

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
    res, err := s.LoggingS3.UpdateLogAwsS3(ctx, operations.UpdateLogAwsS3Request{
        LoggingS3Input: &shared.LoggingS3Input{
            AccessKey: sdk.String("impedit"),
            ACL: sdk.String("explicabo"),
            BucketName: sdk.String("voluptas"),
            CompressionCodec: shared.LoggingS3CompressionCodecZstd.ToPointer(),
            Domain: sdk.String("dignissimos"),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingS3FormatVersionOne.ToPointer(),
            GzipLevel: sdk.Int64(0),
            IamRole: sdk.String("maiores"),
            MessageType: shared.LoggingS3MessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("natus"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingS3PlacementNone.ToPointer(),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            Redundancy: sdk.String("voluptatibus"),
            ResponseCondition: sdk.String("null"),
            SecretKey: sdk.String("voluptas"),
            ServerSideEncryption: sdk.String("asperiores"),
            ServerSideEncryptionKmsKeyID: sdk.String("aperiam"),
        },
        LoggingS3Name: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.UpdateLogAwsS3Security{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingS3Response != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateLogAwsS3Request](../../models/operations/updatelogawss3request.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.UpdateLogAwsS3Security](../../models/operations/updatelogawss3security.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.UpdateLogAwsS3Response](../../models/operations/updatelogawss3response.md), error**

