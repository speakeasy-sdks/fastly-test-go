# Pop
(*Pop*)

## Overview

List Fastly POPs and their locations.

<https://developer.fastly.com/reference/api/utils/pops>
### Available Operations

* [ListPops](#listpops) - List Fastly POPs

## ListPops

Get a list of all Fastly POPs.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/pkg/models/shared"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pop.ListPops(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pops != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListPopsResponse](../../models/operations/listpopsresponse.md), error**

