# LoggingCloudfiles
(*LoggingCloudfiles*)

## Overview

Fastly will upload log messages to your Rackspace Cloud Files account.

<https://developer.fastly.com/reference/api/logging/cloudfiles>
### Available Operations

* [CreateLogCloudfiles](#createlogcloudfiles) - Create a Cloud Files log endpoint
* [DeleteLogCloudfiles](#deletelogcloudfiles) - Delete the Cloud Files log endpoint
* [GetLogCloudfiles](#getlogcloudfiles) - Get a Cloud Files log endpoint
* [ListLogCloudfiles](#listlogcloudfiles) - List Cloud Files log endpoints
* [UpdateLogCloudfiles](#updatelogcloudfiles) - Update the Cloud Files log endpoint

## CreateLogCloudfiles

Create a Cloud Files log endpoint for a particular service and version.

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingCloudfiles.CreateLogCloudfiles(ctx, operations.CreateLogCloudfilesRequest{
        LoggingCloudfilesInput: &shared.LoggingCloudfilesInput{
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: shared.LoggingCloudfilesFormatVersionTwo.ToPointer(),
            GzipLevel: fastly.Int64(0),
            MessageType: shared.LoggingCloudfilesMessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Period: fastly.Int64(3600),
            Placement: shared.LoggingCloudfilesPlacementLessThanNilGreaterThan.ToPointer(),
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

    if res.LoggingCloudfilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateLogCloudfilesRequest](../../models/operations/createlogcloudfilesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.CreateLogCloudfilesResponse](../../models/operations/createlogcloudfilesresponse.md), error**


## DeleteLogCloudfiles

Delete the Cloud Files log endpoint for a particular service and version.

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingCloudfiles.DeleteLogCloudfiles(ctx, operations.DeleteLogCloudfilesRequest{
        LoggingCloudfilesName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogCloudfiles200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteLogCloudfilesRequest](../../models/operations/deletelogcloudfilesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteLogCloudfilesResponse](../../models/operations/deletelogcloudfilesresponse.md), error**


## GetLogCloudfiles

Get the Cloud Files log endpoint for a particular service and version.

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingCloudfiles.GetLogCloudfiles(ctx, operations.GetLogCloudfilesRequest{
        LoggingCloudfilesName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingCloudfilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetLogCloudfilesRequest](../../models/operations/getlogcloudfilesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetLogCloudfilesResponse](../../models/operations/getlogcloudfilesresponse.md), error**


## ListLogCloudfiles

List all of the Cloud Files log endpoints for a particular service and version.

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingCloudfiles.ListLogCloudfiles(ctx, operations.ListLogCloudfilesRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingCloudfilesResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListLogCloudfilesRequest](../../models/operations/listlogcloudfilesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListLogCloudfilesResponse](../../models/operations/listlogcloudfilesresponse.md), error**


## UpdateLogCloudfiles

Update the Cloud Files log endpoint for a particular service and version.

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingCloudfiles.UpdateLogCloudfiles(ctx, operations.UpdateLogCloudfilesRequest{
        LoggingCloudfilesInput: &shared.LoggingCloudfilesInput{
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: shared.LoggingCloudfilesFormatVersionTwo.ToPointer(),
            GzipLevel: fastly.Int64(0),
            MessageType: shared.LoggingCloudfilesMessageTypeClassic.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Period: fastly.Int64(3600),
            Placement: shared.LoggingCloudfilesPlacementLessThanNilGreaterThan.ToPointer(),
            PublicKey: fastly.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: fastly.String("null"),
        },
        LoggingCloudfilesName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingCloudfilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateLogCloudfilesRequest](../../models/operations/updatelogcloudfilesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdateLogCloudfilesResponse](../../models/operations/updatelogcloudfilesresponse.md), error**

