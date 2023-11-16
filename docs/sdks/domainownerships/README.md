# DomainOwnerships
(*DomainOwnerships*)

## Overview

A domain ownership is the record that Fastly keeps after a customer has validated control over a domain.

<https://developer.fastly.com/reference/api/services/domain-ownerships>
### Available Operations

* [ListDomainOwnerships](#listdomainownerships) - List domain-ownerships

## ListDomainOwnerships

List all domain-ownerships.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.DomainOwnerships.ListDomainOwnerships(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListDomainOwnershipsResponse](../../models/operations/listdomainownershipsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
