# WafExclusions

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
    operationSecurity := operations.CreateWafRuleExclusionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafExclusions.CreateWafRuleExclusion(ctx, operations.CreateWafRuleExclusionRequest{
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
        WafExclusionInput: &shared.WafExclusionInput{
            Data: &shared.WafExclusionDataInput{
                Attributes: &shared.WafExclusionDataAttributes{
                    Condition: sdk.String("voluptates"),
                    ExclusionType: shared.WafExclusionDataAttributesExclusionTypeVariable.ToPointer(),
                    Logging: sdk.Bool(false),
                    Name: sdk.String("Wm Hane"),
                    Number: sdk.Int64(1),
                    Variable: shared.WafExclusionDataAttributesVariableReqCookies.ToPointer(),
                },
                Relationships: &shared.RelationshipWafRuleRevisionsInput{
                    WafRuleRevisions: &shared.RelationshipWafRuleRevisionsWafRuleRevisionsInput{
                        Data: []shared.RelationshipMemberWafRuleRevisionInput{
                            shared.RelationshipMemberWafRuleRevisionInput{
                                Type: shared.TypeWafRuleRevisionWafRuleRevision.ToPointer(),
                            },
                            shared.RelationshipMemberWafRuleRevisionInput{
                                Type: shared.TypeWafRuleRevisionWafRuleRevision.ToPointer(),
                            },
                        },
                    },
                },
                Type: shared.TypeWafExclusionWafExclusion.ToPointer(),
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafExclusionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateWafRuleExclusionRequest](../../models/operations/createwafruleexclusionrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.CreateWafRuleExclusionSecurity](../../models/operations/createwafruleexclusionsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.CreateWafRuleExclusionResponse](../../models/operations/createwafruleexclusionresponse.md), error**


## ~~DeleteWafRuleExclusion~~

Delete a WAF exclusion for a particular firewall version.

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
    operationSecurity := operations.DeleteWafRuleExclusionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafExclusions.DeleteWafRuleExclusion(ctx, operations.DeleteWafRuleExclusionRequest{
        ExclusionNumber: 1,
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DeleteWafRuleExclusionRequest](../../models/operations/deletewafruleexclusionrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.DeleteWafRuleExclusionSecurity](../../models/operations/deletewafruleexclusionsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.DeleteWafRuleExclusionResponse](../../models/operations/deletewafruleexclusionresponse.md), error**


## ~~GetWafRuleExclusion~~

Get a specific WAF exclusion object.

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
    operationSecurity := operations.GetWafRuleExclusionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafExclusions.GetWafRuleExclusion(ctx, operations.GetWafRuleExclusionRequest{
        ExclusionNumber: 1,
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafExclusionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetWafRuleExclusionRequest](../../models/operations/getwafruleexclusionrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.GetWafRuleExclusionSecurity](../../models/operations/getwafruleexclusionsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.GetWafRuleExclusionResponse](../../models/operations/getwafruleexclusionresponse.md), error**


## ~~ListWafRuleExclusions~~

List all exclusions for a particular firewall version.

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
    operationSecurity := operations.ListWafRuleExclusionsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafExclusions.ListWafRuleExclusions(ctx, operations.ListWafRuleExclusionsRequest{
        FilterExclusionType: operations.ListWafRuleExclusionsFilterExclusionTypeWaf.ToPointer(),
        FilterName: sdk.String("ex"),
        FilterWafRulesModsecRuleID: sdk.Int64(281153),
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
        Include: sdk.String("waf_rules"),
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafExclusionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListWafRuleExclusionsRequest](../../models/operations/listwafruleexclusionsrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.ListWafRuleExclusionsSecurity](../../models/operations/listwafruleexclusionssecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.ListWafRuleExclusionsResponse](../../models/operations/listwafruleexclusionsresponse.md), error**


## ~~UpdateWafRuleExclusion~~

Update a WAF exclusion for a particular firewall version.

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
    operationSecurity := operations.UpdateWafRuleExclusionSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafExclusions.UpdateWafRuleExclusion(ctx, operations.UpdateWafRuleExclusionRequest{
        ExclusionNumber: 1,
        FirewallID: "fW7g2uUGZzb2W9Euo4Mo0r",
        FirewallVersionNumber: 1,
        WafExclusionInput: &shared.WafExclusionInput{
            Data: &shared.WafExclusionDataInput{
                Attributes: &shared.WafExclusionDataAttributes{
                    Condition: sdk.String("ad"),
                    ExclusionType: shared.WafExclusionDataAttributesExclusionTypeWaf.ToPointer(),
                    Logging: sdk.Bool(false),
                    Name: sdk.String("Leona Rippin IV"),
                    Number: sdk.Int64(1),
                    Variable: shared.WafExclusionDataAttributesVariableReqPostFilename.ToPointer(),
                },
                Relationships: &shared.RelationshipWafRulesInput{
                    WafRules: &shared.RelationshipWafRulesWafRulesInput{
                        Data: []shared.RelationshipMemberWafRuleInput{
                            shared.RelationshipMemberWafRuleInput{
                                Type: shared.TypeWafRuleWafRule.ToPointer(),
                            },
                            shared.RelationshipMemberWafRuleInput{
                                Type: shared.TypeWafRuleWafRule.ToPointer(),
                            },
                            shared.RelationshipMemberWafRuleInput{
                                Type: shared.TypeWafRuleWafRule.ToPointer(),
                            },
                        },
                    },
                },
                Type: shared.TypeWafExclusionWafExclusion.ToPointer(),
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafExclusionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateWafRuleExclusionRequest](../../models/operations/updatewafruleexclusionrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.UpdateWafRuleExclusionSecurity](../../models/operations/updatewafruleexclusionsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.UpdateWafRuleExclusionResponse](../../models/operations/updatewafruleexclusionresponse.md), error**

