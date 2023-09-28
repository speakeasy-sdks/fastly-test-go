# LoggingOpenstack
(*LoggingOpenstack*)

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
    res, err := s.LoggingOpenstack.CreateLogOpenstack(ctx, operations.CreateLogOpenstackRequest{
        LoggingOpenstackInput: &shared.LoggingOpenstackInput{
            AccessKey: fastly.String("voluptas"),
            BucketName: fastly.String("libero"),
            CompressionCodec: shared.LoggingOpenstackCompressionCodecZstd.ToPointer(),
            Format: fastly.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingOpenstackFormatVersionOne.ToPointer(),
            GzipLevel: fastly.Int64(0),
            MessageType: shared.LoggingOpenstackMessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Path: fastly.String("numquam"),
            Period: fastly.Int64(3600),
            Placement: shared.LoggingOpenstackPlacementNone.ToPointer(),
            PublicKey: fastly.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: fastly.String("null"),
            URL: fastly.String("provident"),
            User: fastly.String("ipsa"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateLogOpenstackRequest](../../models/operations/createlogopenstackrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


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
    res, err := s.LoggingOpenstack.DeleteLogOpenstack(ctx, operations.DeleteLogOpenstackRequest{
        LoggingOpenstackName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteLogOpenstackRequest](../../models/operations/deletelogopenstackrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


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
    res, err := s.LoggingOpenstack.GetLogOpenstack(ctx, operations.GetLogOpenstackRequest{
        LoggingOpenstackName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetLogOpenstackRequest](../../models/operations/getlogopenstackrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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
    res, err := s.LoggingOpenstack.ListLogOpenstack(ctx, operations.ListLogOpenstackRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListLogOpenstackRequest](../../models/operations/listlogopenstackrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
    res, err := s.LoggingOpenstack.UpdateLogOpenstack(ctx, operations.UpdateLogOpenstackRequest{
        LoggingOpenstackInput: &shared.LoggingOpenstackInput{
            AccessKey: fastly.String("molestiae"),
            BucketName: fastly.String("magnam"),
            CompressionCodec: shared.LoggingOpenstackCompressionCodecSnappy.ToPointer(),
            Format: fastly.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingOpenstackFormatVersionOne.ToPointer(),
            GzipLevel: fastly.Int64(0),
            MessageType: shared.LoggingOpenstackMessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Path: fastly.String("esse"),
            Period: fastly.Int64(3600),
            Placement: shared.LoggingOpenstackPlacementWafDebug.ToPointer(),
            PublicKey: fastly.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: fastly.String("null"),
            URL: fastly.String("rem"),
            User: fastly.String("fuga"),
        },
        LoggingOpenstackName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateLogOpenstackRequest](../../models/operations/updatelogopenstackrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateLogOpenstackResponse](../../models/operations/updatelogopenstackresponse.md), error**

