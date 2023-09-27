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
    res, err := s.Condition.CreateCondition(ctx, operations.CreateConditionRequest{
        ConditionInput: &shared.ConditionInput{
            Comment: fastly.String("harum"),
            Name: fastly.String("test-condition"),
            Priority: fastly.String("10"),
            Statement: fastly.String("enim"),
            Type: shared.ConditionTypePrefetch.ToPointer(),
            Version: fastly.String("commodi"),
        },
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


## DeleteCondition

Deletes the specified condition.

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
    res, err := s.Condition.DeleteCondition(ctx, operations.DeleteConditionRequest{
        ConditionName: "test-condition",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteConditionRequest](../../models/operations/deleteconditionrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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


## ListConditions

Gets all conditions for a particular service and version.

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


## UpdateCondition

Updates the specified condition.

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
    res, err := s.Condition.UpdateCondition(ctx, operations.UpdateConditionRequest{
        ConditionInput: &shared.ConditionInput{
            Comment: fastly.String("repudiandae"),
            Name: fastly.String("test-condition"),
            Priority: fastly.String("10"),
            Statement: fastly.String("quae"),
            Type: shared.ConditionTypeRequest.ToPointer(),
            Version: fastly.String("quidem"),
        },
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

