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
    res, err := s.Healthcheck.CreateHealthcheck(ctx, operations.CreateHealthcheckRequest{
        Healthcheck: &shared.Healthcheck{
            CheckInterval: fastly.Int64(919483),
            Comment: fastly.String("ullam"),
            ExpectedResponse: fastly.Int64(714242),
            Headers: []string{
                "nihil",
            },
            Host: fastly.String("repellat"),
            HTTPVersion: fastly.String("quibusdam"),
            Initial: fastly.Int64(149448),
            Method: fastly.String("saepe"),
            Name: fastly.String("test-healthcheck"),
            Path: fastly.String("pariatur"),
            Threshold: fastly.Int64(37559),
            Timeout: fastly.Int64(162493),
            Window: fastly.Int64(508315),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HealthcheckResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateHealthcheckRequest](../../models/operations/createhealthcheckrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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
    res, err := s.Healthcheck.DeleteHealthcheck(ctx, operations.DeleteHealthcheckRequest{
        HealthcheckName: "test-healthcheck",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteHealthcheck200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteHealthcheckRequest](../../models/operations/deletehealthcheckrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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
    res, err := s.Healthcheck.GetHealthcheck(ctx, operations.GetHealthcheckRequest{
        HealthcheckName: "test-healthcheck",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HealthcheckResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetHealthcheckRequest](../../models/operations/gethealthcheckrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
    res, err := s.Healthcheck.ListHealthchecks(ctx, operations.ListHealthchecksRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HealthcheckResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListHealthchecksRequest](../../models/operations/listhealthchecksrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
    res, err := s.Healthcheck.UpdateHealthcheck(ctx, operations.UpdateHealthcheckRequest{
        Healthcheck: &shared.Healthcheck{
            CheckInterval: fastly.Int64(615560),
            Comment: fastly.String("magni"),
            ExpectedResponse: fastly.Int64(123820),
            Headers: []string{
                "quo",
            },
            Host: fastly.String("illum"),
            HTTPVersion: fastly.String("pariatur"),
            Initial: fastly.Int64(807319),
            Method: fastly.String("ea"),
            Name: fastly.String("test-healthcheck"),
            Path: fastly.String("excepturi"),
            Threshold: fastly.Int64(139972),
            Timeout: fastly.Int64(407183),
            Window: fastly.Int64(33222),
        },
        HealthcheckName: "test-healthcheck",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HealthcheckResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateHealthcheckRequest](../../models/operations/updatehealthcheckrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateHealthcheckResponse](../../models/operations/updatehealthcheckresponse.md), error**

