# Package
(*Package*)

## Overview

Compute@Edge is a computation platform capable of running custom binary packages that you compile on your own systems and upload to Fastly. These packages are associated with a service version and are deployed to Fastly's edge network.


<https://developer.fastly.com/reference/api/services/package>
### Available Operations

* [GetPackage](#getpackage) - Get details of the service's Compute@Edge package.
* [PutPackage](#putpackage) - Upload a Compute@Edge package.

## GetPackage

List detailed information about the Compute@Edge package for the specified service.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Package.GetPackage(ctx, operations.GetPackageRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PackageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetPackageRequest](../../models/operations/getpackagerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetPackageResponse](../../models/operations/getpackageresponse.md), error**


## PutPackage

Upload a Compute@Edge package associated with the specified service version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Package.PutPackage(ctx, operations.PutPackageRequest{
        Expect: fastly.String("voluptatum"),
        PackageUpload: &shared.PackageUpload{
            Package: &shared.PackageUploadPackage{
                Content: []byte("qui"),
                Package: "quibusdam",
            },
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PackageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.PutPackageRequest](../../models/operations/putpackagerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.PutPackageResponse](../../models/operations/putpackageresponse.md), error**

