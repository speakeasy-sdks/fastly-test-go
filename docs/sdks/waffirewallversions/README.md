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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.CloneWafFirewallVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafFirewallVersions.CloneWafFirewallVersion(ctx, operations.CloneWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
    }, operationSecurity)
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
| `request`                                                                                                | [operations.CloneWafFirewallVersionRequest](../../models/operations/clonewaffirewallversionrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.CloneWafFirewallVersionSecurity](../../models/operations/clonewaffirewallversionsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.CreateWafFirewallVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafFirewallVersions.CreateWafFirewallVersion(ctx, operations.CreateWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        WafFirewallVersionInput: &shared.WafFirewallVersionInput{
            Data: &shared.WafFirewallVersionDataInput{
                Attributes: &shared.WafFirewallVersionDataAttributesInput{
                    AllowedHTTPVersions: sdk.String("debitis"),
                    AllowedMethods: sdk.String("neque"),
                    AllowedRequestContentType: sdk.String("dolorum"),
                    AllowedRequestContentTypeCharset: sdk.String("nostrum"),
                    ArgLength: sdk.Int64(639028),
                    ArgNameLength: sdk.Int64(676243),
                    CombinedFileSizes: sdk.Int64(548361),
                    Comment: sdk.String("accusamus"),
                    CriticalAnomalyScore: sdk.Int64(272683),
                    CrsValidateUTF8Encoding: sdk.Bool(false),
                    ErrorAnomalyScore: sdk.Int64(543678),
                    HighRiskCountryCodes: sdk.String("fugit"),
                    HTTPViolationScoreThreshold: sdk.Int64(282699),
                    InboundAnomalyScoreThreshold: sdk.Int64(856303),
                    LfiScoreThreshold: sdk.Int64(30235),
                    Locked: sdk.Bool(false),
                    MaxFileSize: sdk.Int64(635057),
                    MaxNumArgs: sdk.Int64(710337),
                    NoticeAnomalyScore: sdk.Int64(299643),
                    ParanoiaLevel: sdk.Int64(7884),
                    PhpInjectionScoreThreshold: sdk.Int64(460220),
                    RceScoreThreshold: sdk.Int64(372679),
                    RestrictedExtensions: sdk.String("sit"),
                    RestrictedHeaders: sdk.String("voluptatum"),
                    RfiScoreThreshold: sdk.Int64(558065),
                    SessionFixationScoreThreshold: sdk.Int64(922112),
                    SQLInjectionScoreThreshold: sdk.Int64(361151),
                    TotalArgLength: sdk.Int64(89494),
                    WarningAnomalyScore: sdk.Int64(502710),
                    XSSScoreThreshold: sdk.Int64(405942),
                },
                Type: shared.TypeWafFirewallVersionWafFirewallVersion.ToPointer(),
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateWafFirewallVersionRequest](../../models/operations/createwaffirewallversionrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.CreateWafFirewallVersionSecurity](../../models/operations/createwaffirewallversionsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.DeployActivateWafFirewallVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafFirewallVersions.DeployActivateWafFirewallVersion(ctx, operations.DeployActivateWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeployActivateWafFirewallVersion202ApplicationVndAPIPlusJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DeployActivateWafFirewallVersionRequest](../../models/operations/deployactivatewaffirewallversionrequest.md)   | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `security`                                                                                                                 | [operations.DeployActivateWafFirewallVersionSecurity](../../models/operations/deployactivatewaffirewallversionsecurity.md) | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.GetWafFirewallVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafFirewallVersions.GetWafFirewallVersion(ctx, operations.GetWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
        Include: sdk.String("waf_firewall,waf_active_rules"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetWafFirewallVersionRequest](../../models/operations/getwaffirewallversionrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.GetWafFirewallVersionSecurity](../../models/operations/getwaffirewallversionsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListWafFirewallVersionsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafFirewallVersions.ListWafFirewallVersions(ctx, operations.ListWafFirewallVersionsRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        Include: sdk.String("waf_firewall"),
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallVersionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListWafFirewallVersionsRequest](../../models/operations/listwaffirewallversionsrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.ListWafFirewallVersionsSecurity](../../models/operations/listwaffirewallversionssecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.UpdateWafFirewallVersionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafFirewallVersions.UpdateWafFirewallVersion(ctx, operations.UpdateWafFirewallVersionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
        WafFirewallVersionInput: &shared.WafFirewallVersionInput{
            Data: &shared.WafFirewallVersionDataInput{
                Attributes: &shared.WafFirewallVersionDataAttributesInput{
                    AllowedHTTPVersions: sdk.String("sed"),
                    AllowedMethods: sdk.String("sit"),
                    AllowedRequestContentType: sdk.String("vel"),
                    AllowedRequestContentTypeCharset: sdk.String("nostrum"),
                    ArgLength: sdk.Int64(906172),
                    ArgNameLength: sdk.Int64(622231),
                    CombinedFileSizes: sdk.Int64(8511),
                    Comment: sdk.String("incidunt"),
                    CriticalAnomalyScore: sdk.Int64(968865),
                    CrsValidateUTF8Encoding: sdk.Bool(false),
                    ErrorAnomalyScore: sdk.Int64(209750),
                    HighRiskCountryCodes: sdk.String("harum"),
                    HTTPViolationScoreThreshold: sdk.Int64(115703),
                    InboundAnomalyScoreThreshold: sdk.Int64(99416),
                    LfiScoreThreshold: sdk.Int64(577140),
                    Locked: sdk.Bool(false),
                    MaxFileSize: sdk.Int64(289776),
                    MaxNumArgs: sdk.Int64(695270),
                    NoticeAnomalyScore: sdk.Int64(539074),
                    ParanoiaLevel: sdk.Int64(671957),
                    PhpInjectionScoreThreshold: sdk.Int64(724148),
                    RceScoreThreshold: sdk.Int64(948861),
                    RestrictedExtensions: sdk.String("laboriosam"),
                    RestrictedHeaders: sdk.String("alias"),
                    RfiScoreThreshold: sdk.Int64(227084),
                    SessionFixationScoreThreshold: sdk.Int64(647197),
                    SQLInjectionScoreThreshold: sdk.Int64(454860),
                    TotalArgLength: sdk.Int64(600392),
                    WarningAnomalyScore: sdk.Int64(972083),
                    XSSScoreThreshold: sdk.Int64(588740),
                },
                Type: shared.TypeWafFirewallVersionWafFirewallVersion.ToPointer(),
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafFirewallVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateWafFirewallVersionRequest](../../models/operations/updatewaffirewallversionrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.UpdateWafFirewallVersionSecurity](../../models/operations/updatewaffirewallversionsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.UpdateWafFirewallVersionResponse](../../models/operations/updatewaffirewallversionresponse.md), error**

