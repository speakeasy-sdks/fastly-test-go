# Healthcheck

## Overview

Health checks are used to customize the way Fastly checks on your Backends. If an origin server is marked unhealthy due to health checks, Fastly will stop attempting to send requests to it. If all origin servers are marked unhealthy, Fastly will attempt to serve stale. If no stale object is available, a 503 will be returned to the client.

<https://developer.fastly.com/reference/api/services/healthcheck>
### Available Operations

* [CreateHealthcheck](#createhealthcheck) - Create a health check
* [DeleteHealthcheck](#deletehealthcheck) - Delete a health check
* [GetHealthcheck](#gethealthcheck) - Get a health check
* [ListHealthchecks](#listhealthchecks) - List health checks
* [UpdateHealthcheck](#updatehealthcheck) - Update a health check

## CreateHealthcheck

Create a health check for a particular service and version.

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
    operationSecurity := operations.CreateHealthcheckSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Healthcheck.CreateHealthcheck(ctx, operations.CreateHealthcheckRequest{
        Healthcheck: &shared.Healthcheck{
            CheckInterval: sdk.Int64(703889),
            Comment: sdk.String("in"),
            ExpectedResponse: sdk.Int64(100226),
            Headers: []string{
                "architecto",
            },
            Host: sdk.String("repudiandae"),
            HTTPVersion: sdk.String("ullam"),
            Initial: sdk.Int64(714242),
            Method: sdk.String("nihil"),
            Name: sdk.String("test-healthcheck"),
            Path: sdk.String("repellat"),
            Threshold: sdk.Int64(841140),
            Timeout: sdk.Int64(149448),
            Window: sdk.Int64(904648),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HealthcheckResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateHealthcheckRequest](../../models/operations/createhealthcheckrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.CreateHealthcheckSecurity](../../models/operations/createhealthchecksecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.CreateHealthcheckResponse](../../models/operations/createhealthcheckresponse.md), error**


## DeleteHealthcheck

Delete the health check for a particular service and version.

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
    operationSecurity := operations.DeleteHealthcheckSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Healthcheck.DeleteHealthcheck(ctx, operations.DeleteHealthcheckRequest{
        HealthcheckName: "test-healthcheck",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteHealthcheck200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteHealthcheckRequest](../../models/operations/deletehealthcheckrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.DeleteHealthcheckSecurity](../../models/operations/deletehealthchecksecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.DeleteHealthcheckResponse](../../models/operations/deletehealthcheckresponse.md), error**


## GetHealthcheck

Get the health check for a particular service and version.

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
    operationSecurity := operations.GetHealthcheckSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Healthcheck.GetHealthcheck(ctx, operations.GetHealthcheckRequest{
        HealthcheckName: "test-healthcheck",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HealthcheckResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetHealthcheckRequest](../../models/operations/gethealthcheckrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetHealthcheckSecurity](../../models/operations/gethealthchecksecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetHealthcheckResponse](../../models/operations/gethealthcheckresponse.md), error**


## ListHealthchecks

List all of the health checks for a particular service and version.

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
    operationSecurity := operations.ListHealthchecksSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Healthcheck.ListHealthchecks(ctx, operations.ListHealthchecksRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HealthcheckResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListHealthchecksRequest](../../models/operations/listhealthchecksrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ListHealthchecksSecurity](../../models/operations/listhealthcheckssecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ListHealthchecksResponse](../../models/operations/listhealthchecksresponse.md), error**


## UpdateHealthcheck

Update the health check for a particular service and version.

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
    operationSecurity := operations.UpdateHealthcheckSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.Healthcheck.UpdateHealthcheck(ctx, operations.UpdateHealthcheckRequest{
        Healthcheck: &shared.Healthcheck{
            CheckInterval: sdk.Int64(868126),
            Comment: sdk.String("accusantium"),
            ExpectedResponse: sdk.Int64(162493),
            Headers: []string{
                "praesentium",
            },
            Host: sdk.String("natus"),
            HTTPVersion: sdk.String("magni"),
            Initial: sdk.Int64(123820),
            Method: sdk.String("quo"),
            Name: sdk.String("test-healthcheck"),
            Path: sdk.String("illum"),
            Threshold: sdk.Int64(864934),
            Timeout: sdk.Int64(807319),
            Window: sdk.Int64(411397),
        },
        HealthcheckName: "test-healthcheck",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HealthcheckResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateHealthcheckRequest](../../models/operations/updatehealthcheckrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UpdateHealthcheckSecurity](../../models/operations/updatehealthchecksecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.UpdateHealthcheckResponse](../../models/operations/updatehealthcheckresponse.md), error**

