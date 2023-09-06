# Publish

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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.PublishSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Publish.Publish(ctx, operations.PublishRequest{
        PublishRequest: &shared.PublishRequest{
            Items: []shared.PublishItem{
                shared.PublishItem{
                    Channel: "non",
                    Formats: shared.PublishItemFormats{
                        HTTPResponse: &shared.HTTPResponseFormat{
                            Body: sdk.String("voluptatem"),
                            BodyBin: sdk.String("dolor"),
                            Code: sdk.Int64(580152),
                            Headers: map[string]string{
                                "numquam": "impedit",
                            },
                            Reason: sdk.String("explicabo"),
                        },
                        HTTPStream: &shared.HTTPStreamFormat{
                            Content: sdk.String("voluptas"),
                            ContentBin: sdk.String("aut"),
                        },
                        WsMessage: &shared.WsMessageFormat{
                            Content: sdk.String("dignissimos"),
                            ContentBin: sdk.String("dicta"),
                        },
                    },
                    ID: sdk.String("f93f5f06-42da-4c7a-b515-cc413aa63aae"),
                    PrevID: sdk.String("totam"),
                },
            },
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Publish200TextPlainString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.PublishRequest](../../models/operations/publishrequest.md)   | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `security`                                                               | [operations.PublishSecurity](../../models/operations/publishsecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |


### Response

**[*operations.PublishResponse](../../models/operations/publishresponse.md), error**

