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
                    Channel: "what Lira Ohio",
                    Formats: shared.PublishItemFormats{
                        HTTPResponse: &shared.HTTPResponseFormat{
                            Body: fastly.String("bypassing Southwest er"),
                            BodyBin: fastly.String("Clarksville frictionless"),
                            Code: fastly.Int64(222146),
                            Headers: map[string]string{
                                "eos": "joshingly",
                            },
                            Reason: fastly.String("user Ambrose infrastructure"),
                        },
                        HTTPStream: &shared.HTTPStreamFormat{
                            Content: fastly.String("Buckinghamshire"),
                            ContentBin: fastly.String("withdrawal"),
                        },
                        WsMessage: &shared.WsMessageFormat{
                            Content: fastly.String("Electric"),
                            ContentBin: fastly.String("Oriental"),
                        },
                    },
                    ID: fastly.String("<ID>"),
                    PrevID: fastly.String("ferociously next"),
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

