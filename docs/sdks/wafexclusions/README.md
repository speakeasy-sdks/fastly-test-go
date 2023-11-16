# WafExclusions
(*WafExclusions*)

## Overview

WAF rule exclusions provide a flexible way to handle false positives, allowing specific variables, rules, and the entire WAF to be excluded on a per request basis. You can configure up to 300 WAF exclusions and each exclusion of type `rule` can have up to 30 rules associated with it.


<https://developer.fastly.com/reference/api/waf/rules/exclusions>
### Available Operations

* [~~CreateWafRuleExclusion~~](#createwafruleexclusion) - Create a WAF rule exclusion :warning: **Deprecated**
* [~~DeleteWafRuleExclusion~~](#deletewafruleexclusion) - Delete a WAF rule exclusion :warning: **Deprecated**
* [~~GetWafRuleExclusion~~](#getwafruleexclusion) - Get a WAF rule exclusion :warning: **Deprecated**
* [~~ListWafRuleExclusions~~](#listwafruleexclusions) - List WAF rule exclusions :warning: **Deprecated**
* [~~UpdateWafRuleExclusion~~](#updatewafruleexclusion) - Update a WAF rule exclusion :warning: **Deprecated**

## ~~CreateWafRuleExclusion~~

Create a WAF exclusion for a particular firewall version.

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
    res, err := s.WafExclusions.CreateWafRuleExclusion(ctx, operations.CreateWafRuleExclusionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
        WafExclusion: &components.WafExclusion{
            Data: &components.WafExclusionData{
                Attributes: &components.WafExclusionDataAttributes{
                    Number: fastlytestgo.Int64(1),
                },
                Relationships: components.CreateRelationshipsForWafExclusionRelationshipWafRuleRevisionsInput(
                        components.RelationshipWafRuleRevisionsInput{
                            WafRuleRevisions: &components.RelationshipWafRuleRevisionsWafRuleRevisionsInput{
                                Data: []components.RelationshipMemberWafRuleRevisionInput{
                                    components.RelationshipMemberWafRuleRevisionInput{},
                                },
                            },
                        },
                ),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafExclusionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateWafRuleExclusionRequest](../../models/operations/createwafruleexclusionrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.CreateWafRuleExclusionResponse](../../models/operations/createwafruleexclusionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~DeleteWafRuleExclusion~~

Delete a WAF exclusion for a particular firewall version.

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
	"net/http"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.WafExclusions.DeleteWafRuleExclusion(ctx, operations.DeleteWafRuleExclusionRequest{
        ExclusionNumber: 1,
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteWafRuleExclusionRequest](../../models/operations/deletewafruleexclusionrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.DeleteWafRuleExclusionResponse](../../models/operations/deletewafruleexclusionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~GetWafRuleExclusion~~

Get a specific WAF exclusion object.

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
    res, err := s.WafExclusions.GetWafRuleExclusion(ctx, operations.GetWafRuleExclusionRequest{
        ExclusionNumber: 1,
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafExclusionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetWafRuleExclusionRequest](../../models/operations/getwafruleexclusionrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetWafRuleExclusionResponse](../../models/operations/getwafruleexclusionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~ListWafRuleExclusions~~

List all exclusions for a particular firewall version.

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
    res, err := s.WafExclusions.ListWafRuleExclusions(ctx, operations.ListWafRuleExclusionsRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
        Include: fastlytestgo.String("waf_rules"),
        PageNumber: fastlytestgo.Int64(1),
        PageSize: fastlytestgo.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafExclusionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListWafRuleExclusionsRequest](../../models/operations/listwafruleexclusionsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.ListWafRuleExclusionsResponse](../../models/operations/listwafruleexclusionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~UpdateWafRuleExclusion~~

Update a WAF exclusion for a particular firewall version.

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
    res, err := s.WafExclusions.UpdateWafRuleExclusion(ctx, operations.UpdateWafRuleExclusionRequest{
        ExclusionNumber: 1,
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
        WafExclusion: &components.WafExclusion{
            Data: &components.WafExclusionData{
                Attributes: &components.WafExclusionDataAttributes{
                    Number: fastlytestgo.Int64(1),
                },
                Relationships: components.CreateRelationshipsForWafExclusionRelationshipWafRulesInput(
                        components.RelationshipWafRulesInput{
                            WafRules: &components.RelationshipWafRulesWafRules{
                                Data: []components.RelationshipMemberWafRuleInput{
                                    components.RelationshipMemberWafRuleInput{},
                                },
                            },
                        },
                ),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafExclusionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateWafRuleExclusionRequest](../../models/operations/updatewafruleexclusionrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.UpdateWafRuleExclusionResponse](../../models/operations/updatewafruleexclusionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
