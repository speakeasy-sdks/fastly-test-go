# TLSDomains

## Overview

TLS domains are all the domains (including wildcard domains) included in any [TLS certificate](#tls_certificates)'s Subject Alternative Names (SAN) list. Included in the response is information about which certificates reference this domain as well as the [TLS activation](#tls_activations) indicating which certificate is enabled to serve TLS traffic for the domain.

<https://developer.fastly.com/reference/api/tls/custom-certs/domains>
### Available Operations

* [ListTLSDomains](#listtlsdomains) - List TLS domains

## ListTLSDomains

List all TLS domains.

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
    operationSecurity := operations.ListTLSDomainsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSDomains.ListTLSDomains(ctx, operations.ListTLSDomainsRequest{
        FilterInUse: sdk.String("dicta"),
        FilterTLSCertificatesID: sdk.String("minima"),
        FilterTLSSubscriptionsID: sdk.String("beatae"),
        Include: sdk.String("cupiditate"),
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
        Sort: shared.SortMinusCreatedAt.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSDomainsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListTLSDomainsRequest](../../models/operations/listtlsdomainsrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.ListTLSDomainsSecurity](../../models/operations/listtlsdomainssecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.ListTLSDomainsResponse](../../models/operations/listtlsdomainsresponse.md), error**

