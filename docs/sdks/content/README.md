# Content

## Overview

Fastly makes it possible to see which version of a particular URL is cached on each edge server.

<https://developer.fastly.com/reference/api/utils/content>
### Available Operations

* [ContentCheck](#contentcheck) - Check status of content in each POP's cache

## ContentCheck

Retrieve headers and MD5 hash of the content for a particular URL from each Fastly edge server. This API is limited to 200 requests per hour.

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
    operationSecurity := operations.ContentCheckSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Content.ContentCheck(ctx, operations.ContentCheckRequest{
        URL: sdk.String("https://www.example.com/foo/bar"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Contents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ContentCheckRequest](../../models/operations/contentcheckrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.ContentCheckSecurity](../../models/operations/contentchecksecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.ContentCheckResponse](../../models/operations/contentcheckresponse.md), error**

