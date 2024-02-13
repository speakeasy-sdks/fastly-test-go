# LoggingLogentries
(*LoggingLogentries*)

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

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.LoggingLogentries.CreateLogLogentries(ctx, operations.CreateLogLogentriesRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogentriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateLogLogentriesRequest](../../models/operations/createloglogentriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.CreateLogLogentriesResponse](../../models/operations/createloglogentriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~DeleteLogLogentries~~

Delete the Logentry for a particular service and version.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.LoggingLogentries.DeleteLogLogentries(ctx, operations.DeleteLogLogentriesRequest{
        LoggingLogentriesName: "test-log-endpoint",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteLogLogentriesRequest](../../models/operations/deleteloglogentriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteLogLogentriesResponse](../../models/operations/deleteloglogentriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~GetLogLogentries~~

Get the Logentry for a particular service and version.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.LoggingLogentries.GetLogLogentries(ctx, operations.GetLogLogentriesRequest{
        LoggingLogentriesName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogentriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetLogLogentriesRequest](../../models/operations/getloglogentriesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetLogLogentriesResponse](../../models/operations/getloglogentriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~ListLogLogentries~~

List all of the Logentries for a particular service and version.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.LoggingLogentries.ListLogLogentries(ctx, operations.ListLogLogentriesRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListLogLogentriesRequest](../../models/operations/listloglogentriesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListLogLogentriesResponse](../../models/operations/listloglogentriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~UpdateLogLogentries~~

Update the Logentry for a particular service and version.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.LoggingLogentries.UpdateLogLogentries(ctx, operations.UpdateLogLogentriesRequest{
        LoggingLogentriesName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingLogentriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateLogLogentriesRequest](../../models/operations/updateloglogentriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdateLogLogentriesResponse](../../models/operations/updateloglogentriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
