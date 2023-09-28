# Publish
(*Publish*)

## Overview

Publishing sends messages to [Fanout](https://developer.fastly.com/learning/concepts/real-time-messaging/fanout) subscribers. Fanout is designed to be [GRIP-compatible](https://pushpin.org/docs/protocols/grip/), such that `https://api.fastly.com/service/{service_id}` can be used as a GRIP URL in application configurations.

<https://developer.fastly.com/reference/api/services/service>
### Available Operations

* [Publish](#publish) - Send messages to Fanout subscribers

## Publish

Send one or more messages to [Fanout](https://developer.fastly.com/learning/concepts/real-time-messaging/fanout) subscribers. Each message specifies a channel, and Fanout will deliver the message to all subscribers of its channel.
> **IMPORTANT:** For compatibility with GRIP, this endpoint requires a trailing slash, and the API token may be provided in the `Authorization` header (instead of the `Fastly-Key` header) using the `Bearer` scheme.


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
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Publish.Publish(ctx, operations.PublishRequest{
        PublishRequest: &shared.PublishRequest{
            Items: []shared.PublishItem{
                shared.PublishItem{
                    Channel: "natus",
                    Formats: shared.PublishItemFormats{
                        HTTPResponse: &shared.HTTPResponseFormat{
                            Body: fastly.String("velit"),
                            BodyBin: fastly.String("voluptatibus"),
                            Code: fastly.Int64(374323),
                            Headers: map[string]string{
                                "asperiores": "aperiam",
                            },
                            Reason: fastly.String("ea"),
                        },
                        HTTPStream: &shared.HTTPStreamFormat{
                            Content: fastly.String("quaerat"),
                            ContentBin: fastly.String("consequuntur"),
                        },
                        WsMessage: &shared.WsMessageFormat{
                            Content: fastly.String("repellendus"),
                            ContentBin: fastly.String("officia"),
                        },
                    },
                    ID: fastly.String("c7af515c-c413-4aa6-baae-8d67864dbb67"),
                    PrevID: fastly.String("corporis"),
                },
            },
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Publish200TextPlainString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.PublishRequest](../../models/operations/publishrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.PublishResponse](../../models/operations/publishresponse.md), error**

