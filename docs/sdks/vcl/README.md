# Vcl

## Overview

A VCL is a Varnish configuration file used to customize the configuration for a Service.

<https://developer.fastly.com/reference/api/vcl-services/vcl>
### Available Operations

* [CreateCustomVcl](#createcustomvcl) - Create a custom VCL file
* [DeleteCustomVcl](#deletecustomvcl) - Delete a custom VCL file
* [GetCustomVcl](#getcustomvcl) - Get a custom VCL file
* [GetCustomVclBoilerplate](#getcustomvclboilerplate) - Get boilerplate VCL
* [GetCustomVclGenerated](#getcustomvclgenerated) - Get the generated VCL for a service
* [GetCustomVclRaw](#getcustomvclraw) - Download a custom VCL file
* [ListCustomVcl](#listcustomvcl) - List custom VCL files
* [SetCustomVclMain](#setcustomvclmain) - Set a custom VCL file as main
* [UpdateCustomVcl](#updatecustomvcl) - Update a custom VCL file

## CreateCustomVcl

Upload a VCL for a particular service and version.

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
    res, err := s.Vcl.CreateCustomVcl(ctx, operations.CreateCustomVclRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        Vcl: &shared.Vcl{
            Content: sdk.String("reiciendis"),
            Main: sdk.Bool(false),
            Name: sdk.String("test-vcl"),
        },
        VersionID: 1,
    }, operations.CreateCustomVclSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VclResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateCustomVclRequest](../../models/operations/createcustomvclrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.CreateCustomVclSecurity](../../models/operations/createcustomvclsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.CreateCustomVclResponse](../../models/operations/createcustomvclresponse.md), error**


## DeleteCustomVcl

Delete the uploaded VCL for a particular service and version.

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
    res, err := s.Vcl.DeleteCustomVcl(ctx, operations.DeleteCustomVclRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VclName: "test-vcl",
        VersionID: 1,
    }, operations.DeleteCustomVclSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteCustomVcl200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteCustomVclRequest](../../models/operations/deletecustomvclrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.DeleteCustomVclSecurity](../../models/operations/deletecustomvclsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.DeleteCustomVclResponse](../../models/operations/deletecustomvclresponse.md), error**


## GetCustomVcl

Get the uploaded VCL for a particular service and version.

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
    res, err := s.Vcl.GetCustomVcl(ctx, operations.GetCustomVclRequest{
        NoContent: sdk.String("vel"),
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VclName: "test-vcl",
        VersionID: 1,
    }, operations.GetCustomVclSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VclResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetCustomVclRequest](../../models/operations/getcustomvclrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.GetCustomVclSecurity](../../models/operations/getcustomvclsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.GetCustomVclResponse](../../models/operations/getcustomvclresponse.md), error**


## GetCustomVclBoilerplate

Return boilerplate VCL with the service's TTL from the [settings](/reference/api/vcl-services/settings/).

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
    res, err := s.Vcl.GetCustomVclBoilerplate(ctx, operations.GetCustomVclBoilerplateRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetCustomVclBoilerplateSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCustomVclBoilerplate200TextPlainString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCustomVclBoilerplateRequest](../../models/operations/getcustomvclboilerplaterequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.GetCustomVclBoilerplateSecurity](../../models/operations/getcustomvclboilerplatesecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.GetCustomVclBoilerplateResponse](../../models/operations/getcustomvclboilerplateresponse.md), error**


## GetCustomVclGenerated

Display the generated VCL for a particular service and version.

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
    res, err := s.Vcl.GetCustomVclGenerated(ctx, operations.GetCustomVclGeneratedRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetCustomVclGeneratedSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VclResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCustomVclGeneratedRequest](../../models/operations/getcustomvclgeneratedrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.GetCustomVclGeneratedSecurity](../../models/operations/getcustomvclgeneratedsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.GetCustomVclGeneratedResponse](../../models/operations/getcustomvclgeneratedresponse.md), error**


## GetCustomVclRaw

Download the specified VCL.

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
    res, err := s.Vcl.GetCustomVclRaw(ctx, operations.GetCustomVclRawRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VclName: "test-vcl",
        VersionID: 1,
    }, operations.GetCustomVclRawSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCustomVclRaw200TextPlainString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetCustomVclRawRequest](../../models/operations/getcustomvclrawrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.GetCustomVclRawSecurity](../../models/operations/getcustomvclrawsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.GetCustomVclRawResponse](../../models/operations/getcustomvclrawresponse.md), error**


## ListCustomVcl

List the uploaded VCLs for a particular service and version.

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
    res, err := s.Vcl.ListCustomVcl(ctx, operations.ListCustomVclRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListCustomVclSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VclResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListCustomVclRequest](../../models/operations/listcustomvclrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListCustomVclSecurity](../../models/operations/listcustomvclsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.ListCustomVclResponse](../../models/operations/listcustomvclresponse.md), error**


## SetCustomVclMain

Set the specified VCL as the main.

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
    res, err := s.Vcl.SetCustomVclMain(ctx, operations.SetCustomVclMainRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VclName: "test-vcl",
        VersionID: 1,
    }, operations.SetCustomVclMainSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VclResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.SetCustomVclMainRequest](../../models/operations/setcustomvclmainrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.SetCustomVclMainSecurity](../../models/operations/setcustomvclmainsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.SetCustomVclMainResponse](../../models/operations/setcustomvclmainresponse.md), error**


## UpdateCustomVcl

Update the uploaded VCL for a particular service and version.

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
    res, err := s.Vcl.UpdateCustomVcl(ctx, operations.UpdateCustomVclRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        Vcl: &shared.Vcl{
            Content: sdk.String("architecto"),
            Main: sdk.Bool(false),
            Name: sdk.String("test-vcl"),
        },
        VclName: "test-vcl",
        VersionID: 1,
    }, operations.UpdateCustomVclSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VclResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateCustomVclRequest](../../models/operations/updatecustomvclrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.UpdateCustomVclSecurity](../../models/operations/updatecustomvclsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.UpdateCustomVclResponse](../../models/operations/updatecustomvclresponse.md), error**

