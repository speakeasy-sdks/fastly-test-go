# ApexRedirect

## Overview

Supports redirecting traffic for apex domains, subdomains, or wildcard domains to a WWW subdomain.

<https://developer.fastly.com/reference/api/vcl-services/apex-redirect>
### Available Operations

* [DeleteApexRedirect](#deleteapexredirect) - Delete an apex redirect
* [GetApexRedirect](#getapexredirect) - Get an apex redirect
* [ListApexRedirects](#listapexredirects) - List apex redirects
* [UpdateApexRedirect](#updateapexredirect) - Update an apex redirect

## DeleteApexRedirect

Delete an apex redirect by its ID.

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
    operationSecurity := operations.DeleteApexRedirectSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.ApexRedirect.DeleteApexRedirect(ctx, operations.DeleteApexRedirectRequest{
        ApexRedirectID: "debitis",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteApexRedirect200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteApexRedirectRequest](../../models/operations/deleteapexredirectrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.DeleteApexRedirectSecurity](../../models/operations/deleteapexredirectsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.DeleteApexRedirectResponse](../../models/operations/deleteapexredirectresponse.md), error**


## GetApexRedirect

Get an apex redirect by its ID.

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
    operationSecurity := operations.GetApexRedirectSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.ApexRedirect.GetApexRedirect(ctx, operations.GetApexRedirectRequest{
        ApexRedirectID: "ipsa",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApexRedirect != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetApexRedirectRequest](../../models/operations/getapexredirectrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.GetApexRedirectSecurity](../../models/operations/getapexredirectsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.GetApexRedirectResponse](../../models/operations/getapexredirectresponse.md), error**


## ListApexRedirects

List all apex redirects for a particular service and version.

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
    operationSecurity := operations.ListApexRedirectsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.ApexRedirect.ListApexRedirects(ctx, operations.ListApexRedirectsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApexRedirects != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListApexRedirectsRequest](../../models/operations/listapexredirectsrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ListApexRedirectsSecurity](../../models/operations/listapexredirectssecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.ListApexRedirectsResponse](../../models/operations/listapexredirectsresponse.md), error**


## UpdateApexRedirect

Update an apex redirect by its ID.

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
    operationSecurity := operations.UpdateApexRedirectSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.ApexRedirect.UpdateApexRedirect(ctx, operations.UpdateApexRedirectRequest{
        ApexRedirectInput: &shared.ApexRedirectInput{
            Domains: []string{
                "tempora",
                "suscipit",
                "molestiae",
                "minus",
            },
            FeatureRevision: sdk.Int64(812169),
            StatusCode: shared.ApexRedirectStatusCodeThreeHundredAndSeven.ToPointer(),
        },
        ApexRedirectID: "iusto",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApexRedirect != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateApexRedirectRequest](../../models/operations/updateapexredirectrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.UpdateApexRedirectSecurity](../../models/operations/updateapexredirectsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.UpdateApexRedirectResponse](../../models/operations/updateapexredirectresponse.md), error**

