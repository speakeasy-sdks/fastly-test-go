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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.CreateLogFtpSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingFtp.CreateLogFtp(ctx, operations.CreateLogFtpRequest{
        LoggingFtpInput: &shared.LoggingFtpInput{
            Address: sdk.String("14736 Cristobal Forge"),
            CompressionCodec: shared.LoggingFtpCompressionCodecZstd.ToPointer(),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingFtpFormatVersionTwo.ToPointer(),
            GzipLevel: sdk.Int64(0),
            Hostname: sdk.String("steel-subconscious.net"),
            Ipv4: sdk.String("30.172.132.70"),
            MessageType: shared.LoggingFtpMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Password: sdk.String("aspernatur"),
            Path: sdk.String("dolores"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingFtpPlacementLessThanNilGreaterThan.ToPointer(),
            Port: sdk.Int64(704474),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            User: sdk.String("aliquid"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingFtpResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateLogFtpRequest](../../models/operations/createlogftprequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.CreateLogFtpSecurity](../../models/operations/createlogftpsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.DeleteLogFtpSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingFtp.DeleteLogFtp(ctx, operations.DeleteLogFtpRequest{
        LoggingFtpName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogFtp200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteLogFtpRequest](../../models/operations/deletelogftprequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.DeleteLogFtpSecurity](../../models/operations/deletelogftpsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.GetLogFtpSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingFtp.GetLogFtp(ctx, operations.GetLogFtpRequest{
        LoggingFtpName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingFtpResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetLogFtpRequest](../../models/operations/getlogftprequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.GetLogFtpSecurity](../../models/operations/getlogftpsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListLogFtpSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingFtp.ListLogFtp(ctx, operations.ListLogFtpRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingFtpResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListLogFtpRequest](../../models/operations/listlogftprequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.ListLogFtpSecurity](../../models/operations/listlogftpsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.UpdateLogFtpSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingFtp.UpdateLogFtp(ctx, operations.UpdateLogFtpRequest{
        LoggingFtpInput: &shared.LoggingFtpInput{
            Address: sdk.String("5812 Casimir Lock"),
            CompressionCodec: shared.LoggingFtpCompressionCodecZstd.ToPointer(),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingFtpFormatVersionOne.ToPointer(),
            GzipLevel: sdk.Int64(0),
            Hostname: sdk.String("scared-underpass.com"),
            Ipv4: sdk.String("195.191.191.23"),
            MessageType: shared.LoggingFtpMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Password: sdk.String("saepe"),
            Path: sdk.String("ipsum"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingFtpPlacementNone.ToPointer(),
            Port: sdk.Int64(749255),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            User: sdk.String("quos"),
        },
        LoggingFtpName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingFtpResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateLogFtpRequest](../../models/operations/updatelogftprequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.UpdateLogFtpSecurity](../../models/operations/updatelogftpsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.UpdateLogFtpResponse](../../models/operations/updatelogftpresponse.md), error**

