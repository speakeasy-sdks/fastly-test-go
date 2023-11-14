# Events
(*Events*)

## Overview

[Event logs](https://docs.fastly.com/en/guides/reviewing-service-activity-with-the-event-log) are used to audit actions performed by customers.


<https://developer.fastly.com/reference/api/account/events>
### Available Operations

* [GetEvent](#getevent) - Get an event
* [ListEvents](#listevents) - List events

## GetEvent

Get a specific event.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Events.GetEvent(ctx, operations.GetEventRequest{
        EventID: "1PTzLK8g1NRKMGu5kUb8SC",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EventResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetEventRequest](../../models/operations/geteventrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetEventResponse](../../models/operations/geteventresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListEvents

List all events for a particular customer. Events can be filtered by user, customer and event type. Events can be sorted by date.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Events.ListEvents(ctx, operations.ListEventsRequest{
        FilterCustomerID: fastlytestgo.String("x4xCwxxJxGCx123Rx5xTx"),
        PageNumber: fastlytestgo.Int64(1),
        PageSize: fastlytestgo.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListEventsRequest](../../models/operations/listeventsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.ListEventsResponse](../../models/operations/listeventsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
