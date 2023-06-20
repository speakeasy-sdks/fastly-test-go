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

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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

    ctx := context.Background()
    res, err := s.WafActiveRules.BulkUpdateWafActiveRules(ctx, operations.BulkUpdateWafActiveRulesRequest{
        BulkWafActiveRuleInput: &shared.BulkWafActiveRuleInput{
            Attributes: &shared.BulkWafActiveRuleAttributes{
                ModsecRuleID: sdk.Int64(117315),
                Revision: &shared.WafRuleRevisionOrLatest{},
                Status: shared.BulkWafActiveRuleAttributesStatusBlock.ToPointer(),
            },
            Relationships: &shared.RelationshipWafFirewallVersionInput{
                WafFirewallVersion: &shared.RelationshipWafFirewallVersionWafFirewallVersionInput{
                    Data: []shared.RelationshipMemberWafFirewallVersionInput{
                        shared.RelationshipMemberWafFirewallVersionInput{
                            Type: shared.TypeWafFirewallVersionWafFirewallVersion.ToPointer(),
                        },
                        shared.RelationshipMemberWafFirewallVersionInput{
                            Type: shared.TypeWafFirewallVersionWafFirewallVersion.ToPointer(),
                        },
                    },
                },
            },
            Type: shared.TypeWafActiveRuleWafActiveRule.ToPointer(),
        },
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        VersionID: 1,
    }, operations.BulkUpdateWafActiveRulesSecurity{
        Token: "",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.BulkUpdateWafActiveRulesRequest](../../models/operations/bulkupdatewafactiverulesrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.BulkUpdateWafActiveRulesSecurity](../../models/operations/bulkupdatewafactiverulessecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.BulkUpdateWafActiveRulesResponse](../../models/operations/bulkupdatewafactiverulesresponse.md), error**


## ~~CreateWafActiveRule~~

Create an active rule for a particular firewall version.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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

    ctx := context.Background()
    res, err := s.WafActiveRules.CreateWafActiveRule(ctx, operations.CreateWafActiveRuleRequest{
        BulkWafActiveRulesInput: &shared.BulkWafActiveRulesInput{
            Data: []shared.WafActiveRuleDataInput{
                shared.WafActiveRuleDataInput{
                    Attributes: &shared.WafActiveRuleDataAttributes{
                        ModsecRuleID: sdk.Int64(235263),
                        Revision: &shared.WafRuleRevisionOrLatest{},
                        Status: shared.WafActiveRuleDataAttributesStatusBlock.ToPointer(),
                    },
                    Relationships: &shared.RelationshipWafFirewallVersionInput{
                        WafFirewallVersion: &shared.RelationshipWafFirewallVersionWafFirewallVersionInput{
                            Data: []shared.RelationshipMemberWafFirewallVersionInput{
                                shared.RelationshipMemberWafFirewallVersionInput{
                                    Type: shared.TypeWafFirewallVersionWafFirewallVersion.ToPointer(),
                                },
                                shared.RelationshipMemberWafFirewallVersionInput{
                                    Type: shared.TypeWafFirewallVersionWafFirewallVersion.ToPointer(),
                                },
                                shared.RelationshipMemberWafFirewallVersionInput{
                                    Type: shared.TypeWafFirewallVersionWafFirewallVersion.ToPointer(),
                                },
                            },
                        },
                    },
                    Type: shared.TypeWafActiveRuleWafActiveRule.ToPointer(),
                },
                shared.WafActiveRuleDataInput{
                    Attributes: &shared.WafActiveRuleDataAttributes{
                        ModsecRuleID: sdk.Int64(123844),
                        Revision: &shared.WafRuleRevisionOrLatest{},
                        Status: shared.WafActiveRuleDataAttributesStatusBlock.ToPointer(),
                    },
                    Relationships: &shared.RelationshipWafRuleRevisionInput{
                        WafRuleRevisions: &shared.RelationshipWafRuleRevisionWafRuleRevisionsInput{
                            Data: []shared.RelationshipMemberWafRuleRevisionInput{
                                shared.RelationshipMemberWafRuleRevisionInput{
                                    Type: shared.TypeWafRuleRevisionWafRuleRevision.ToPointer(),
                                },
                                shared.RelationshipMemberWafRuleRevisionInput{
                                    Type: shared.TypeWafRuleRevisionWafRuleRevision.ToPointer(),
                                },
                                shared.RelationshipMemberWafRuleRevisionInput{
                                    Type: shared.TypeWafRuleRevisionWafRuleRevision.ToPointer(),
                                },
                            },
                        },
                    },
                    Type: shared.TypeWafActiveRuleWafActiveRule.ToPointer(),
                },
            },
        },
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        VersionID: 1,
    }, operations.CreateWafActiveRuleSecurity{
        Token: "",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateWafActiveRuleRequest](../../models/operations/createwafactiverulerequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.CreateWafActiveRuleSecurity](../../models/operations/createwafactiverulesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.CreateWafActiveRuleResponse](../../models/operations/createwafactiveruleresponse.md), error**


## ~~CreateWafActiveRulesTag~~

Create active rules by tag. This endpoint will create active rules using the latest revision available for each rule.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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

    ctx := context.Background()
    res, err := s.WafActiveRules.CreateWafActiveRulesTag(ctx, operations.CreateWafActiveRulesTagRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        VersionID: 1,
        WafActiveRuleInput: &shared.WafActiveRuleInput{
            Data: &shared.WafActiveRuleDataInput{
                Attributes: &shared.WafActiveRuleDataAttributes{
                    ModsecRuleID: sdk.Int64(399667),
                    Revision: &shared.WafRuleRevisionOrLatest{},
                    Status: shared.WafActiveRuleDataAttributesStatusBlock.ToPointer(),
                },
                Relationships: &shared.RelationshipWafFirewallVersionInput{
                    WafFirewallVersion: &shared.RelationshipWafFirewallVersionWafFirewallVersionInput{
                        Data: []shared.RelationshipMemberWafFirewallVersionInput{
                            shared.RelationshipMemberWafFirewallVersionInput{
                                Type: shared.TypeWafFirewallVersionWafFirewallVersion.ToPointer(),
                            },
                            shared.RelationshipMemberWafFirewallVersionInput{
                                Type: shared.TypeWafFirewallVersionWafFirewallVersion.ToPointer(),
                            },
                        },
                    },
                },
                Type: shared.TypeWafActiveRuleWafActiveRule.ToPointer(),
            },
        },
        WafTagName: "test-waf-tag",
    }, operations.CreateWafActiveRulesTagSecurity{
        Token: "",
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
| `request`                                                                                                | [operations.CreateWafActiveRulesTagRequest](../../models/operations/createwafactiverulestagrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.CreateWafActiveRulesTagSecurity](../../models/operations/createwafactiverulestagsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.CreateWafActiveRulesTagResponse](../../models/operations/createwafactiverulestagresponse.md), error**


## ~~DeleteWafActiveRule~~

Delete an active rule for a particular firewall version.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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

    ctx := context.Background()
    res, err := s.WafActiveRules.DeleteWafActiveRule(ctx, operations.DeleteWafActiveRuleRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        VersionID: 1,
        WafRuleID: "3krg2uUGZzb2W9Euo4moOR",
    }, operations.DeleteWafActiveRuleSecurity{
        Token: "",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteWafActiveRuleRequest](../../models/operations/deletewafactiverulerequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.DeleteWafActiveRuleSecurity](../../models/operations/deletewafactiverulesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.DeleteWafActiveRuleResponse](../../models/operations/deletewafactiveruleresponse.md), error**


## ~~GetWafActiveRule~~

Get a specific active rule object. Includes details of the rule revision associated with the active rule object by default.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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

    ctx := context.Background()
    res, err := s.WafActiveRules.GetWafActiveRule(ctx, operations.GetWafActiveRuleRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        Include: sdk.String("waf_rule_revision,waf_firewall_version"),
        VersionID: 1,
        WafRuleID: "3krg2uUGZzb2W9Euo4moOR",
    }, operations.GetWafActiveRuleSecurity{
        Token: "",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetWafActiveRuleRequest](../../models/operations/getwafactiverulerequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.GetWafActiveRuleSecurity](../../models/operations/getwafactiverulesecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.GetWafActiveRuleResponse](../../models/operations/getwafactiveruleresponse.md), error**


## ~~ListWafActiveRules~~

List all active rules for a particular firewall version.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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

    ctx := context.Background()
    res, err := s.WafActiveRules.ListWafActiveRules(ctx, operations.ListWafActiveRulesRequest{
        FilterOutdated: sdk.String("perferendis"),
        FilterStatus: sdk.String("eum"),
        FilterWafRuleRevisionMessage: sdk.String("voluptas"),
        FilterWafRuleRevisionModsecRuleID: sdk.String("iste"),
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        Include: sdk.String("waf_rule_revision,waf_firewall_version"),
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
        VersionID: 1,
    }, operations.ListWafActiveRulesSecurity{
        Token: "",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListWafActiveRulesRequest](../../models/operations/listwafactiverulesrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ListWafActiveRulesSecurity](../../models/operations/listwafactiverulessecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ListWafActiveRulesResponse](../../models/operations/listwafactiverulesresponse.md), error**


## ~~UpdateWafActiveRule~~

Update an active rule's status for a particular firewall version.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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

    ctx := context.Background()
    res, err := s.WafActiveRules.UpdateWafActiveRule(ctx, operations.UpdateWafActiveRuleRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        VersionID: 1,
        WafActiveRuleInput: &shared.WafActiveRuleInput{
            Data: &shared.WafActiveRuleDataInput{
                Attributes: &shared.WafActiveRuleDataAttributes{
                    ModsecRuleID: sdk.Int64(661607),
                    Revision: &shared.WafRuleRevisionOrLatest{},
                    Status: shared.WafActiveRuleDataAttributesStatusLog.ToPointer(),
                },
                Relationships: &shared.RelationshipWafRuleRevisionInput{
                    WafRuleRevisions: &shared.RelationshipWafRuleRevisionWafRuleRevisionsInput{
                        Data: []shared.RelationshipMemberWafRuleRevisionInput{
                            shared.RelationshipMemberWafRuleRevisionInput{
                                Type: shared.TypeWafRuleRevisionWafRuleRevision.ToPointer(),
                            },
                            shared.RelationshipMemberWafRuleRevisionInput{
                                Type: shared.TypeWafRuleRevisionWafRuleRevision.ToPointer(),
                            },
                            shared.RelationshipMemberWafRuleRevisionInput{
                                Type: shared.TypeWafRuleRevisionWafRuleRevision.ToPointer(),
                            },
                            shared.RelationshipMemberWafRuleRevisionInput{
                                Type: shared.TypeWafRuleRevisionWafRuleRevision.ToPointer(),
                            },
                        },
                    },
                },
                Type: shared.TypeWafActiveRuleWafActiveRule.ToPointer(),
            },
        },
        WafRuleID: "3krg2uUGZzb2W9Euo4moOR",
    }, operations.UpdateWafActiveRuleSecurity{
        Token: "",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateWafActiveRuleRequest](../../models/operations/updatewafactiverulerequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.UpdateWafActiveRuleSecurity](../../models/operations/updatewafactiverulesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.UpdateWafActiveRuleResponse](../../models/operations/updatewafactiveruleresponse.md), error**
