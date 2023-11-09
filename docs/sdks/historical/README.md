# Historical
(*Historical*)

## Overview

The Historical Stats API allows you to programmatically retrieve historical caching statistics derived from your Fastly services. You can use these metrics to help you optimize your site’s data caching and analyze your site’s traffic.

<https://developer.fastly.com/reference/api/metrics-stats/historical-stats>
### Available Operations

* [GetHistStats](#gethiststats) - Get historical stats
* [GetHistStatsAggregated](#gethiststatsaggregated) - Get aggregated historical stats
* [GetHistStatsField](#gethiststatsfield) - Get historical stats for a single field
* [GetHistStatsService](#gethiststatsservice) - Get historical stats for a single service
* [GetHistStatsServiceField](#gethiststatsservicefield) - Get historical stats for a single service/field combination
* [GetRegions](#getregions) - Get region codes
* [GetUsage](#getusage) - Get usage statistics
* [GetUsageMonth](#getusagemonth) - Get month-to-date usage statistics
* [GetUsageService](#getusageservice) - Get usage statistics per service

## GetHistStats

Fetches historical stats for each of your Fastly services and groups the results by service ID.

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
    res, err := s.Historical.GetHistStats(ctx, operations.GetHistStatsRequest{
        By: components.ByDay.ToPointer(),
        Region: components.RegionUsa.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HistoricalResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetHistStatsRequest](../../models/operations/gethiststatsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetHistStatsResponse](../../models/operations/gethiststatsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetHistStatsAggregated

Fetches historical stats information aggregated across all of your Fastly services.

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
    res, err := s.Historical.GetHistStatsAggregated(ctx, operations.GetHistStatsAggregatedRequest{
        By: components.ByDay.ToPointer(),
        Region: components.RegionUsa.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HistoricalAggregateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetHistStatsAggregatedRequest](../../models/operations/gethiststatsaggregatedrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetHistStatsAggregatedResponse](../../models/operations/gethiststatsaggregatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetHistStatsField

Fetches the specified field from the historical stats for each of your services and groups the results by service ID.

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
    res, err := s.Historical.GetHistStatsField(ctx, operations.GetHistStatsFieldRequest{
        By: components.ByDay.ToPointer(),
        Field: "hit_ratio",
        Region: components.RegionUsa.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HistoricalFieldResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetHistStatsFieldRequest](../../models/operations/gethiststatsfieldrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetHistStatsFieldResponse](../../models/operations/gethiststatsfieldresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetHistStatsService

Fetches historical stats for a given service.

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
    res, err := s.Historical.GetHistStatsService(ctx, operations.GetHistStatsServiceRequest{
        By: components.ByDay.ToPointer(),
        Region: components.RegionUsa.ToPointer(),
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HistoricalAggregateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetHistStatsServiceRequest](../../models/operations/gethiststatsservicerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetHistStatsServiceResponse](../../models/operations/gethiststatsserviceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetHistStatsServiceField

Fetches the specified field from the historical stats for a given service.

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
    res, err := s.Historical.GetHistStatsServiceField(ctx, operations.GetHistStatsServiceFieldRequest{
        By: components.ByDay.ToPointer(),
        Field: "hit_ratio",
        Region: components.RegionUsa.ToPointer(),
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HistoricalFieldAggregateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetHistStatsServiceFieldRequest](../../models/operations/gethiststatsservicefieldrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetHistStatsServiceFieldResponse](../../models/operations/gethiststatsservicefieldresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetRegions

Fetches the list of codes for regions that are covered by the Fastly CDN service.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Historical.GetRegions(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.HistoricalRegionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetRegionsResponse](../../models/operations/getregionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetUsage

Returns usage information aggregated across all Fastly services and grouped by region. To aggregate across all Fastly services by time period, see [`/stats/aggregate`](#get-hist-stats-aggregated).

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
    res, err := s.Historical.GetUsage(ctx, operations.GetUsageRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.HistoricalUsageAggregateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetUsageRequest](../../models/operations/getusagerequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetUsageResponse](../../models/operations/getusageresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetUsageMonth

Returns month-to-date usage details for a given month and year. Usage details are aggregated by service and across all Fastly services, and then grouped by region. This endpoint does not use the `from` or `to` fields for selecting the date for which data is requested. Instead, it uses `month` and `year` integer fields. Both fields are optional and default to the current month and year respectively. When set, an optional `billable_units` field will convert bandwidth to GB and divide requests by 10,000.

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
    res, err := s.Historical.GetUsageMonth(ctx, operations.GetUsageMonthRequest{
        BillableUnits: fastly.Bool(true),
        Month: fastly.String("05"),
        Year: fastly.String("2020"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HistoricalUsageMonthResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetUsageMonthRequest](../../models/operations/getusagemonthrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetUsageMonthResponse](../../models/operations/getusagemonthresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetUsageService

Returns usage information aggregated by service and grouped by service and region. For service stats by time period, see [`/stats`](#get-hist-stats) and [`/stats/field/:field`](#get-hist-stats-field).

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
    res, err := s.Historical.GetUsageService(ctx, operations.GetUsageServiceRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.HistoricalUsageServiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetUsageServiceRequest](../../models/operations/getusageservicerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetUsageServiceResponse](../../models/operations/getusageserviceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
