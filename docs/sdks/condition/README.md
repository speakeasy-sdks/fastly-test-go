# Condition
(*Condition*)

## Overview

Conditions are used to control whether logic defined in configured VCL objects is applied for a particular client request. A condition contains a VCL conditional expression that evaluates to either true or false and is used to determine whether the condition is met. The type of the condition determines where it is executed and the VCL variables that can be evaluated as part of the conditional logic.

<https://developer.fastly.com/reference/api/vcl-services/condition>
### Available Operations

* [CreateCondition](#createcondition) - Create a condition
* [DeleteCondition](#deletecondition) - Delete a condition
* [GetCondition](#getcondition) - Describe a condition
* [ListConditions](#listconditions) - List conditions
* [UpdateCondition](#updatecondition) - Update a condition

## CreateCondition

Creates a new condition.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Condition.CreateCondition(ctx, operations.CreateConditionRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConditionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateConditionRequest](../../models/operations/createconditionrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateConditionResponse](../../models/operations/createconditionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteCondition

Deletes the specified condition.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Condition.DeleteCondition(ctx, operations.DeleteConditionRequest{
        ConditionName: "test-condition",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteConditionRequest](../../models/operations/deleteconditionrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DeleteConditionResponse](../../models/operations/deleteconditionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCondition

Gets the specified condition.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Condition.GetCondition(ctx, operations.GetConditionRequest{
        ConditionName: "test-condition",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConditionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetConditionRequest](../../models/operations/getconditionrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetConditionResponse](../../models/operations/getconditionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListConditions

Gets all conditions for a particular service and version.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Condition.ListConditions(ctx, operations.ListConditionsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConditionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListConditionsRequest](../../models/operations/listconditionsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListConditionsResponse](../../models/operations/listconditionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateCondition

Updates the specified condition.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.Condition.UpdateCondition(ctx, operations.UpdateConditionRequest{
        ConditionName: "test-condition",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConditionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateConditionRequest](../../models/operations/updateconditionrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateConditionResponse](../../models/operations/updateconditionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
