# Realtime

## Overview

The real-time analytics API offers a standardized set of data about traffic received by a specified service in one-second time periods up to the last complete second.

<https://developer.fastly.com/reference/api/metrics-stats/realtime>
### Available Operations

* [GetStatsLast120Seconds](#getstatslast120seconds) - Get real-time data for the last 120 seconds
* [GetStatsLast120SecondsLimitEntries](#getstatslast120secondslimitentries) - Get a limited number of real-time data entries
* [GetStatsLastSecond](#getstatslastsecond) - Get real-time data from specified time

## GetStatsLast120Seconds

Get data for the 120 seconds preceding the latest timestamp available for a service.

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
    res, err := s.Realtime.GetStatsLast120Seconds(ctx, operations.GetStatsLast120SecondsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.GetStatsLast120SecondsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Realtime != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetStatsLast120SecondsRequest](../../models/operations/getstatslast120secondsrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.GetStatsLast120SecondsSecurity](../../models/operations/getstatslast120secondssecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.GetStatsLast120SecondsResponse](../../models/operations/getstatslast120secondsresponse.md), error**


## GetStatsLast120SecondsLimitEntries

Get data for the 120 seconds preceding the latest timestamp available for a service, up to a maximum of `max_entries` entries.

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
    res, err := s.Realtime.GetStatsLast120SecondsLimitEntries(ctx, operations.GetStatsLast120SecondsLimitEntriesRequest{
        MaxEntries: 1,
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.GetStatsLast120SecondsLimitEntriesSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Realtime != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetStatsLast120SecondsLimitEntriesRequest](../../models/operations/getstatslast120secondslimitentriesrequest.md)   | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `security`                                                                                                                     | [operations.GetStatsLast120SecondsLimitEntriesSecurity](../../models/operations/getstatslast120secondslimitentriessecurity.md) | :heavy_check_mark:                                                                                                             | The security requirements to use for the request.                                                                              |


### Response

**[*operations.GetStatsLast120SecondsLimitEntriesResponse](../../models/operations/getstatslast120secondslimitentriesresponse.md), error**


## GetStatsLastSecond

Get real-time data for the specified reporting period. Specify `0` to get a single entry for the last complete second. The `Timestamp` field included in the response provides the time index of the latest entry in the dataset and can be provided as the `start_timestamp` of the next request for a seamless continuation of the dataset from one request to the next.

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
    res, err := s.Realtime.GetStatsLastSecond(ctx, operations.GetStatsLastSecondRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        TimestampInSeconds: 1608560817,
    }, operations.GetStatsLastSecondSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Realtime != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetStatsLastSecondRequest](../../models/operations/getstatslastsecondrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.GetStatsLastSecondSecurity](../../models/operations/getstatslastsecondsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.GetStatsLastSecondResponse](../../models/operations/getstatslastsecondresponse.md), error**

