# VclDiff
(*.VclDiff*)

## Overview

Compare the changes in generated VCL between two versions of a service. This is sometimes called a "diff" because the comparison may highlight "differences" between the versions. To compare the configuration changes between two versions of a service represented in YAML format instead, use the related [diff](/reference/api/utils/diff/#diff-service-versions) endpoint.

<https://developer.fastly.com/reference/api/vcl-services/diff>
### Available Operations

* [VclDiffServiceVersions](#vcldiffserviceversions) - Get a comparison of the VCL changes between two service versions

## VclDiffServiceVersions

Get a comparison of the VCL changes between two service versions.

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
    res, err := s.VclDiff.VclDiffServiceVersions(ctx, operations.VclDiffServiceVersionsRequest{
        FromVersionID: 1,
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        ToVersionID: 2,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VclDiff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.VclDiffServiceVersionsRequest](../../models/operations/vcldiffserviceversionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.VclDiffServiceVersionsResponse](../../models/operations/vcldiffserviceversionsresponse.md), error**

