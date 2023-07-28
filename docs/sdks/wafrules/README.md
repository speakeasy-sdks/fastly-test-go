# WafRules

## Overview

Rules are universally available for every firewall. Rules can have one or multiple [rule revisions](/reference/api/waf/rules/revisions/). You can add rules to your firewall by creating [active rules](/reference/api/waf/rules/active/).

<https://developer.fastly.com/reference/api/waf/rules>
### Available Operations

* [~~GetWafRule~~](#getwafrule) - Get a rule :warning: **Deprecated**
* [~~ListWafRules~~](#listwafrules) - List available WAF rules :warning: **Deprecated**

## ~~GetWafRule~~

Get a specific rule. The `id` provided can be the ModSecurity Rule ID or the Fastly generated rule ID.

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
    operationSecurity := operations.GetWafRuleSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafRules.GetWafRule(ctx, operations.GetWafRuleRequest{
        Include: sdk.String("waf_tags,waf_rule_revisions"),
        WafRuleID: "3krg2uUGZzb2W9Euo4moOR",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetWafRuleRequest](../../models/operations/getwafrulerequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.GetWafRuleSecurity](../../models/operations/getwafrulesecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.GetWafRuleResponse](../../models/operations/getwafruleresponse.md), error**


## ~~ListWafRules~~

List all available WAF rules.

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
    operationSecurity := operations.ListWafRulesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.WafRules.ListWafRules(ctx, operations.ListWafRulesRequest{
        FilterModsecRuleID: sdk.String("porro"),
        FilterWafFirewallIDNotMatch: sdk.String("autem"),
        FilterWafRuleRevisionsSource: sdk.String("nobis"),
        FilterWafTagsName: sdk.String("laboriosam"),
        Include: sdk.String("waf_tags,waf_rule_revisions"),
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.WafRulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListWafRulesRequest](../../models/operations/listwafrulesrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.ListWafRulesSecurity](../../models/operations/listwafrulessecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.ListWafRulesResponse](../../models/operations/listwafrulesresponse.md), error**

