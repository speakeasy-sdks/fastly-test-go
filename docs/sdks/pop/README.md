# Pop

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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListPopsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Pop.ListPops(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pops != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `security`                                                                 | [operations.ListPopsSecurity](../../models/operations/listpopssecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.ListPopsResponse](../../models/operations/listpopsresponse.md), error**

