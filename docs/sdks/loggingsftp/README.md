# LoggingSftp

## Overview

Fastly will upload log messages periodically to the server in the format specified in the SFTP object.

<https://developer.fastly.com/reference/api/logging/sftp>
### Available Operations

* [CreateLogSftp](#createlogsftp) - Create an SFTP log endpoint
* [DeleteLogSftp](#deletelogsftp) - Delete an SFTP log endpoint
* [GetLogSftp](#getlogsftp) - Get an SFTP log endpoint
* [ListLogSftp](#listlogsftp) - List SFTP log endpoints
* [UpdateLogSftp](#updatelogsftp) - Update an SFTP log endpoint

## CreateLogSftp

Create a SFTP for a particular service and version.

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
    operationSecurity := operations.CreateLogSftpSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSftp.CreateLogSftp(ctx, operations.CreateLogSftpRequest{
        LoggingSftpInput: &shared.LoggingSftpInput{
            Address: sdk.String("example.com"),
            CompressionCodec: shared.LoggingSftpCompressionCodecZstd.ToPointer(),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSftpFormatVersionOne.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingSftpMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Password: sdk.String("porro"),
            Path: sdk.String("quod"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingSftpPlacementNone.ToPointer(),
            Port: sdk.Int64(70447),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            SecretKey: sdk.String("adipisci"),
            SSHKnownHosts: sdk.String("fuga"),
            User: sdk.String("id"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSftpResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateLogSftpRequest](../../models/operations/createlogsftprequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.CreateLogSftpSecurity](../../models/operations/createlogsftpsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.CreateLogSftpResponse](../../models/operations/createlogsftpresponse.md), error**


## DeleteLogSftp

Delete the SFTP for a particular service and version.

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
    operationSecurity := operations.DeleteLogSftpSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSftp.DeleteLogSftp(ctx, operations.DeleteLogSftpRequest{
        LoggingSftpName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogSftp200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteLogSftpRequest](../../models/operations/deletelogsftprequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.DeleteLogSftpSecurity](../../models/operations/deletelogsftpsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.DeleteLogSftpResponse](../../models/operations/deletelogsftpresponse.md), error**


## GetLogSftp

Get the SFTP for a particular service and version.

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
    operationSecurity := operations.GetLogSftpSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSftp.GetLogSftp(ctx, operations.GetLogSftpRequest{
        LoggingSftpName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSftpResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetLogSftpRequest](../../models/operations/getlogsftprequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.GetLogSftpSecurity](../../models/operations/getlogsftpsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.GetLogSftpResponse](../../models/operations/getlogsftpresponse.md), error**


## ListLogSftp

List all of the SFTPs for a particular service and version.

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
    operationSecurity := operations.ListLogSftpSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSftp.ListLogSftp(ctx, operations.ListLogSftpRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSftpResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListLogSftpRequest](../../models/operations/listlogsftprequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.ListLogSftpSecurity](../../models/operations/listlogsftpsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.ListLogSftpResponse](../../models/operations/listlogsftpresponse.md), error**


## UpdateLogSftp

Update the SFTP for a particular service and version.

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
    operationSecurity := operations.UpdateLogSftpSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSftp.UpdateLogSftp(ctx, operations.UpdateLogSftpRequest{
        LoggingSftpInput: &shared.LoggingSftpInput{
            Address: sdk.String("example.com"),
            CompressionCodec: shared.LoggingSftpCompressionCodecSnappy.ToPointer(),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSftpFormatVersionOne.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingSftpMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Password: sdk.String("culpa"),
            Path: sdk.String("est"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingSftpPlacementLessThanNilGreaterThan.ToPointer(),
            Port: sdk.Int64(517309),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            SecretKey: sdk.String("fugiat"),
            SSHKnownHosts: sdk.String("vel"),
            User: sdk.String("ducimus"),
        },
        LoggingSftpName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSftpResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateLogSftpRequest](../../models/operations/updatelogsftprequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.UpdateLogSftpSecurity](../../models/operations/updatelogsftpsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.UpdateLogSftpResponse](../../models/operations/updatelogsftpresponse.md), error**

