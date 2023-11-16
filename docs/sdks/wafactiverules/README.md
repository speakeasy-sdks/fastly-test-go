# WafActiveRules
(*WafActiveRules*)

## Overview

An active rule represents a [rule revision](/reference/api/waf/rules/revisions/) added to a particular [firewall version](/reference/api/waf/firewall-version/).

<https://developer.fastly.com/reference/api/waf/rules/active>
### Available Operations

* [~~BulkUpdateWafActiveRules~~](#bulkupdatewafactiverules) - Update multiple active rules :warning: **Deprecated**
* [~~CreateWafActiveRule~~](#createwafactiverule) - Add a rule to a WAF as an active rule :warning: **Deprecated**
* [~~CreateWafActiveRulesTag~~](#createwafactiverulestag) - Create active rules by tag :warning: **Deprecated**
* [~~DeleteWafActiveRule~~](#deletewafactiverule) - Delete an active rule :warning: **Deprecated**
* [~~GetWafActiveRule~~](#getwafactiverule) - Get an active WAF rule object :warning: **Deprecated**
* [~~ListWafActiveRules~~](#listwafactiverules) - List active rules on a WAF :warning: **Deprecated**
* [~~UpdateWafActiveRule~~](#updatewafactiverule) - Update an active rule :warning: **Deprecated**

## ~~BulkUpdateWafActiveRules~~

Bulk update all active rules on a [firewall version](https://developer.fastly.com/reference/api/waf/firewall-version/). This endpoint will not add new active rules, only update existing active rules.

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
    res, err := s.WafActiveRules.BulkUpdateWafActiveRules(ctx, operations.BulkUpdateWafActiveRulesRequest{
        BulkWafActiveRule: &components.BulkWafActiveRule{
            Attributes: &components.Attributes{
                Revision: components.CreateWafRuleRevisionOrLatestInteger(
                2,
                ),
            },
            Relationships: components.CreateRelationshipsForWafActiveRuleInputRelationshipWafFirewallVersionInput(
                    components.RelationshipWafFirewallVersionInput{
                        WafFirewallVersion: &components.RelationshipWafFirewallVersionWafFirewallVersionInput{
                            Data: []components.RelationshipMemberWafFirewallVersionInput{
                                components.RelationshipMemberWafFirewallVersionInput{},
                            },
                        },
                    },
            ),
        },
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        VersionID: 1,
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.BulkUpdateWafActiveRulesRequest](../../models/operations/bulkupdatewafactiverulesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.BulkUpdateWafActiveRulesResponse](../../models/operations/bulkupdatewafactiverulesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~CreateWafActiveRule~~

Create an active rule for a particular firewall version.

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
    res, err := s.WafActiveRules.CreateWafActiveRule(ctx, operations.CreateWafActiveRuleRequest{
        BulkWafActiveRules: &components.BulkWafActiveRules{
            Data: []components.WafActiveRuleData1{
                components.WafActiveRuleData1{
                    Attributes: &components.WafActiveRuleDataAttributes{
                        Revision: components.CreateWafRuleRevisionOrLatestStr(
                        "latest",
                        ),
                    },
                    Relationships: components.CreateRelationshipsForWafActiveRuleInputRelationshipWafRuleRevisionInput(
                            components.RelationshipWafRuleRevisionInput{
                                WafRuleRevisions: &components.RelationshipWafRuleRevisionWafRuleRevisions{
                                    Data: []components.RelationshipMemberWafRuleRevisionInput{
                                        components.RelationshipMemberWafRuleRevisionInput{},
                                    },
                                },
                            },
                    ),
                },
            },
        },
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafActiveRuleCreationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateWafActiveRuleRequest](../../models/operations/createwafactiverulerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.CreateWafActiveRuleResponse](../../models/operations/createwafactiveruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~CreateWafActiveRulesTag~~

Create active rules by tag. This endpoint will create active rules using the latest revision available for each rule.

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
    res, err := s.WafActiveRules.CreateWafActiveRulesTag(ctx, operations.CreateWafActiveRulesTagRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        VersionID: 1,
        WafActiveRule: &components.WafActiveRule1{
            Data: &components.WafActiveRuleData1{
                Attributes: &components.WafActiveRuleDataAttributes{
                    Revision: components.CreateWafRuleRevisionOrLatestStr(
                    "latest",
                    ),
                },
                Relationships: components.CreateRelationshipsForWafActiveRuleInputRelationshipWafFirewallVersionInput(
                        components.RelationshipWafFirewallVersionInput{
                            WafFirewallVersion: &components.RelationshipWafFirewallVersionWafFirewallVersionInput{
                                Data: []components.RelationshipMemberWafFirewallVersionInput{
                                    components.RelationshipMemberWafFirewallVersionInput{},
                                },
                            },
                        },
                ),
            },
        },
        WafTagName: "test-waf-tag",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateWafActiveRulesTagRequest](../../models/operations/createwafactiverulestagrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.CreateWafActiveRulesTagResponse](../../models/operations/createwafactiverulestagresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~DeleteWafActiveRule~~

Delete an active rule for a particular firewall version.

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
    res, err := s.WafActiveRules.DeleteWafActiveRule(ctx, operations.DeleteWafActiveRuleRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        VersionID: 1,
        WafRuleID: "3krg2uUGZzb2W9Euo4moOR",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteWafActiveRuleRequest](../../models/operations/deletewafactiverulerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteWafActiveRuleResponse](../../models/operations/deletewafactiveruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~GetWafActiveRule~~

Get a specific active rule object. Includes details of the rule revision associated with the active rule object by default.

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
    res, err := s.WafActiveRules.GetWafActiveRule(ctx, operations.GetWafActiveRuleRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        Include: fastlytestgo.String("waf_rule_revision,waf_firewall_version"),
        VersionID: 1,
        WafRuleID: "3krg2uUGZzb2W9Euo4moOR",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafActiveRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetWafActiveRuleRequest](../../models/operations/getwafactiverulerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetWafActiveRuleResponse](../../models/operations/getwafactiveruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~ListWafActiveRules~~

List all active rules for a particular firewall version.

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
    res, err := s.WafActiveRules.ListWafActiveRules(ctx, operations.ListWafActiveRulesRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        Include: fastlytestgo.String("waf_rule_revision,waf_firewall_version"),
        PageNumber: fastlytestgo.Int64(1),
        PageSize: fastlytestgo.Int64(20),
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafActiveRulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListWafActiveRulesRequest](../../models/operations/listwafactiverulesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListWafActiveRulesResponse](../../models/operations/listwafactiverulesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~UpdateWafActiveRule~~

Update an active rule's status for a particular firewall version.

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
    res, err := s.WafActiveRules.UpdateWafActiveRule(ctx, operations.UpdateWafActiveRuleRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        VersionID: 1,
        WafActiveRule: &components.WafActiveRule1{
            Data: &components.WafActiveRuleData1{
                Attributes: &components.WafActiveRuleDataAttributes{
                    Revision: components.CreateWafRuleRevisionOrLatestStr(
                    "latest",
                    ),
                },
                Relationships: components.CreateRelationshipsForWafActiveRuleInputRelationshipWafRuleRevisionInput(
                        components.RelationshipWafRuleRevisionInput{
                            WafRuleRevisions: &components.RelationshipWafRuleRevisionWafRuleRevisions{
                                Data: []components.RelationshipMemberWafRuleRevisionInput{
                                    components.RelationshipMemberWafRuleRevisionInput{},
                                },
                            },
                        },
                ),
            },
        },
        WafRuleID: "3krg2uUGZzb2W9Euo4moOR",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafActiveRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateWafActiveRuleRequest](../../models/operations/updatewafactiverulerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdateWafActiveRuleResponse](../../models/operations/updatewafactiveruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
