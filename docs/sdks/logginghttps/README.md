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
    operationSecurity := operations.CreateLogHTTPSSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHTTPS.CreateLogHTTPS(ctx, operations.CreateLogHTTPSRequest{
        LoggingHttps4: &shared.LoggingHttps4{
            ContentType: sdk.String("quisquam"),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingHTTPSFormatVersionTwo.ToPointer(),
            HeaderName: sdk.String("omnis"),
            HeaderValue: sdk.String("quis"),
            JSONFormat: shared.LoggingHTTPSJSONFormatZero.ToPointer(),
            MessageType: shared.LoggingMessageTypeClassic.ToPointer(),
            Method: shared.LoggingHTTPSMethodPut.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingHTTPSPlacementWafDebug.ToPointer(),
            RequestMaxBytes: sdk.Int64(231701),
            RequestMaxEntries: sdk.Int64(878870),
            ResponseCondition: sdk.String("null"),
            TLSCaCert: sdk.String("tenetur"),
            TLSClientCert: sdk.String("dignissimos"),
            TLSClientKey: sdk.String("hic"),
            TLSHostname: sdk.String("distinctio"),
            URL: sdk.String("quod"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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
    operationSecurity := operations.DeleteLogHTTPSSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHTTPS.DeleteLogHTTPS(ctx, operations.DeleteLogHTTPSRequest{
        LoggingHTTPSName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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
    operationSecurity := operations.GetLogHTTPSSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHTTPS.GetLogHTTPS(ctx, operations.GetLogHTTPSRequest{
        LoggingHTTPSName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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
    operationSecurity := operations.ListLogHTTPSSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHTTPS.ListLogHTTPS(ctx, operations.ListLogHTTPSRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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
    operationSecurity := operations.UpdateLogHTTPSSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingHTTPS.UpdateLogHTTPS(ctx, operations.UpdateLogHTTPSRequest{
        LoggingHttps4: &shared.LoggingHttps4{
            ContentType: sdk.String("odio"),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingHTTPSFormatVersionTwo.ToPointer(),
            HeaderName: sdk.String("facilis"),
            HeaderValue: sdk.String("vero"),
            JSONFormat: shared.LoggingHTTPSJSONFormatOne.ToPointer(),
            MessageType: shared.LoggingMessageTypeClassic.ToPointer(),
            Method: shared.LoggingHTTPSMethodPost.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingHTTPSPlacementLessThanNilGreaterThan.ToPointer(),
            RequestMaxBytes: sdk.Int64(848944),
            RequestMaxEntries: sdk.Int64(194342),
            ResponseCondition: sdk.String("null"),
            TLSCaCert: sdk.String("natus"),
            TLSClientCert: sdk.String("impedit"),
            TLSClientKey: sdk.String("aut"),
            TLSHostname: sdk.String("voluptatibus"),
            URL: sdk.String("exercitationem"),
        },
        LoggingHTTPSName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
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

