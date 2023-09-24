# LoggingFtp

## Overview

Fastly will upload log messages periodically to the server in the format specified in the FTP object.

<https://developer.fastly.com/reference/api/logging/ftp>
### Available Operations

* [CreateLogFtp](#createlogftp) - Create an FTP log endpoint
* [DeleteLogFtp](#deletelogftp) - Delete an FTP log endpoint
* [GetLogFtp](#getlogftp) - Get an FTP log endpoint
* [ListLogFtp](#listlogftp) - List FTP log endpoints
* [UpdateLogFtp](#updatelogftp) - Update an FTP log endpoint

## CreateLogFtp

Create a FTP for a particular service and version.

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
    res, err := s.LoggingFtp.CreateLogFtp(ctx, operations.CreateLogFtpRequest{
        LoggingFtpInput: &shared.LoggingFtpInput{
            Address: fastly.String("15887 Marguerite Manor"),
            CompressionCodec: shared.LoggingFtpCompressionCodecZstd.ToPointer(),
            Format: fastly.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingFtpFormatVersionOne.ToPointer(),
            GzipLevel: fastly.Int64(0),
            Hostname: fastly.String("decisive-radiosonde.net"),
            Ipv4: fastly.String("101.118.144.215"),
            MessageType: shared.LoggingFtpMessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Password: fastly.String("qui"),
            Path: fastly.String("neque"),
            Period: fastly.Int64(3600),
            Placement: shared.LoggingFtpPlacementNone.ToPointer(),
            Port: fastly.Int64(164959),
            PublicKey: fastly.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: fastly.String("null"),
            User: fastly.String("odio"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingFtpResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateLogFtpRequest](../../models/operations/createlogftprequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateLogFtpResponse](../../models/operations/createlogftpresponse.md), error**


## DeleteLogFtp

Delete the FTP for a particular service and version.

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
    res, err := s.LoggingFtp.DeleteLogFtp(ctx, operations.DeleteLogFtpRequest{
        LoggingFtpName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogFtp200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteLogFtpRequest](../../models/operations/deletelogftprequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DeleteLogFtpResponse](../../models/operations/deletelogftpresponse.md), error**


## GetLogFtp

Get the FTP for a particular service and version.

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
    res, err := s.LoggingFtp.GetLogFtp(ctx, operations.GetLogFtpRequest{
        LoggingFtpName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingFtpResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetLogFtpRequest](../../models/operations/getlogftprequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetLogFtpResponse](../../models/operations/getlogftpresponse.md), error**


## ListLogFtp

List all of the FTPs for a particular service and version.

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
    res, err := s.LoggingFtp.ListLogFtp(ctx, operations.ListLogFtpRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingFtpResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListLogFtpRequest](../../models/operations/listlogftprequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.ListLogFtpResponse](../../models/operations/listlogftpresponse.md), error**


## UpdateLogFtp

Update the FTP for a particular service and version.

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
    res, err := s.LoggingFtp.UpdateLogFtp(ctx, operations.UpdateLogFtpRequest{
        LoggingFtpInput: &shared.LoggingFtpInput{
            Address: fastly.String("37907 Rohan Cliffs"),
            CompressionCodec: shared.LoggingFtpCompressionCodecGzip.ToPointer(),
            Format: fastly.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingFtpFormatVersionOne.ToPointer(),
            GzipLevel: fastly.Int64(0),
            Hostname: fastly.String("bony-revolution.name"),
            Ipv4: fastly.String("187.149.11.246"),
            MessageType: shared.LoggingFtpMessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Password: fastly.String("dolorem"),
            Path: fastly.String("dolore"),
            Period: fastly.Int64(3600),
            Placement: shared.LoggingFtpPlacementNone.ToPointer(),
            Port: fastly.Int64(240829),
            PublicKey: fastly.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: fastly.String("null"),
            User: fastly.String("dolorum"),
        },
        LoggingFtpName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingFtpResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpdateLogFtpRequest](../../models/operations/updatelogftprequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.UpdateLogFtpResponse](../../models/operations/updatelogftpresponse.md), error**

