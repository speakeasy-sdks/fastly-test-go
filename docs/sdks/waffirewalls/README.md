# WafFirewalls

## Overview

A firewall represents a configuration of a Web Application Firewall (WAF) on a service. A firewall has many firewall versions through which you can manage rules.

<https://developer.fastly.com/reference/api/waf/firewall>
### Available Operations

* [~~CreateWafFirewall~~](#createwaffirewall) - Create a firewall :warning: **Deprecated**
* [~~DeleteWafFirewall~~](#deletewaffirewall) - Delete a firewall :warning: **Deprecated**
* [~~GetWafFirewall~~](#getwaffirewall) - Get a firewall :warning: **Deprecated**
* [~~ListWafFirewalls~~](#listwaffirewalls) - List firewalls :warning: **Deprecated**
* [~~UpdateWafFirewall~~](#updatewaffirewall) - Update a firewall :warning: **Deprecated**

## ~~CreateWafFirewall~~

Create a firewall object for a particular service and service version using a defined `prefetch_condition` and `response`. If the `prefetch_condition` or the `response` is missing from the request body, Fastly will generate a default object on your service.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    operationSecurity := operations.CreateWafFirewallSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafFirewalls.CreateWafFirewall(ctx, shared.WafFirewallInput{
        Data: &shared.WafFirewallDataInput{
            Attributes: &shared.WafFirewallDataAttributesInput{
                Disabled: sdk.Bool(false),
                PrefetchCondition: sdk.String("repellendus"),
                Response: sdk.String("delectus"),
            },
            Type: shared.TypeWafFirewallWafFirewall.ToPointer(),
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.WafFirewallInput](../../models/shared/waffirewallinput.md)                           | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.CreateWafFirewallSecurity](../../models/operations/createwaffirewallsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.CreateWafFirewallResponse](../../models/operations/createwaffirewallresponse.md), error**


## ~~DeleteWafFirewall~~

Delete the firewall object for a particular service and service version.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    operationSecurity := operations.DeleteWafFirewallSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafFirewalls.DeleteWafFirewall(ctx, operations.DeleteWafFirewallRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        WafFirewallInput: &shared.WafFirewallInput{
            Data: &shared.WafFirewallDataInput{
                Attributes: &shared.WafFirewallDataAttributesInput{
                    Disabled: sdk.Bool(false),
                    PrefetchCondition: sdk.String("voluptates"),
                    Response: sdk.String("perferendis"),
                },
                Type: shared.TypeWafFirewallWafFirewall.ToPointer(),
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteWafFirewallRequest](../../models/operations/deletewaffirewallrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.DeleteWafFirewallSecurity](../../models/operations/deletewaffirewallsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.DeleteWafFirewallResponse](../../models/operations/deletewaffirewallresponse.md), error**


## ~~GetWafFirewall~~

Get a specific firewall object.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    operationSecurity := operations.GetWafFirewallSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafFirewalls.GetWafFirewall(ctx, operations.GetWafFirewallRequest{
        FilterServiceVersionNumber: sdk.String("est"),
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        Include: shared.FirewallIncludeWafFirewallVersions.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetWafFirewallRequest](../../models/operations/getwaffirewallrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetWafFirewallSecurity](../../models/operations/getwaffirewallsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetWafFirewallResponse](../../models/operations/getwaffirewallresponse.md), error**


## ~~ListWafFirewalls~~

List all firewall objects.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    operationSecurity := operations.ListWafFirewallsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafFirewalls.ListWafFirewalls(ctx, operations.ListWafFirewallsRequest{
        FilterServiceID: sdk.String("quidem"),
        FilterServiceVersionNumber: sdk.String("reprehenderit"),
        Include: shared.FirewallIncludeWafFirewallVersions.ToPointer(),
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListWafFirewallsRequest](../../models/operations/listwaffirewallsrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ListWafFirewallsSecurity](../../models/operations/listwaffirewallssecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ListWafFirewallsResponse](../../models/operations/listwaffirewallsresponse.md), error**


## ~~UpdateWafFirewall~~

Update a firewall object for a particular service and service version. Specifying a `service_version_number` is required.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    operationSecurity := operations.UpdateWafFirewallSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafFirewalls.UpdateWafFirewall(ctx, operations.UpdateWafFirewallRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        WafFirewallInput: &shared.WafFirewallInput{
            Data: &shared.WafFirewallDataInput{
                Attributes: &shared.WafFirewallDataAttributesInput{
                    Disabled: sdk.Bool(false),
                    PrefetchCondition: sdk.String("facere"),
                    Response: sdk.String("fuga"),
                },
                Type: shared.TypeWafFirewallWafFirewall.ToPointer(),
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateWafFirewallRequest](../../models/operations/updatewaffirewallrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UpdateWafFirewallSecurity](../../models/operations/updatewaffirewallsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.UpdateWafFirewallResponse](../../models/operations/updatewaffirewallresponse.md), error**

