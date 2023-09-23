# WafActiveRules

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
    res, err := s.WafActiveRules.BulkUpdateWafActiveRules(ctx, operations.BulkUpdateWafActiveRulesRequest{
        BulkWafActiveRule: &shared.BulkWafActiveRule{
            Attributes: &shared.BulkWafActiveRuleAttributes{
                ModsecRuleID: fastly.Int64(710529),
                Revision: &shared.WafRuleRevisionOrLatest{},
                Status: shared.BulkWafActiveRuleAttributesStatusScore.ToPointer(),
            },
            Relationships: &shared.RelationshipsForWafActiveRuleInput{},
            Type: fastly.String("neque"),
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


## ~~CreateWafActiveRule~~

Create an active rule for a particular firewall version.

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
    res, err := s.WafActiveRules.CreateWafActiveRule(ctx, operations.CreateWafActiveRuleRequest{
        BulkWafActiveRules: &shared.BulkWafActiveRules{
            Data: []shared.WafActiveRuleData1{
                shared.WafActiveRuleData1{
                    Attributes: &shared.WafActiveRuleDataAttributes{
                        ModsecRuleID: fastly.Int64(677115),
                        Revision: &shared.WafRuleRevisionOrLatest{},
                        Status: shared.WafActiveRuleDataAttributesStatusBlock.ToPointer(),
                    },
                    Relationships: &shared.RelationshipsForWafActiveRuleInput{},
                    Type: fastly.String("officia"),
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


## ~~CreateWafActiveRulesTag~~

Create active rules by tag. This endpoint will create active rules using the latest revision available for each rule.

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
    res, err := s.WafActiveRules.CreateWafActiveRulesTag(ctx, operations.CreateWafActiveRulesTagRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        VersionID: 1,
        WafActiveRule1: &shared.WafActiveRule1{
            Data: &shared.WafActiveRuleData1{
                Attributes: &shared.WafActiveRuleDataAttributes{
                    ModsecRuleID: fastly.Int64(676243),
                    Revision: &shared.WafRuleRevisionOrLatest{},
                    Status: shared.WafActiveRuleDataAttributesStatusBlock.ToPointer(),
                },
                Relationships: &shared.RelationshipsForWafActiveRuleInput{},
                Type: fastly.String("accusamus"),
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


## ~~DeleteWafActiveRule~~

Delete an active rule for a particular firewall version.

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


## ~~GetWafActiveRule~~

Get a specific active rule object. Includes details of the rule revision associated with the active rule object by default.

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
    res, err := s.WafActiveRules.GetWafActiveRule(ctx, operations.GetWafActiveRuleRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        Include: fastly.String("waf_rule_revision,waf_firewall_version"),
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


## ~~ListWafActiveRules~~

List all active rules for a particular firewall version.

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
    res, err := s.WafActiveRules.ListWafActiveRules(ctx, operations.ListWafActiveRulesRequest{
        FilterOutdated: fastly.String("tempora"),
        FilterStatus: fastly.String("atque"),
        FilterWafRuleRevisionMessage: fastly.String("fugit"),
        FilterWafRuleRevisionModsecRuleID: fastly.String("ut"),
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        Include: fastly.String("waf_rule_revision,waf_firewall_version"),
        PageNumber: fastly.Int64(1),
        PageSize: fastly.Int64(20),
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


## ~~UpdateWafActiveRule~~

Update an active rule's status for a particular firewall version.

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
    res, err := s.WafActiveRules.UpdateWafActiveRule(ctx, operations.UpdateWafActiveRuleRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        VersionID: 1,
        WafActiveRule1: &shared.WafActiveRule1{
            Data: &shared.WafActiveRuleData1{
                Attributes: &shared.WafActiveRuleDataAttributes{
                    ModsecRuleID: fastly.Int64(856303),
                    Revision: &shared.WafRuleRevisionOrLatest{},
                    Status: shared.WafActiveRuleDataAttributesStatusLog.ToPointer(),
                },
                Relationships: &shared.RelationshipsForWafActiveRuleInput{},
                Type: fastly.String("culpa"),
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

