# Version

## Overview

A Version represents a specific instance of the configuration for a service. A Version can be cloned, locked, activated, or deactivated.

<https://developer.fastly.com/reference/api/services/version>
### Available Operations

* [ActivateServiceVersion](#activateserviceversion) - Activate a service version
* [CloneServiceVersion](#cloneserviceversion) - Clone a service version
* [CreateServiceVersion](#createserviceversion) - Create a service version
* [DeactivateServiceVersion](#deactivateserviceversion) - Deactivate a service version
* [GetServiceVersion](#getserviceversion) - Get a version of a service
* [ListServiceVersions](#listserviceversions) - List versions of a service
* [LockServiceVersion](#lockserviceversion) - Lock a service version
* [UpdateServiceVersion](#updateserviceversion) - Update a service version
* [ValidateServiceVersion](#validateserviceversion) - Validate a service version

## ActivateServiceVersion

Activate the current version.

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
    operationSecurity := operations.ActivateServiceVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Version.ActivateServiceVersion(ctx, operations.ActivateServiceVersionRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ActivateServiceVersionRequest](../../models/operations/activateserviceversionrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.ActivateServiceVersionSecurity](../../models/operations/activateserviceversionsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.ActivateServiceVersionResponse](../../models/operations/activateserviceversionresponse.md), error**


## CloneServiceVersion

Clone the current configuration into a new version.

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
    operationSecurity := operations.CloneServiceVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Version.CloneServiceVersion(ctx, operations.CloneServiceVersionRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Version != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CloneServiceVersionRequest](../../models/operations/cloneserviceversionrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.CloneServiceVersionSecurity](../../models/operations/cloneserviceversionsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.CloneServiceVersionResponse](../../models/operations/cloneserviceversionresponse.md), error**


## CreateServiceVersion

Create a version for a particular service.

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
    operationSecurity := operations.CreateServiceVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Version.CreateServiceVersion(ctx, operations.CreateServiceVersionRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateServiceVersionRequest](../../models/operations/createserviceversionrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.CreateServiceVersionSecurity](../../models/operations/createserviceversionsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.CreateServiceVersionResponse](../../models/operations/createserviceversionresponse.md), error**


## DeactivateServiceVersion

Deactivate the current version.

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
    operationSecurity := operations.DeactivateServiceVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Version.DeactivateServiceVersion(ctx, operations.DeactivateServiceVersionRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.DeactivateServiceVersionRequest](../../models/operations/deactivateserviceversionrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.DeactivateServiceVersionSecurity](../../models/operations/deactivateserviceversionsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.DeactivateServiceVersionResponse](../../models/operations/deactivateserviceversionresponse.md), error**


## GetServiceVersion

Get the version for a particular service.

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
    operationSecurity := operations.GetServiceVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Version.GetServiceVersion(ctx, operations.GetServiceVersionRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetServiceVersionRequest](../../models/operations/getserviceversionrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.GetServiceVersionSecurity](../../models/operations/getserviceversionsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.GetServiceVersionResponse](../../models/operations/getserviceversionresponse.md), error**


## ListServiceVersions

List the versions for a particular service.

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
    operationSecurity := operations.ListServiceVersionsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Version.ListServiceVersions(ctx, operations.ListServiceVersionsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListServiceVersionsRequest](../../models/operations/listserviceversionsrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.ListServiceVersionsSecurity](../../models/operations/listserviceversionssecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.ListServiceVersionsResponse](../../models/operations/listserviceversionsresponse.md), error**


## LockServiceVersion

Locks the specified version.

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
    operationSecurity := operations.LockServiceVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Version.LockServiceVersion(ctx, operations.LockServiceVersionRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Version != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.LockServiceVersionRequest](../../models/operations/lockserviceversionrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.LockServiceVersionSecurity](../../models/operations/lockserviceversionsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.LockServiceVersionResponse](../../models/operations/lockserviceversionresponse.md), error**


## UpdateServiceVersion

Update a particular version for a particular service.

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
    operationSecurity := operations.UpdateServiceVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Version.UpdateServiceVersion(ctx, operations.UpdateServiceVersionRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionInput: &shared.VersionInput{
            Active: sdk.Bool(false),
            Comment: sdk.String("reiciendis"),
            Deployed: sdk.Bool(false),
            Locked: sdk.Bool(false),
            Staging: sdk.Bool(false),
            Testing: sdk.Bool(false),
        },
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateServiceVersionRequest](../../models/operations/updateserviceversionrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.UpdateServiceVersionSecurity](../../models/operations/updateserviceversionsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.UpdateServiceVersionResponse](../../models/operations/updateserviceversionresponse.md), error**


## ValidateServiceVersion

Validate the version for a particular service and version.

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
    operationSecurity := operations.ValidateServiceVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Version.ValidateServiceVersion(ctx, operations.ValidateServiceVersionRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ValidateServiceVersion200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ValidateServiceVersionRequest](../../models/operations/validateserviceversionrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.ValidateServiceVersionSecurity](../../models/operations/validateserviceversionsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.ValidateServiceVersionResponse](../../models/operations/validateserviceversionresponse.md), error**

