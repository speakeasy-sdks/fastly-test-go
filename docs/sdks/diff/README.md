# Diff
(*Diff*)

## Overview

See the line-by-line changes in configuration between two different versions of a service.

<https://developer.fastly.com/reference/api/utils/diff>
### Available Operations

* [DiffServiceVersions](#diffserviceversions) - Diff two service versions

## DiffServiceVersions

Get diff between two versions.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Diff.DiffServiceVersions(ctx, operations.DiffServiceVersionsRequest{
        FromVersionID: 1,
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        ToVersionID: 2,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DiffResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DiffServiceVersionsRequest](../../models/operations/diffserviceversionsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DiffServiceVersionsResponse](../../models/operations/diffserviceversionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
