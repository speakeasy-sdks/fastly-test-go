# LoggingOpenstack

## Overview

Fastly will upload log messages to the OpenStack bucket in the format specified in the openstack object.

<https://developer.fastly.com/reference/api/logging/openstack>
### Available Operations

* [CreateLogOpenstack](#createlogopenstack) - Create an OpenStack log endpoint
* [DeleteLogOpenstack](#deletelogopenstack) - Delete an OpenStack log endpoint
* [GetLogOpenstack](#getlogopenstack) - Get an OpenStack log endpoint
* [ListLogOpenstack](#listlogopenstack) - List OpenStack log endpoints
* [UpdateLogOpenstack](#updatelogopenstack) - Update an OpenStack log endpoint

## CreateLogOpenstack

Create a openstack for a particular service and version.

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
    operationSecurity := operations.CreateLogOpenstackSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingOpenstack.CreateLogOpenstack(ctx, operations.CreateLogOpenstackRequest{
        LoggingOpenstackInput: &shared.LoggingOpenstackInput{
            AccessKey: sdk.String("laborum"),
            BucketName: sdk.String("placeat"),
            CompressionCodec: shared.LoggingOpenstackCompressionCodecZstd.ToPointer(),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingOpenstackFormatVersionOne.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingOpenstackMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("autem"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingOpenstackPlacementLessThanNilGreaterThan.ToPointer(),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            URL: sdk.String("quas"),
            User: sdk.String("assumenda"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingOpenstackResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateLogOpenstackRequest](../../models/operations/createlogopenstackrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.CreateLogOpenstackSecurity](../../models/operations/createlogopenstacksecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.CreateLogOpenstackResponse](../../models/operations/createlogopenstackresponse.md), error**


## DeleteLogOpenstack

Delete the openstack for a particular service and version.

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
    operationSecurity := operations.DeleteLogOpenstackSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingOpenstack.DeleteLogOpenstack(ctx, operations.DeleteLogOpenstackRequest{
        LoggingOpenstackName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogOpenstack200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteLogOpenstackRequest](../../models/operations/deletelogopenstackrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.DeleteLogOpenstackSecurity](../../models/operations/deletelogopenstacksecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.DeleteLogOpenstackResponse](../../models/operations/deletelogopenstackresponse.md), error**


## GetLogOpenstack

Get the openstack for a particular service and version.

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
    operationSecurity := operations.GetLogOpenstackSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingOpenstack.GetLogOpenstack(ctx, operations.GetLogOpenstackRequest{
        LoggingOpenstackName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingOpenstackResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetLogOpenstackRequest](../../models/operations/getlogopenstackrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.GetLogOpenstackSecurity](../../models/operations/getlogopenstacksecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.GetLogOpenstackResponse](../../models/operations/getlogopenstackresponse.md), error**


## ListLogOpenstack

List all of the openstacks for a particular service and version.

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
    operationSecurity := operations.ListLogOpenstackSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingOpenstack.ListLogOpenstack(ctx, operations.ListLogOpenstackRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingOpenstackResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListLogOpenstackRequest](../../models/operations/listlogopenstackrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ListLogOpenstackSecurity](../../models/operations/listlogopenstacksecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ListLogOpenstackResponse](../../models/operations/listlogopenstackresponse.md), error**


## UpdateLogOpenstack

Update the openstack for a particular service and version.

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
    operationSecurity := operations.UpdateLogOpenstackSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingOpenstack.UpdateLogOpenstack(ctx, operations.UpdateLogOpenstackRequest{
        LoggingOpenstackInput: &shared.LoggingOpenstackInput{
            AccessKey: sdk.String("nulla"),
            BucketName: sdk.String("voluptas"),
            CompressionCodec: shared.LoggingOpenstackCompressionCodecGzip.ToPointer(),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingOpenstackFormatVersionOne.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingOpenstackMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("tempora"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingOpenstackPlacementNone.ToPointer(),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            URL: sdk.String("explicabo"),
            User: sdk.String("provident"),
        },
        LoggingOpenstackName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingOpenstackResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateLogOpenstackRequest](../../models/operations/updatelogopenstackrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.UpdateLogOpenstackSecurity](../../models/operations/updatelogopenstacksecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.UpdateLogOpenstackResponse](../../models/operations/updatelogopenstackresponse.md), error**

