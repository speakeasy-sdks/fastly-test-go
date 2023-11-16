# Purge
(*Purge*)

## Overview

Instant Purging removes content from Fastly immediately so it can be refreshed from your origin servers. While the default approach for issuing an individual URL Instant Purge uses the Fastly API, `https://api.fastly.com/`, it is not required.

<https://developer.fastly.com/reference/api/purging>
### Available Operations

* [PurgeAll](#purgeall) - Purge everything from a service
* [PurgeSingleURL](#purgesingleurl) - Purge a URL
* [PurgeTag](#purgetag) - Purge by surrogate key tag

## PurgeAll

Instant Purge everything from a service.

Purge-all requests cannot be done in soft mode and will always immediately invalidate all cached content associated with the service. To do a soft-purge-all, consider applying a constant [surrogate key](https://docs.fastly.com/en/guides/getting-started-with-surrogate-keys) tag (e.g., `"all"`) to all objects.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Purge.PurgeAll(ctx, operations.PurgeAllRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.PurgeAllRequest](../../models/operations/purgeallrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.PurgeAllResponse](../../models/operations/purgeallresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PurgeSingleURL

Instant Purge an individual URL.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Purge.PurgeSingleURL(ctx, operations.PurgeSingleURLRequest{
        CachedURL: "www.example.com/path/to/object-to-purge",
        FastlySoftPurge: fastlytestgo.Int64(1),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PurgeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PurgeSingleURLRequest](../../models/operations/purgesingleurlrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PurgeSingleURLResponse](../../models/operations/purgesingleurlresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PurgeTag

Instant Purge a particular service of items tagged with a Surrogate Key. Only one surrogate key can be purged at a time. Multiple keys can be purged using a batch surrogate key purge request.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Purge.PurgeTag(ctx, operations.PurgeTagRequest{
        FastlySoftPurge: fastlytestgo.Int64(1),
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        SurrogateKey: "key_1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PurgeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.PurgeTagRequest](../../models/operations/purgetagrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.PurgeTagResponse](../../models/operations/purgetagresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
