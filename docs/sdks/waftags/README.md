# WafTags
(*.WafTags*)

## Overview

Tags for categorizing WAF [rules](/reference/api/waf/rules/).

<https://developer.fastly.com/reference/api/waf/tags>
### Available Operations

* [~~ListWafTags~~](#listwaftags) - List tags :warning: **Deprecated**

## ~~ListWafTags~~

List all tags.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.WafTags.ListWafTags(ctx, operations.ListWafTagsRequest{
        Include: components.WafTagIncludeWafRules.ToPointer(),
        PageNumber: fastly.Int64(1),
        PageSize: fastly.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafTagsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListWafTagsRequest](../../models/operations/listwaftagsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ListWafTagsResponse](../../models/operations/listwaftagsresponse.md), error**

