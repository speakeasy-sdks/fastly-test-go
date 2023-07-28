# DictionaryInfo

## Overview

Dictionary Info is a set of metadata describing properties of a dictionary which change as items are added and removed.

<https://developer.fastly.com/reference/api/dictionaries/dictionary-info>
### Available Operations

* [GetDictionaryInfo](#getdictionaryinfo) - Get edge dictionary metadata

## GetDictionaryInfo

Retrieve metadata for a single dictionary by ID for a version and service.

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
    operationSecurity := operations.GetDictionaryInfoSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.DictionaryInfo.GetDictionaryInfo(ctx, operations.GetDictionaryInfoRequest{
        DictionaryID: "3vjTN8v1O7nOAY7aNDGOL",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DictionaryInfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetDictionaryInfoRequest](../../models/operations/getdictionaryinforequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.GetDictionaryInfoSecurity](../../models/operations/getdictionaryinfosecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.GetDictionaryInfoResponse](../../models/operations/getdictionaryinforesponse.md), error**

