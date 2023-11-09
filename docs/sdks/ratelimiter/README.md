# RateLimiter
(*RateLimiter*)

## Overview

Rate limiters add configurable origin request rate limiting to a service. This information is part of a limited availability release. For more information, see our [product and feature lifecycle](https://docs.fastly.com/products/fastly-product-lifecycle#limited-availability) descriptions. To use this feature you must purchase a Professional or Premier Platform subscription for either [Signal Sciences Cloud WAF](https://docs.fastly.com/products/signal-sciences-cloud-waf) or [Signal Sciences Next-Gen WAF](https://docs.fastly.com/products/signal-sciences-next-gen-waf) and have a [paid account with a contract](https://docs.fastly.com/en/guides/accounts-and-pricing-plans#paid-accounts-with-contractual-commitments) for [full-site delivery](https://docs.fastly.com/products/fastlys-legacy-full-site-delivery-services).

<https://developer.fastly.com/reference/api/vcl-services/rate-limiter>
### Available Operations

* [DeleteRateLimiter](#deleteratelimiter) - Delete a rate limiter
* [GetRateLimiter](#getratelimiter) - Get a rate limiter
* [ListRateLimiters](#listratelimiters) - List rate limiters

## DeleteRateLimiter

Delete a rate limiter by its ID.

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
    res, err := s.RateLimiter.DeleteRateLimiter(ctx, operations.DeleteRateLimiterRequest{
        RateLimiterID: "s7aqgcJjqqKhwiTRMaP11",
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
| `request`                                                                                  | [operations.DeleteRateLimiterRequest](../../models/operations/deleteratelimiterrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteRateLimiterResponse](../../models/operations/deleteratelimiterresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetRateLimiter

Get a rate limiter by its ID.

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
    res, err := s.RateLimiter.GetRateLimiter(ctx, operations.GetRateLimiterRequest{
        RateLimiterID: "s7aqgcJjqqKhwiTRMaP11",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RateLimiterResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetRateLimiterRequest](../../models/operations/getratelimiterrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetRateLimiterResponse](../../models/operations/getratelimiterresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListRateLimiters

List all rate limiters for a particular service and version.

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
    res, err := s.RateLimiter.ListRateLimiters(ctx, operations.ListRateLimitersRequest{
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListRateLimitersRequest](../../models/operations/listratelimitersrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListRateLimitersResponse](../../models/operations/listratelimitersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
