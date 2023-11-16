# WafFirewallVersions
(*WafFirewallVersions*)

## Overview

Firewall version objects contain all of the rules and settings for your WAF and remain empty until properly configured. To understand the behavior of thresholds and scores, see [Managing rules](https://docs.fastly.com/en/guides/managing-rules-on-the-fastly-waf#understanding-scoring-active-rule-behavior). Newly created firewall versions are initiated without any associated rules. See [Active Rules](/reference/api/waf/rules/active/) for details. Changes to your WAF's rules and settings can be made by [cloning](/reference/api/waf/firewall-version/#clone-waf-firewall-version) an existing firewall version, making the changes, and then [activating](/reference/api/waf/firewall-version/#deploy-activate-waf-firewall-version) the new firewall version.

<https://developer.fastly.com/reference/api/waf/firewall/version>
### Available Operations

* [~~CloneWafFirewallVersion~~](#clonewaffirewallversion) - Clone a firewall version :warning: **Deprecated**
* [~~CreateWafFirewallVersion~~](#createwaffirewallversion) - Create a firewall version :warning: **Deprecated**
* [~~DeployActivateWafFirewallVersion~~](#deployactivatewaffirewallversion) - Deploy or activate a firewall version :warning: **Deprecated**
* [~~GetWafFirewallVersion~~](#getwaffirewallversion) - Get a firewall version :warning: **Deprecated**
* [~~ListWafFirewallVersions~~](#listwaffirewallversions) - List firewall versions :warning: **Deprecated**
* [~~UpdateWafFirewallVersion~~](#updatewaffirewallversion) - Update a firewall version :warning: **Deprecated**

## ~~CloneWafFirewallVersion~~

Clone a specific, existing firewall version into a new, draft firewall version.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.WafFirewallVersions.CloneWafFirewallVersion(ctx, operations.CloneWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CloneWafFirewallVersionRequest](../../models/operations/clonewaffirewallversionrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.CloneWafFirewallVersionResponse](../../models/operations/clonewaffirewallversionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~CreateWafFirewallVersion~~

Create a new, draft firewall version.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.WafFirewallVersions.CreateWafFirewallVersion(ctx, operations.CreateWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        WafFirewallVersion: &components.WafFirewallVersion{
            Data: &components.WafFirewallVersionData{
                Attributes: &components.WafFirewallVersionDataAttributes{
                    Comment: fastlytestgo.String("Ergonomic executive chair upholstered in bonded black leather and PVC padded seat and back for all-day comfort and support"),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateWafFirewallVersionRequest](../../models/operations/createwaffirewallversionrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.CreateWafFirewallVersionResponse](../../models/operations/createwaffirewallversionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~DeployActivateWafFirewallVersion~~

Deploy or activate a specific firewall version. If a firewall has been disabled, deploying a firewall version will automatically enable the firewall again.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.WafFirewallVersions.DeployActivateWafFirewallVersion(ctx, operations.DeployActivateWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeployActivateWafFirewallVersionRequest](../../models/operations/deployactivatewaffirewallversionrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.DeployActivateWafFirewallVersionResponse](../../models/operations/deployactivatewaffirewallversionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~GetWafFirewallVersion~~

Get details about a specific firewall version.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.WafFirewallVersions.GetWafFirewallVersion(ctx, operations.GetWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
        Include: fastlytestgo.String("waf_firewall,waf_active_rules"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetWafFirewallVersionRequest](../../models/operations/getwaffirewallversionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetWafFirewallVersionResponse](../../models/operations/getwaffirewallversionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~ListWafFirewallVersions~~

Get a list of firewall versions associated with a specific firewall.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.WafFirewallVersions.ListWafFirewallVersions(ctx, operations.ListWafFirewallVersionsRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        Include: fastlytestgo.String("waf_firewall"),
        PageNumber: fastlytestgo.Int64(1),
        PageSize: fastlytestgo.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallVersionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListWafFirewallVersionsRequest](../../models/operations/listwaffirewallversionsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.ListWafFirewallVersionsResponse](../../models/operations/listwaffirewallversionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~UpdateWafFirewallVersion~~

Update a specific firewall version.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.WafFirewallVersions.UpdateWafFirewallVersion(ctx, operations.UpdateWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
        WafFirewallVersion: &components.WafFirewallVersion{
            Data: &components.WafFirewallVersionData{
                Attributes: &components.WafFirewallVersionDataAttributes{
                    Comment: fastlytestgo.String("The slim & simple Maple Gaming Keyboard from Dev Byte comes with a sleek body and 7- Color RGB LED Back-lighting for smart functionality"),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateWafFirewallVersionRequest](../../models/operations/updatewaffirewallversionrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.UpdateWafFirewallVersionResponse](../../models/operations/updatewaffirewallversionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
