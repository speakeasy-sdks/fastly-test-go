# LoggingS3
(*LoggingS3*)

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
    res, err := s.LoggingS3.CreateLogAwsS3(ctx, operations.CreateLogAwsS3Request{
        LoggingS3Input: &shared.LoggingS3Input{
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: shared.LoggingS3FormatVersionTwo.ToPointer(),
            GzipLevel: fastly.Int64(0),
            MessageType: shared.LoggingS3MessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Period: fastly.Int64(3600),
            Placement: shared.LoggingS3PlacementWafDebug.ToPointer(),
            PublicKey: fastly.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: fastly.String("null"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateLogAwsS3Request](../../models/operations/createlogawss3request.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
    res, err := s.LoggingS3.DeleteLogAwsS3(ctx, operations.DeleteLogAwsS3Request{
        LoggingS3Name: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteLogAwsS3Request](../../models/operations/deletelogawss3request.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
    res, err := s.LoggingS3.GetLogAwsS3(ctx, operations.GetLogAwsS3Request{
        LoggingS3Name: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetLogAwsS3Request](../../models/operations/getlogawss3request.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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
    res, err := s.LoggingS3.ListLogAwsS3(ctx, operations.ListLogAwsS3Request{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListLogAwsS3Request](../../models/operations/listlogawss3request.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


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
    res, err := s.LoggingS3.UpdateLogAwsS3(ctx, operations.UpdateLogAwsS3Request{
        LoggingS3Input: &shared.LoggingS3Input{
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: shared.LoggingS3FormatVersionTwo.ToPointer(),
            GzipLevel: fastly.Int64(0),
            MessageType: shared.LoggingS3MessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Period: fastly.Int64(3600),
            Placement: shared.LoggingS3PlacementWafDebug.ToPointer(),
            PublicKey: fastly.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: fastly.String("null"),
        },
        LoggingS3Name: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateLogAwsS3Request](../../models/operations/updatelogawss3request.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateLogAwsS3Response](../../models/operations/updatelogawss3response.md), error**

