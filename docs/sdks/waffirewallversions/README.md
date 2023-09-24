# WafFirewallVersions

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


## ~~CreateWafFirewallVersion~~

Create a new, draft firewall version.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    res, err := s.WafFirewallVersions.CreateWafFirewallVersion(ctx, operations.CreateWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        WafFirewallVersionInput: &shared.WafFirewallVersionInput{
            Data: &shared.WafFirewallVersionDataInput{
                Attributes: &shared.WafFirewallVersionDataAttributesInput{
                    AllowedHTTPVersions: fastly.String("harum"),
                    AllowedMethods: fastly.String("dicta"),
                    AllowedRequestContentType: fastly.String("architecto"),
                    AllowedRequestContentTypeCharset: fastly.String("occaecati"),
                    ArgLength: fastly.Int64(289776),
                    ArgNameLength: fastly.Int64(695270),
                    CombinedFileSizes: fastly.Int64(539074),
                    Comment: fastly.String("laborum"),
                    CriticalAnomalyScore: fastly.Int64(724148),
                    CrsValidateUTF8Encoding: fastly.Bool(false),
                    ErrorAnomalyScore: fastly.Int64(948861),
                    HighRiskCountryCodes: fastly.String("laboriosam"),
                    HTTPViolationScoreThreshold: fastly.Int64(2703),
                    InboundAnomalyScoreThreshold: fastly.Int64(227084),
                    LfiScoreThreshold: fastly.Int64(647197),
                    Locked: fastly.Bool(false),
                    MaxFileSize: fastly.Int64(454860),
                    MaxNumArgs: fastly.Int64(600392),
                    NoticeAnomalyScore: fastly.Int64(972083),
                    ParanoiaLevel: fastly.Int64(588740),
                    PhpInjectionScoreThreshold: fastly.Int64(833819),
                    RceScoreThreshold: fastly.Int64(962771),
                    RestrictedExtensions: fastly.String("voluptates"),
                    RestrictedHeaders: fastly.String("perferendis"),
                    RfiScoreThreshold: fastly.Int64(667285),
                    SessionFixationScoreThreshold: fastly.Int64(696483),
                    SQLInjectionScoreThreshold: fastly.Int64(440666),
                    TotalArgLength: fastly.Int64(813679),
                    WarningAnomalyScore: fastly.Int64(685092),
                    XSSScoreThreshold: fastly.Int64(509807),
                },
                Type: fastly.String("mollitia"),
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


## ~~DeployActivateWafFirewallVersion~~

Deploy or activate a specific firewall version. If a firewall has been disabled, deploying a firewall version will automatically enable the firewall again.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    res, err := s.WafFirewallVersions.DeployActivateWafFirewallVersion(ctx, operations.DeployActivateWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeployActivateWafFirewallVersion202ApplicationVndAPIPlusJSONObject != nil {
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


## ~~GetWafFirewallVersion~~

Get details about a specific firewall version.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    res, err := s.WafFirewallVersions.GetWafFirewallVersion(ctx, operations.GetWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
        Include: fastly.String("waf_firewall,waf_active_rules"),
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


## ~~ListWafFirewallVersions~~

Get a list of firewall versions associated with a specific firewall.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    res, err := s.WafFirewallVersions.ListWafFirewallVersions(ctx, operations.ListWafFirewallVersionsRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        Include: fastly.String("waf_firewall"),
        PageNumber: fastly.Int64(1),
        PageSize: fastly.Int64(20),
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


## ~~UpdateWafFirewallVersion~~

Update a specific firewall version.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    res, err := s.WafFirewallVersions.UpdateWafFirewallVersion(ctx, operations.UpdateWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
        WafFirewallVersionInput: &shared.WafFirewallVersionInput{
            Data: &shared.WafFirewallVersionDataInput{
                Attributes: &shared.WafFirewallVersionDataAttributesInput{
                    AllowedHTTPVersions: fastly.String("veniam"),
                    AllowedMethods: fastly.String("voluptatem"),
                    AllowedRequestContentType: fastly.String("quisquam"),
                    AllowedRequestContentTypeCharset: fastly.String("repudiandae"),
                    ArgLength: fastly.Int64(97243),
                    ArgNameLength: fastly.Int64(542457),
                    CombinedFileSizes: fastly.Int64(442036),
                    Comment: fastly.String("asperiores"),
                    CriticalAnomalyScore: fastly.Int64(519952),
                    CrsValidateUTF8Encoding: fastly.Bool(false),
                    ErrorAnomalyScore: fastly.Int64(383103),
                    HighRiskCountryCodes: fastly.String("quidem"),
                    HTTPViolationScoreThreshold: fastly.Int64(806670),
                    InboundAnomalyScoreThreshold: fastly.Int64(90885),
                    LfiScoreThreshold: fastly.Int64(461007),
                    Locked: fastly.Bool(false),
                    MaxFileSize: fastly.Int64(227759),
                    MaxNumArgs: fastly.Int64(826825),
                    NoticeAnomalyScore: fastly.Int64(410301),
                    ParanoiaLevel: fastly.Int64(539118),
                    PhpInjectionScoreThreshold: fastly.Int64(623295),
                    RceScoreThreshold: fastly.Int64(887265),
                    RestrictedExtensions: fastly.String("officiis"),
                    RestrictedHeaders: fastly.String("accusamus"),
                    RfiScoreThreshold: fastly.Int64(618826),
                    SessionFixationScoreThreshold: fastly.Int64(328303),
                    SQLInjectionScoreThreshold: fastly.Int64(133461),
                    TotalArgLength: fastly.Int64(404425),
                    WarningAnomalyScore: fastly.Int64(980581),
                    XSSScoreThreshold: fastly.Int64(544647),
                },
                Type: fastly.String("at"),
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

