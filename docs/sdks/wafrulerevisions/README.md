# WafRuleRevisions
(*WafRuleRevisions*)

## Overview

Rule revisions track new variations of [rules](/reference/api/waf/rules/) as they change over time with enhancements, fixes, and improvements. This object allows you to find a specific variation of a rule for use in your application. An [active rule](/reference/api/waf/rules/active/) on a firewall uses a specific rule revision.

<https://developer.fastly.com/reference/api/waf/rules/revisions>
### Available Operations

* [~~GetWafRuleRevision~~](#getwafrulerevision) - Get a revision of a rule :warning: **Deprecated**
* [~~ListWafRuleRevisions~~](#listwafrulerevisions) - List revisions for a rule :warning: **Deprecated**

## ~~GetWafRuleRevision~~

Get a specific rule revision.

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.WafRuleRevisions.GetWafRuleRevision(ctx, operations.GetWafRuleRevisionRequest{
        Include: fastly.String("source,vcl,waf_rule"),
        WafRuleID: "3krg2uUGZzb2W9Euo4moOR",
        WafRuleRevisionNumber: 2,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafRuleRevisionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetWafRuleRevisionRequest](../../models/operations/getwafrulerevisionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetWafRuleRevisionResponse](../../models/operations/getwafrulerevisionresponse.md), error**


## ~~ListWafRuleRevisions~~

List all revisions for a specific rule. The `rule_id` provided can be the ModSecurity Rule ID or the Fastly generated rule ID.

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.WafRuleRevisions.ListWafRuleRevisions(ctx, operations.ListWafRuleRevisionsRequest{
        Include: shared.WafRuleRevisionIncludeWafRule.ToPointer(),
        PageNumber: fastly.Int64(1),
        PageSize: fastly.Int64(20),
        WafRuleID: "3krg2uUGZzb2W9Euo4moOR",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WafRuleRevisionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListWafRuleRevisionsRequest](../../models/operations/listwafrulerevisionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ListWafRuleRevisionsResponse](../../models/operations/listwafrulerevisionsresponse.md), error**

