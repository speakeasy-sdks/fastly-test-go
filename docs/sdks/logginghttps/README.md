# LoggingHTTPS

## Overview

Fastly will upload log messages to an HTTPS endpoint in the format specified in the HTTPS object. The HTTPS endpoint requires proof of domain ownership before logs can be received. Learn how to validate your domain in our [HTTPS endpoint documentation](https://docs.fastly.com/en/guides/log-streaming-https).

<https://developer.fastly.com/reference/api/logging/https>
### Available Operations

* [CreateLogHTTPS](#createloghttps) - Create an HTTPS log endpoint
* [DeleteLogHTTPS](#deleteloghttps) - Delete an HTTPS log endpoint
* [GetLogHTTPS](#getloghttps) - Get an HTTPS log endpoint
* [ListLogHTTPS](#listloghttps) - List HTTPS log endpoints
* [UpdateLogHTTPS](#updateloghttps) - Update an HTTPS log endpoint

## CreateLogHTTPS

Create an HTTPS object for a particular service and version.

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
    res, err := s.LoggingHTTPS.CreateLogHTTPS(ctx, operations.CreateLogHTTPSRequest{
        LoggingHttps4: &shared.LoggingHttps4{
            ContentType: sdk.String("porro"),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingHTTPSFormatVersionOne.ToPointer(),
            HeaderName: sdk.String("quas"),
            HeaderValue: sdk.String("praesentium"),
            JSONFormat: shared.LoggingHTTPSJSONFormatZero.ToPointer(),
            MessageType: shared.LoggingMessageTypeClassic.ToPointer(),
            Method: shared.LoggingHTTPSMethodPut.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingHTTPSPlacementNone.ToPointer(),
            RequestMaxBytes: sdk.Int64(681393),
            RequestMaxEntries: sdk.Int64(649463),
            ResponseCondition: sdk.String("null"),
            TLSCaCert: sdk.String("incidunt"),
            TLSClientCert: sdk.String("atque"),
            TLSClientKey: sdk.String("explicabo"),
            TLSHostname: sdk.String("minima"),
            URL: sdk.String("nisi"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateLogHTTPSSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHTTPSResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateLogHTTPSRequest](../../models/operations/createloghttpsrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.CreateLogHTTPSSecurity](../../models/operations/createloghttpssecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.CreateLogHTTPSResponse](../../models/operations/createloghttpsresponse.md), error**


## DeleteLogHTTPS

Delete the HTTPS object for a particular service and version.

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
    res, err := s.LoggingHTTPS.DeleteLogHTTPS(ctx, operations.DeleteLogHTTPSRequest{
        LoggingHTTPSName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteLogHTTPSSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogHTTPS200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteLogHTTPSRequest](../../models/operations/deleteloghttpsrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.DeleteLogHTTPSSecurity](../../models/operations/deleteloghttpssecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.DeleteLogHTTPSResponse](../../models/operations/deleteloghttpsresponse.md), error**


## GetLogHTTPS

Get the HTTPS object for a particular service and version.

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
    res, err := s.LoggingHTTPS.GetLogHTTPS(ctx, operations.GetLogHTTPSRequest{
        LoggingHTTPSName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetLogHTTPSSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHTTPSResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetLogHTTPSRequest](../../models/operations/getloghttpsrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.GetLogHTTPSSecurity](../../models/operations/getloghttpssecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.GetLogHTTPSResponse](../../models/operations/getloghttpsresponse.md), error**


## ListLogHTTPS

List all of the HTTPS objects for a particular service and version.

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
    res, err := s.LoggingHTTPS.ListLogHTTPS(ctx, operations.ListLogHTTPSRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListLogHTTPSSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHTTPSResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListLogHTTPSRequest](../../models/operations/listloghttpsrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.ListLogHTTPSSecurity](../../models/operations/listloghttpssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.ListLogHTTPSResponse](../../models/operations/listloghttpsresponse.md), error**


## UpdateLogHTTPS

Update the HTTPS object for a particular service and version.

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
    res, err := s.LoggingHTTPS.UpdateLogHTTPS(ctx, operations.UpdateLogHTTPSRequest{
        LoggingHttps4: &shared.LoggingHttps4{
            ContentType: sdk.String("fugit"),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingHTTPSFormatVersionTwo.ToPointer(),
            HeaderName: sdk.String("consequuntur"),
            HeaderValue: sdk.String("ratione"),
            JSONFormat: shared.LoggingHTTPSJSONFormatZero.ToPointer(),
            MessageType: shared.LoggingMessageTypeClassic.ToPointer(),
            Method: shared.LoggingHTTPSMethodPut.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingHTTPSPlacementWafDebug.ToPointer(),
            RequestMaxBytes: sdk.Int64(543806),
            RequestMaxEntries: sdk.Int64(92260),
            ResponseCondition: sdk.String("null"),
            TLSCaCert: sdk.String("esse"),
            TLSClientCert: sdk.String("eveniet"),
            TLSClientKey: sdk.String("accusamus"),
            TLSHostname: sdk.String("veritatis"),
            URL: sdk.String("esse"),
        },
        LoggingHTTPSName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.UpdateLogHTTPSSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingHTTPSResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateLogHTTPSRequest](../../models/operations/updateloghttpsrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.UpdateLogHTTPSSecurity](../../models/operations/updateloghttpssecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.UpdateLogHTTPSResponse](../../models/operations/updateloghttpsresponse.md), error**

