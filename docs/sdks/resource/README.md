# Resource

## Overview

A resource link represents a link between a shared resource (such as an kv store or config store) and a service version.

<https://developer.fastly.com/reference/api/services/resource>
### Available Operations

* [CreateResource](#createresource) - Create a resource link
* [DeleteResource](#deleteresource) - Delete a resource link
* [GetResource](#getresource) - Display a resource link
* [ListResources](#listresources) - List resource links
* [UpdateResource](#updateresource) - Update a resource link

## CreateResource

Create a link between a resource and a service version.

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
    operationSecurity := operations.CreateResourceSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Resource.CreateResource(ctx, operations.CreateResourceRequest{
        Resource: &shared.Resource{
            Name: sdk.String("test-resource"),
            ResourceID: sdk.String("3vjTN8v1O7nOAY7aNDGOL"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateResourceRequest](../../models/operations/createresourcerequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.CreateResourceSecurity](../../models/operations/createresourcesecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.CreateResourceResponse](../../models/operations/createresourceresponse.md), error**


## DeleteResource

Delete a link between a resource and a service version.

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
    operationSecurity := operations.DeleteResourceSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Resource.DeleteResource(ctx, operations.DeleteResourceRequest{
        ID: "7Lsb7Y76rChV9hSrv3KgFl",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteResource200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteResourceRequest](../../models/operations/deleteresourcerequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.DeleteResourceSecurity](../../models/operations/deleteresourcesecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.DeleteResourceResponse](../../models/operations/deleteresourceresponse.md), error**


## GetResource

Display a resource link by its identifier.

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
    operationSecurity := operations.GetResourceSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Resource.GetResource(ctx, operations.GetResourceRequest{
        ID: "7Lsb7Y76rChV9hSrv3KgFl",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetResourceRequest](../../models/operations/getresourcerequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.GetResourceSecurity](../../models/operations/getresourcesecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.GetResourceResponse](../../models/operations/getresourceresponse.md), error**


## ListResources

List links between resources and services

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
    operationSecurity := operations.ListResourcesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Resource.ListResources(ctx, operations.ListResourcesRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListResourcesRequest](../../models/operations/listresourcesrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListResourcesSecurity](../../models/operations/listresourcessecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.ListResourcesResponse](../../models/operations/listresourcesresponse.md), error**


## UpdateResource

Update a link between a resource and a service version.

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
    operationSecurity := operations.UpdateResourceSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Resource.UpdateResource(ctx, operations.UpdateResourceRequest{
        ID: "7Lsb7Y76rChV9hSrv3KgFl",
        Resource: &shared.Resource{
            Name: sdk.String("test-resource"),
            ResourceID: sdk.String("3vjTN8v1O7nOAY7aNDGOL"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateResourceRequest](../../models/operations/updateresourcerequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.UpdateResourceSecurity](../../models/operations/updateresourcesecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.UpdateResourceResponse](../../models/operations/updateresourceresponse.md), error**

