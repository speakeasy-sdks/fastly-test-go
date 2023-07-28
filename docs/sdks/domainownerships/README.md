# DomainOwnerships

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
	"context"
	"log"
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListDomainOwnershipsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.DomainOwnerships.ListDomainOwnerships(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListDomainOwnerships200ApplicationVndAPIPlusJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `security`                                                                                         | [operations.ListDomainOwnershipsSecurity](../../models/operations/listdomainownershipssecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.ListDomainOwnershipsResponse](../../models/operations/listdomainownershipsresponse.md), error**

