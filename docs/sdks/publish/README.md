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

    ctx := context.Background()
    res, err := s.Publish.Publish(ctx, operations.PublishRequest{
        PublishRequest: &shared.PublishRequest{
            Items: []shared.PublishItem{
                shared.PublishItem{
                    Channel: "tempora",
                    Formats: shared.PublishItemFormats{
                        HTTPResponse: &shared.HTTPResponseFormat{
                            Body: sdk.String("tempora"),
                            BodyBin: sdk.String("voluptate"),
                            Code: sdk.Int64(970076),
                            Headers: map[string]string{
                                "sit": "non",
                                "officiis": "praesentium",
                            },
                            Reason: sdk.String("facilis"),
                        },
                        HTTPStream: &shared.HTTPStreamFormat{
                            Content: sdk.String("quaerat"),
                            ContentBin: sdk.String("incidunt"),
                        },
                        WsMessage: &shared.WsMessageFormat{
                            Content: sdk.String("ipsam"),
                            ContentBin: sdk.String("debitis"),
                        },
                    },
                    ID: sdk.String("80ca55ef-d20e-4457-a185-8b6a89fbe3a5"),
                    PrevID: sdk.String("officia"),
                },
                shared.PublishItem{
                    Channel: "dolorum",
                    Formats: shared.PublishItemFormats{
                        HTTPResponse: &shared.HTTPResponseFormat{
                            Body: sdk.String("corrupti"),
                            BodyBin: sdk.String("accusamus"),
                            Code: sdk.Int64(272683),
                            Headers: map[string]string{
                                "fugit": "ut",
                                "fugiat": "voluptatem",
                                "culpa": "expedita",
                            },
                            Reason: sdk.String("magnam"),
                        },
                        HTTPStream: &shared.HTTPStreamFormat{
                            Content: sdk.String("consequatur"),
                            ContentBin: sdk.String("esse"),
                        },
                        WsMessage: &shared.WsMessageFormat{
                            Content: sdk.String("ipsam"),
                            ContentBin: sdk.String("sit"),
                        },
                    },
                    ID: sdk.String("88e51862-065e-4904-b3b1-194b8abf603a"),
                    PrevID: sdk.String("voluptate"),
                },
                shared.PublishItem{
                    Channel: "unde",
                    Formats: shared.PublishItemFormats{
                        HTTPResponse: &shared.HTTPResponseFormat{
                            Body: sdk.String("reiciendis"),
                            BodyBin: sdk.String("provident"),
                            Code: sdk.Int64(833819),
                            Headers: map[string]string{
                                "voluptates": "perferendis",
                                "est": "quidem",
                                "reprehenderit": "facere",
                                "fuga": "praesentium",
                            },
                            Reason: sdk.String("mollitia"),
                        },
                        HTTPStream: &shared.HTTPStreamFormat{
                            Content: sdk.String("veniam"),
                            ContentBin: sdk.String("voluptatem"),
                        },
                        WsMessage: &shared.WsMessageFormat{
                            Content: sdk.String("quisquam"),
                            ContentBin: sdk.String("repudiandae"),
                        },
                    },
                    ID: sdk.String("187f86bc-173d-4689-aee9-526f8d986e88"),
                    PrevID: sdk.String("sunt"),
                },
                shared.PublishItem{
                    Channel: "recusandae",
                    Formats: shared.PublishItemFormats{
                        HTTPResponse: &shared.HTTPResponseFormat{
                            Body: sdk.String("dolorum"),
                            BodyBin: sdk.String("repellendus"),
                            Code: sdk.Int64(287119),
                            Headers: map[string]string{
                                "doloremque": "repudiandae",
                                "dicta": "accusantium",
                                "beatae": "dolores",
                                "enim": "laboriosam",
                            },
                            Reason: sdk.String("velit"),
                        },
                        HTTPStream: &shared.HTTPStreamFormat{
                            Content: sdk.String("a"),
                            ContentBin: sdk.String("molestias"),
                        },
                        WsMessage: &shared.WsMessageFormat{
                            Content: sdk.String("magnam"),
                            ContentBin: sdk.String("saepe"),
                        },
                    },
                    ID: sdk.String("29e973e9-22a5-47a1-9be3-e060807e2b6e"),
                    PrevID: sdk.String("ratione"),
                },
            },
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.PublishSecurity{
        Token: "",
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.PublishRequest](../../models/operations/publishrequest.md)   | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `security`                                                               | [operations.PublishSecurity](../../models/operations/publishsecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |


### Response

**[*operations.PublishResponse](../../models/operations/publishresponse.md), error**

