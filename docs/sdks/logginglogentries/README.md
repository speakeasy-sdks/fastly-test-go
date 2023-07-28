# LoggingLogentries

## Overview

The Logentries integration has been discontinued.  No new Logentries endpoints can be created.

<https://developer.fastly.com/reference/api/logging/logentries>
### Available Operations

* [~~CreateLogLogentries~~](#createloglogentries) - Create a Logentries log endpoint :warning: **Deprecated**
* [~~DeleteLogLogentries~~](#deleteloglogentries) - Delete a Logentries log endpoint :warning: **Deprecated**
* [~~GetLogLogentries~~](#getloglogentries) - Get a Logentries log endpoint :warning: **Deprecated**
* [~~ListLogLogentries~~](#listloglogentries) - List Logentries log endpoints :warning: **Deprecated**
* [~~UpdateLogLogentries~~](#updateloglogentries) - Update a Logentries log endpoint :warning: **Deprecated**

## ~~CreateLogLogentries~~

Create a Logentry for a particular service and version.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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
    operationSecurity := operations.CreateLogLogentriesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingLogentries.CreateLogLogentries(ctx, operations.CreateLogLogentriesRequest{
        LoggingLogentries3: &shared.LoggingLogentries3{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingLogentriesFormatVersionTwo.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingLogentriesPlacementNone.ToPointer(),
            Port: sdk.Int64(458139),
            Region: shared.LoggingLogentriesRegionEu.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("provident"),
            UseTLS: shared.LoggingUseTLSOne.ToPointer(),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogentriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateLogLogentriesRequest](../../models/operations/createloglogentriesrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.CreateLogLogentriesSecurity](../../models/operations/createloglogentriessecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.CreateLogLogentriesResponse](../../models/operations/createloglogentriesresponse.md), error**


## ~~DeleteLogLogentries~~

Delete the Logentry for a particular service and version.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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
    operationSecurity := operations.DeleteLogLogentriesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingLogentries.DeleteLogLogentries(ctx, operations.DeleteLogLogentriesRequest{
        LoggingLogentriesName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogLogentries200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteLogLogentriesRequest](../../models/operations/deleteloglogentriesrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.DeleteLogLogentriesSecurity](../../models/operations/deleteloglogentriessecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.DeleteLogLogentriesResponse](../../models/operations/deleteloglogentriesresponse.md), error**


## ~~GetLogLogentries~~

Get the Logentry for a particular service and version.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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
    operationSecurity := operations.GetLogLogentriesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingLogentries.GetLogLogentries(ctx, operations.GetLogLogentriesRequest{
        LoggingLogentriesName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogentriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetLogLogentriesRequest](../../models/operations/getloglogentriesrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.GetLogLogentriesSecurity](../../models/operations/getloglogentriessecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.GetLogLogentriesResponse](../../models/operations/getloglogentriesresponse.md), error**


## ~~ListLogLogentries~~

List all of the Logentries for a particular service and version.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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
    operationSecurity := operations.ListLogLogentriesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingLogentries.ListLogLogentries(ctx, operations.ListLogLogentriesRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogentriesResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListLogLogentriesRequest](../../models/operations/listloglogentriesrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ListLogLogentriesSecurity](../../models/operations/listloglogentriessecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.ListLogLogentriesResponse](../../models/operations/listloglogentriesresponse.md), error**


## ~~UpdateLogLogentries~~

Update the Logentry for a particular service and version.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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
    operationSecurity := operations.UpdateLogLogentriesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingLogentries.UpdateLogLogentries(ctx, operations.UpdateLogLogentriesRequest{
        LoggingLogentries3: &shared.LoggingLogentries3{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingLogentriesFormatVersionTwo.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingLogentriesPlacementWafDebug.ToPointer(),
            Port: sdk.Int64(457223),
            Region: shared.LoggingLogentriesRegionUs.ToPointer(),
            ResponseCondition: sdk.String("null"),
            Token: sdk.String("a"),
            UseTLS: shared.LoggingUseTLSOne.ToPointer(),
        },
        LoggingLogentriesName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogentriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateLogLogentriesRequest](../../models/operations/updateloglogentriesrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.UpdateLogLogentriesSecurity](../../models/operations/updateloglogentriessecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.UpdateLogLogentriesResponse](../../models/operations/updateloglogentriesresponse.md), error**

