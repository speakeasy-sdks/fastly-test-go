# Domain

## Overview

A domain represents the domain name through which visitors will retrieve content. There can be multiple domains for a service.

<https://developer.fastly.com/reference/api/services/domain>
### Available Operations

* [CheckDomain](#checkdomain) - Validate DNS configuration for a single domain on a service
* [CheckDomains](#checkdomains) - Validate DNS configuration for all domains on a service
* [CreateDomain](#createdomain) - Add a domain name to a service
* [DeleteDomain](#deletedomain) - Remove a domain from a service
* [GetDomain](#getdomain) - Describe a domain
* [ListDomains](#listdomains) - List domains
* [UpdateDomain](#updatedomain) - Update a domain

## CheckDomain

Checks the status of a specific domain's DNS record for a Service Version. Returns an array in the same format as domain/check_all.

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
    operationSecurity := operations.CheckDomainSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Domain.CheckDomain(ctx, operations.CheckDomainRequest{
        DomainName: "www.example.com",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DomainCheckResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CheckDomainRequest](../../models/operations/checkdomainrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.CheckDomainSecurity](../../models/operations/checkdomainsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.CheckDomainResponse](../../models/operations/checkdomainresponse.md), error**


## CheckDomains

Checks the status of all domains' DNS records for a Service Version. Returns an array of 3 items for each domain; the first is the details for the domain, the second is the current CNAME of the domain, and the third is a boolean indicating whether or not it has been properly setup to use Fastly.

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
    operationSecurity := operations.CheckDomainsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Domain.CheckDomains(ctx, operations.CheckDomainsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DomainCheckResponseList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CheckDomainsRequest](../../models/operations/checkdomainsrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.CheckDomainsSecurity](../../models/operations/checkdomainssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.CheckDomainsResponse](../../models/operations/checkdomainsresponse.md), error**


## CreateDomain

Create a domain for a particular service and version.

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
    operationSecurity := operations.CreateDomainSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Domain.CreateDomain(ctx, operations.CreateDomainRequest{
        Domain: &shared.Domain{
            Comment: sdk.String("deleniti"),
            Name: sdk.String("www.example.com"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DomainResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateDomainRequest](../../models/operations/createdomainrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.CreateDomainSecurity](../../models/operations/createdomainsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.CreateDomainResponse](../../models/operations/createdomainresponse.md), error**


## DeleteDomain

Delete the domain for a particular service and versions.

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
    operationSecurity := operations.DeleteDomainSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Domain.DeleteDomain(ctx, operations.DeleteDomainRequest{
        DomainName: "www.example.com",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDomain200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteDomainRequest](../../models/operations/deletedomainrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.DeleteDomainSecurity](../../models/operations/deletedomainsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.DeleteDomainResponse](../../models/operations/deletedomainresponse.md), error**


## GetDomain

Get the domain for a particular service and version.

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
    operationSecurity := operations.GetDomainSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Domain.GetDomain(ctx, operations.GetDomainRequest{
        DomainName: "www.example.com",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DomainResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetDomainRequest](../../models/operations/getdomainrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.GetDomainSecurity](../../models/operations/getdomainsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.GetDomainResponse](../../models/operations/getdomainresponse.md), error**


## ListDomains

List all the domains for a particular service and version.

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
    operationSecurity := operations.ListDomainsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Domain.ListDomains(ctx, operations.ListDomainsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DomainsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListDomainsRequest](../../models/operations/listdomainsrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.ListDomainsSecurity](../../models/operations/listdomainssecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.ListDomainsResponse](../../models/operations/listdomainsresponse.md), error**


## UpdateDomain

Update the domain for a particular service and version.

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
    operationSecurity := operations.UpdateDomainSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Domain.UpdateDomain(ctx, operations.UpdateDomainRequest{
        Domain: &shared.Domain{
            Comment: sdk.String("omnis"),
            Name: sdk.String("www.example.com"),
        },
        DomainName: "www.example.com",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DomainResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateDomainRequest](../../models/operations/updatedomainrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.UpdateDomainSecurity](../../models/operations/updatedomainsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.UpdateDomainResponse](../../models/operations/updatedomainresponse.md), error**

