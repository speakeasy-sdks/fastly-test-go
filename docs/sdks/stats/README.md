# Stats
(*Stats*)

## Overview

Stats give you information on the usage and performance of your Service. They can be requested by Service and over a particular time span. Stats are broken down per POP, giving you information on how your Services are being used across the world. There is now a more flexible, and fully featured [Stats API](/reference/api/metrics-stats/historical-stats/) available.

<https://developer.fastly.com/reference/api/metrics-stats/stats>
### Available Operations

* [GetServiceStats](#getservicestats) - Get stats for a service

## GetServiceStats

Get the stats from a service for a block of time. This lists all stats by PoP location, starting with AMS. This call requires parameters to select block of time to query. Use either a timestamp range (using start_time and end_time) or a specified month/year combo (using month and year).

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
    res, err := s.Stats.GetServiceStats(ctx, operations.GetServiceStatsRequest{
        EndTime: fastlytestgo.Int64(1608560817),
        Month: fastlytestgo.String("05"),
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        StartTime: fastlytestgo.Int64(1608560817),
        Year: fastlytestgo.String("2020"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Stats != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetServiceStatsRequest](../../models/operations/getservicestatsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetServiceStatsResponse](../../models/operations/getservicestatsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
