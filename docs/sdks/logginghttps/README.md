# LoggingHTTPS
(*LoggingHTTPS*)

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
    res, err := s.LoggingHTTPS.CreateLogHTTPS(ctx, operations.CreateLogHTTPSRequest{
        LoggingHTTPS: &shared.LoggingHTTPS{
            ContentType: fastly.String("wireless Bicycle"),
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: shared.LoggingHTTPSFormatVersionOne.ToPointer(),
            HeaderName: fastly.String("schemas Beauty Rolls"),
            HeaderValue: fastly.String("shallow"),
            JSONFormat: shared.LoggingHTTPSJSONFormatZero.ToPointer(),
            MessageType: shared.LoggingMessageTypeClassic.ToPointer(),
            Method: shared.LoggingHTTPSMethodPut.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingHTTPSPlacementWafDebug.ToPointer(),
            RequestMaxBytes: fastly.Int64(879956),
            RequestMaxEntries: fastly.Int64(842598),
            ResponseCondition: fastly.String("null"),
            TLSCaCert: fastly.String("Coordinator although truthfully"),
            TLSClientCert: fastly.String("especially World Northeast"),
            TLSClientKey: fastly.String("nor Passenger"),
            TLSHostname: fastly.String("Forward"),
            URL: fastly.String("http://narrow-puffin.name"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateLogHTTPSRequest](../../models/operations/createloghttpsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
    res, err := s.LoggingHTTPS.DeleteLogHTTPS(ctx, operations.DeleteLogHTTPSRequest{
        LoggingHTTPSName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteLogHTTPSRequest](../../models/operations/deleteloghttpsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
    res, err := s.LoggingHTTPS.GetLogHTTPS(ctx, operations.GetLogHTTPSRequest{
        LoggingHTTPSName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetLogHTTPSRequest](../../models/operations/getloghttpsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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
    res, err := s.LoggingHTTPS.ListLogHTTPS(ctx, operations.ListLogHTTPSRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListLogHTTPSRequest](../../models/operations/listloghttpsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


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
    res, err := s.LoggingHTTPS.UpdateLogHTTPS(ctx, operations.UpdateLogHTTPSRequest{
        LoggingHTTPS: &shared.LoggingHTTPS{
            ContentType: fastly.String("Iowa Sausages"),
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: shared.LoggingHTTPSFormatVersionOne.ToPointer(),
            HeaderName: fastly.String("impactful invoice aggregate"),
            HeaderValue: fastly.String("Assistant"),
            JSONFormat: shared.LoggingHTTPSJSONFormatOne.ToPointer(),
            MessageType: shared.LoggingMessageTypeClassic.ToPointer(),
            Method: shared.LoggingHTTPSMethodPost.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingHTTPSPlacementNone.ToPointer(),
            RequestMaxBytes: fastly.Int64(382846),
            RequestMaxEntries: fastly.Int64(490200),
            ResponseCondition: fastly.String("null"),
            TLSCaCert: fastly.String("rightfully possimus"),
            TLSClientCert: fastly.String("Credit"),
            TLSClientKey: fastly.String("blue"),
            TLSHostname: fastly.String("indexing North"),
            URL: fastly.String("https://imaginary-louse.biz"),
        },
        LoggingHTTPSName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateLogHTTPSRequest](../../models/operations/updateloghttpsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateLogHTTPSResponse](../../models/operations/updateloghttpsresponse.md), error**

