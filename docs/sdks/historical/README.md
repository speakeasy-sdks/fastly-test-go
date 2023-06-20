# Historical

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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Historical.GetHistStats(ctx, operations.GetHistStatsRequest{
        By: shared.ByDay.ToPointer(),
        From: sdk.String("incidunt"),
        Region: shared.RegionUsa.ToPointer(),
        To: sdk.String("qui"),
    }, operations.GetHistStatsSecurity{
        Token: "",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetHistStatsRequest](../../models/operations/gethiststatsrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.GetHistStatsSecurity](../../models/operations/gethiststatssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.GetHistStatsResponse](../../models/operations/gethiststatsresponse.md), error**


## GetHistStatsAggregated

Fetches historical stats information aggregated across all of your Fastly services.

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
    res, err := s.Historical.GetHistStatsAggregated(ctx, operations.GetHistStatsAggregatedRequest{
        By: shared.ByDay.ToPointer(),
        From: sdk.String("cupiditate"),
        Region: shared.RegionUsa.ToPointer(),
        To: sdk.String("maxime"),
    }, operations.GetHistStatsAggregatedSecurity{
        Token: "",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetHistStatsAggregatedRequest](../../models/operations/gethiststatsaggregatedrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.GetHistStatsAggregatedSecurity](../../models/operations/gethiststatsaggregatedsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.GetHistStatsAggregatedResponse](../../models/operations/gethiststatsaggregatedresponse.md), error**


## GetHistStatsField

Fetches the specified field from the historical stats for each of your services and groups the results by service ID.

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
    res, err := s.Historical.GetHistStatsField(ctx, operations.GetHistStatsFieldRequest{
        By: shared.ByDay.ToPointer(),
        Field: "hit_ratio",
        From: sdk.String("pariatur"),
        Region: shared.RegionUsa.ToPointer(),
        To: sdk.String("soluta"),
    }, operations.GetHistStatsFieldSecurity{
        Token: "",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetHistStatsFieldRequest](../../models/operations/gethiststatsfieldrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.GetHistStatsFieldSecurity](../../models/operations/gethiststatsfieldsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.GetHistStatsFieldResponse](../../models/operations/gethiststatsfieldresponse.md), error**


## GetHistStatsService

Fetches historical stats for a given service.

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
    res, err := s.Historical.GetHistStatsService(ctx, operations.GetHistStatsServiceRequest{
        By: shared.ByDay.ToPointer(),
        From: sdk.String("dicta"),
        Region: shared.RegionUsa.ToPointer(),
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        To: sdk.String("laborum"),
    }, operations.GetHistStatsServiceSecurity{
        Token: "",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetHistStatsServiceRequest](../../models/operations/gethiststatsservicerequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.GetHistStatsServiceSecurity](../../models/operations/gethiststatsservicesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.GetHistStatsServiceResponse](../../models/operations/gethiststatsserviceresponse.md), error**


## GetHistStatsServiceField

Fetches the specified field from the historical stats for a given service.

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
    res, err := s.Historical.GetHistStatsServiceField(ctx, operations.GetHistStatsServiceFieldRequest{
        By: shared.ByDay.ToPointer(),
        Field: "hit_ratio",
        From: sdk.String("totam"),
        Region: shared.RegionUsa.ToPointer(),
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        To: sdk.String("incidunt"),
    }, operations.GetHistStatsServiceFieldSecurity{
        Token: "",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetHistStatsServiceFieldRequest](../../models/operations/gethiststatsservicefieldrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.GetHistStatsServiceFieldSecurity](../../models/operations/gethiststatsservicefieldsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.GetHistStatsServiceFieldResponse](../../models/operations/gethiststatsservicefieldresponse.md), error**


## GetRegions

Fetches the list of codes for regions that are covered by the Fastly CDN service.

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
    res, err := s.Historical.GetRegions(ctx, operations.GetRegionsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HistoricalRegionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `security`                                                                     | [operations.GetRegionsSecurity](../../models/operations/getregionssecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.GetRegionsResponse](../../models/operations/getregionsresponse.md), error**


## GetUsage

Returns usage information aggregated across all Fastly services and grouped by region. To aggregate across all Fastly services by time period, see [`/stats/aggregate`](#get-hist-stats-aggregated).

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
    res, err := s.Historical.GetUsage(ctx, operations.GetUsageRequest{
        From: sdk.String("aspernatur"),
        To: sdk.String("dolores"),
    }, operations.GetUsageSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HistoricalUsageAggregateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetUsageRequest](../../models/operations/getusagerequest.md)   | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.GetUsageSecurity](../../models/operations/getusagesecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.GetUsageResponse](../../models/operations/getusageresponse.md), error**


## GetUsageMonth

Returns month-to-date usage details for a given month and year. Usage details are aggregated by service and across all Fastly services, and then grouped by region. This endpoint does not use the `from` or `to` fields for selecting the date for which data is requested. Instead, it uses `month` and `year` integer fields. Both fields are optional and default to the current month and year respectively. When set, an optional `billable_units` field will convert bandwidth to GB and divide requests by 10,000.

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
    res, err := s.Historical.GetUsageMonth(ctx, operations.GetUsageMonthRequest{
        BillableUnits: sdk.Bool(true),
        Month: sdk.String("05"),
        Year: sdk.String("2020"),
    }, operations.GetUsageMonthSecurity{
        Token: "",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetUsageMonthRequest](../../models/operations/getusagemonthrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.GetUsageMonthSecurity](../../models/operations/getusagemonthsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.GetUsageMonthResponse](../../models/operations/getusagemonthresponse.md), error**


## GetUsageService

Returns usage information aggregated by service and grouped by service and region. For service stats by time period, see [`/stats`](#get-hist-stats) and [`/stats/field/:field`](#get-hist-stats-field).

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
    res, err := s.Historical.GetUsageService(ctx, operations.GetUsageServiceRequest{
        From: sdk.String("distinctio"),
        To: sdk.String("facilis"),
    }, operations.GetUsageServiceSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HistoricalUsageServiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetUsageServiceRequest](../../models/operations/getusageservicerequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.GetUsageServiceSecurity](../../models/operations/getusageservicesecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.GetUsageServiceResponse](../../models/operations/getusageserviceresponse.md), error**

