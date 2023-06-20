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

    ctx := context.Background()
    res, err := s.LoggingOpenstack.CreateLogOpenstack(ctx, operations.CreateLogOpenstackRequest{
        LoggingOpenstackInput: &shared.LoggingOpenstackInput{
            AccessKey: sdk.String("sapiente"),
            BucketName: sdk.String("dicta"),
            CompressionCodec: shared.LoggingOpenstackCompressionCodecSnappy.ToPointer(),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingOpenstackFormatVersionOne.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingOpenstackMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("ullam"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingOpenstackPlacementWafDebug.ToPointer(),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            URL: sdk.String("aut"),
            User: sdk.String("voluptatum"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateLogOpenstackSecurity{
        Token: "",
    })
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

    ctx := context.Background()
    res, err := s.LoggingOpenstack.DeleteLogOpenstack(ctx, operations.DeleteLogOpenstackRequest{
        LoggingOpenstackName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteLogOpenstackSecurity{
        Token: "",
    })
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

    ctx := context.Background()
    res, err := s.LoggingOpenstack.GetLogOpenstack(ctx, operations.GetLogOpenstackRequest{
        LoggingOpenstackName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetLogOpenstackSecurity{
        Token: "",
    })
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

    ctx := context.Background()
    res, err := s.LoggingOpenstack.ListLogOpenstack(ctx, operations.ListLogOpenstackRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListLogOpenstackSecurity{
        Token: "",
    })
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

    ctx := context.Background()
    res, err := s.LoggingOpenstack.UpdateLogOpenstack(ctx, operations.UpdateLogOpenstackRequest{
        LoggingOpenstackInput: &shared.LoggingOpenstackInput{
            AccessKey: sdk.String("qui"),
            BucketName: sdk.String("quibusdam"),
            CompressionCodec: shared.LoggingOpenstackCompressionCodecSnappy.ToPointer(),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingOpenstackFormatVersionTwo.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingOpenstackMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("itaque"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingOpenstackPlacementLessThanNilGreaterThan.ToPointer(),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            URL: sdk.String("architecto"),
            User: sdk.String("omnis"),
        },
        LoggingOpenstackName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.UpdateLogOpenstackSecurity{
        Token: "",
    })
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
