# Service

## Overview

A Service represents the configuration for a website, app, API, or anything else to be served through Fastly. A Service can have many Versions, through which Backends, Domains, and more can be configured.

<https://developer.fastly.com/reference/api/services/service>
### Available Operations

* [CreateService](#createservice) - Create a service
* [DeleteService](#deleteservice) - Delete a service
* [GetService](#getservice) - Get a service
* [GetServiceDetail](#getservicedetail) - Get service details
* [ListServiceDomains](#listservicedomains) - List the domains within a service
* [ListServices](#listservices) - List services
* [SearchService](#searchservice) - Search for a service by name
* [UpdateService](#updateservice) - Update a service

## CreateService

Create a service.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Service.CreateService(ctx, shared.ServiceCreate1{
        Comment: sdk.String("recusandae"),
        CustomerID: sdk.String("x4xCwxxJxGCx123Rx5xTx"),
        Name: sdk.String("test-service"),
        Type: shared.ServiceCreateTypeVcl.ToPointer(),
    }, operations.CreateServiceSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.ServiceCreate1](../../models/shared/servicecreate1.md)                       | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.CreateServiceSecurity](../../models/operations/createservicesecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.CreateServiceResponse](../../models/operations/createserviceresponse.md), error**


## DeleteService

Delete a service.

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

    ctx := context.Background()
    res, err := s.Service.DeleteService(ctx, operations.DeleteServiceRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.DeleteServiceSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteService200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteServiceRequest](../../models/operations/deleteservicerequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.DeleteServiceSecurity](../../models/operations/deleteservicesecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.DeleteServiceResponse](../../models/operations/deleteserviceresponse.md), error**


## GetService

Get a specific service by id.

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

    ctx := context.Background()
    res, err := s.Service.GetService(ctx, operations.GetServiceRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.GetServiceSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetServiceRequest](../../models/operations/getservicerequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.GetServiceSecurity](../../models/operations/getservicesecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.GetServiceResponse](../../models/operations/getserviceresponse.md), error**


## GetServiceDetail

List detailed information on a specified service.

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

    ctx := context.Background()
    res, err := s.Service.GetServiceDetail(ctx, operations.GetServiceDetailRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        Version: sdk.Int64(1),
    }, operations.GetServiceDetailSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServiceDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetServiceDetailRequest](../../models/operations/getservicedetailrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.GetServiceDetailSecurity](../../models/operations/getservicedetailsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.GetServiceDetailResponse](../../models/operations/getservicedetailresponse.md), error**


## ListServiceDomains

List the domains within a service.

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

    ctx := context.Background()
    res, err := s.Service.ListServiceDomains(ctx, operations.ListServiceDomainsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.ListServiceDomainsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DomainResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListServiceDomainsRequest](../../models/operations/listservicedomainsrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ListServiceDomainsSecurity](../../models/operations/listservicedomainssecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ListServiceDomainsResponse](../../models/operations/listservicedomainsresponse.md), error**


## ListServices

List services.

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

    ctx := context.Background()
    res, err := s.Service.ListServices(ctx, operations.ListServicesRequest{
        Direction: shared.DirectionAscend.ToPointer(),
        Page: sdk.Int64(1),
        PerPage: sdk.Int64(20),
        Sort: sdk.String("created"),
    }, operations.ListServicesSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServiceListResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListServicesRequest](../../models/operations/listservicesrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.ListServicesSecurity](../../models/operations/listservicessecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.ListServicesResponse](../../models/operations/listservicesresponse.md), error**


## SearchService

Get a specific service by name.

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

    ctx := context.Background()
    res, err := s.Service.SearchService(ctx, operations.SearchServiceRequest{
        Name: "test-service",
    }, operations.SearchServiceSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.SearchServiceRequest](../../models/operations/searchservicerequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.SearchServiceSecurity](../../models/operations/searchservicesecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.SearchServiceResponse](../../models/operations/searchserviceresponse.md), error**


## UpdateService

Update a service.

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

    ctx := context.Background()
    res, err := s.Service.UpdateService(ctx, operations.UpdateServiceRequest{
        Service: &shared.Service{
            Comment: sdk.String("provident"),
            CustomerID: sdk.String("x4xCwxxJxGCx123Rx5xTx"),
            Name: sdk.String("test-service"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operations.UpdateServiceSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateServiceRequest](../../models/operations/updateservicerequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.UpdateServiceSecurity](../../models/operations/updateservicesecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.UpdateServiceResponse](../../models/operations/updateserviceresponse.md), error**

