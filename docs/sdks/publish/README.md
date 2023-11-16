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
    res, err := s.Publish.Publish(ctx, operations.PublishRequest{
        PublishRequest: &components.PublishRequest{
            Items: []components.PublishItem{
                components.PublishItem{
                    Channel: "string",
                    Formats: components.Formats{
                        HTTPResponse: &components.HTTPResponseFormat{
                            Headers: map[string]string{
                                "key": "string",
                            },
                        },
                        HTTPStream: &components.HTTPStreamFormat{},
                        WsMessage: &components.WsMessageFormat{},
                    },
                },
            },
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
