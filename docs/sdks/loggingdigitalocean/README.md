# LoggingDigitalocean
(*.LoggingDigitalocean*)

## Overview

Fastly will upload log messages to the DigitalOcean Space in the format specified in the DigitalOcean Spaces object.

<https://developer.fastly.com/reference/api/logging/digitalocean>
### Available Operations

* [CreateLogDigocean](#createlogdigocean) - Create a DigitalOcean Spaces log endpoint
* [DeleteLogDigocean](#deletelogdigocean) - Delete a DigitalOcean Spaces log endpoint
* [GetLogDigocean](#getlogdigocean) - Get a DigitalOcean Spaces log endpoint
* [ListLogDigocean](#listlogdigocean) - List DigitalOcean Spaces log endpoints
* [UpdateLogDigocean](#updatelogdigocean) - Update a DigitalOcean Spaces log endpoint

## CreateLogDigocean

Create a DigitalOcean Space for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingDigitalocean.CreateLogDigocean(ctx, operations.CreateLogDigoceanRequest{
        LoggingDigitalocean: &components.LoggingDigitalocean{
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: components.LoggingDigitaloceanFormatVersionTwo.ToPointer(),
            GzipLevel: fastly.Int64(0),
            MessageType: components.LoggingDigitaloceanMessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Period: fastly.Int64(3600),
            Placement: components.LoggingDigitaloceanPlacementWafDebug.ToPointer(),
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

    if res.LoggingDigitaloceanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateLogDigoceanRequest](../../models/operations/createlogdigoceanrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateLogDigoceanResponse](../../models/operations/createlogdigoceanresponse.md), error**


## DeleteLogDigocean

Delete the DigitalOcean Space for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingDigitalocean.DeleteLogDigocean(ctx, operations.DeleteLogDigoceanRequest{
        LoggingDigitaloceanName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteLogDigoceanRequest](../../models/operations/deletelogdigoceanrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteLogDigoceanResponse](../../models/operations/deletelogdigoceanresponse.md), error**


## GetLogDigocean

Get the DigitalOcean Space for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingDigitalocean.GetLogDigocean(ctx, operations.GetLogDigoceanRequest{
        LoggingDigitaloceanName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDigitaloceanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetLogDigoceanRequest](../../models/operations/getlogdigoceanrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetLogDigoceanResponse](../../models/operations/getlogdigoceanresponse.md), error**


## ListLogDigocean

List all of the DigitalOcean Spaces for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingDigitalocean.ListLogDigocean(ctx, operations.ListLogDigoceanRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListLogDigoceanRequest](../../models/operations/listlogdigoceanrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListLogDigoceanResponse](../../models/operations/listlogdigoceanresponse.md), error**


## UpdateLogDigocean

Update the DigitalOcean Space for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingDigitalocean.UpdateLogDigocean(ctx, operations.UpdateLogDigoceanRequest{
        LoggingDigitalocean: &components.LoggingDigitalocean{
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: components.LoggingDigitaloceanFormatVersionTwo.ToPointer(),
            GzipLevel: fastly.Int64(0),
            MessageType: components.LoggingDigitaloceanMessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Period: fastly.Int64(3600),
            Placement: components.LoggingDigitaloceanPlacementWafDebug.ToPointer(),
            PublicKey: fastly.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: fastly.String("null"),
        },
        LoggingDigitaloceanName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingDigitaloceanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateLogDigoceanRequest](../../models/operations/updatelogdigoceanrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateLogDigoceanResponse](../../models/operations/updatelogdigoceanresponse.md), error**

