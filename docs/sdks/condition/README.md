# Condition

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
	"context"
	"log"
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Condition.CreateCondition(ctx, operations.CreateConditionRequest{
        ConditionInput: &shared.ConditionInput{
            Comment: sdk.String("excepturi"),
            Name: sdk.String("test-condition"),
            Priority: sdk.String("10"),
            Statement: sdk.String("pariatur"),
            Type: shared.ConditionTypeCache.ToPointer(),
            Version: sdk.String("praesentium"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateConditionSecurity{
        Token: "",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateConditionRequest](../../models/operations/createconditionrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.CreateConditionSecurity](../../models/operations/createconditionsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.CreateConditionResponse](../../models/operations/createconditionresponse.md), error**


## DeleteCondition

Deletes the specified condition.

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
    res, err := s.Condition.DeleteCondition(ctx, operations.DeleteConditionRequest{
        ConditionName: "test-condition",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteConditionSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteCondition200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteConditionRequest](../../models/operations/deleteconditionrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.DeleteConditionSecurity](../../models/operations/deleteconditionsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.DeleteConditionResponse](../../models/operations/deleteconditionresponse.md), error**


## GetCondition

Gets the specified condition.

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
    res, err := s.Condition.GetCondition(ctx, operations.GetConditionRequest{
        ConditionName: "test-condition",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetConditionSecurity{
        Token: "",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetConditionRequest](../../models/operations/getconditionrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.GetConditionSecurity](../../models/operations/getconditionsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.GetConditionResponse](../../models/operations/getconditionresponse.md), error**


## ListConditions

Gets all conditions for a particular service and version.

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
    res, err := s.Condition.ListConditions(ctx, operations.ListConditionsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListConditionsSecurity{
        Token: "",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListConditionsRequest](../../models/operations/listconditionsrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.ListConditionsSecurity](../../models/operations/listconditionssecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.ListConditionsResponse](../../models/operations/listconditionsresponse.md), error**


## UpdateCondition

Updates the specified condition.

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
    res, err := s.Condition.UpdateCondition(ctx, operations.UpdateConditionRequest{
        ConditionInput: &shared.ConditionInput{
            Comment: sdk.String("rem"),
            Name: sdk.String("test-condition"),
            Priority: sdk.String("10"),
            Statement: sdk.String("voluptates"),
            Type: shared.ConditionTypeRequest.ToPointer(),
            Version: sdk.String("repudiandae"),
        },
        ConditionName: "test-condition",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.UpdateConditionSecurity{
        Token: "",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateConditionRequest](../../models/operations/updateconditionrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.UpdateConditionSecurity](../../models/operations/updateconditionsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.UpdateConditionResponse](../../models/operations/updateconditionresponse.md), error**
