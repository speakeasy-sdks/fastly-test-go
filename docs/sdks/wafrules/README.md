# WafRules
(*WafRules*)

## Overview

Rules are universally available for every firewall. Rules can have one or multiple [rule revisions](/reference/api/waf/rules/revisions/). You can add rules to your firewall by creating [active rules](/reference/api/waf/rules/active/).

<https://developer.fastly.com/reference/api/waf/rules>
### Available Operations

* [~~GetWafRule~~](#getwafrule) - Get a rule :warning: **Deprecated**
* [~~ListWafRules~~](#listwafrules) - List available WAF rules :warning: **Deprecated**

## ~~GetWafRule~~

Get a specific rule. The `id` provided can be the ModSecurity Rule ID or the Fastly generated rule ID.

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
    res, err := s.WafRules.GetWafRule(ctx, operations.GetWafRuleRequest{
        Include: fastlytestgo.String("waf_tags,waf_rule_revisions"),
        WafRuleID: "3krg2uUGZzb2W9Euo4moOR",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetWafRuleRequest](../../models/operations/getwafrulerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetWafRuleResponse](../../models/operations/getwafruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~ListWafRules~~

List all available WAF rules.

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
    res, err := s.WafRules.ListWafRules(ctx, operations.ListWafRulesRequest{
        Include: fastlytestgo.String("waf_tags,waf_rule_revisions"),
        PageNumber: fastlytestgo.Int64(1),
        PageSize: fastlytestgo.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafRulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListWafRulesRequest](../../models/operations/listwafrulesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListWafRulesResponse](../../models/operations/listwafrulesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
