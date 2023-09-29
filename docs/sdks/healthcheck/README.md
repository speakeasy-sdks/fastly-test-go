# Healthcheck
(*Healthcheck*)

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
            CheckInterval: fastly.Int64(894822),
            Comment: fastly.String("The automobile layout consists of a front-engine design, with transaxle-type transmissions mounted at the rear of the engine and four wheel drive"),
            ExpectedResponse: fastly.Int64(563880),
            Headers: []string{
                "synergistic",
            },
            Host: fastly.String("glistening-concept.biz"),
            HTTPVersion: fastly.String("meh Caguas"),
            Initial: fastly.Int64(288154),
            Method: fastly.String("whenever"),
            Name: fastly.String("test-healthcheck"),
            Path: fastly.String("/opt/share"),
            Threshold: fastly.Int64(837733),
            Timeout: fastly.Int64(697295),
            Window: fastly.Int64(644894),
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
            CheckInterval: fastly.Int64(931535),
            Comment: fastly.String("Carbonite web goalkeeper gloves are ergonomically designed to give easy fit"),
            ExpectedResponse: fastly.Int64(848950),
            Headers: []string{
                "Branding",
            },
            Host: fastly.String("past-armament.org"),
            HTTPVersion: fastly.String("Canadian"),
            Initial: fastly.Int64(693029),
            Method: fastly.String("Nissan"),
            Name: fastly.String("test-healthcheck"),
            Path: fastly.String("/root"),
            Threshold: fastly.Int64(324473),
            Timeout: fastly.Int64(242944),
            Window: fastly.Int64(455756),
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

