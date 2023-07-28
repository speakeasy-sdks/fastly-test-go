# LoggingSumologic

## Overview

Fastly will POST messages to the Sumo Logic account in the format specified in the Sumologic object.

<https://developer.fastly.com/reference/api/logging/sumologic>
### Available Operations

* [CreateLogSumologic](#createlogsumologic) - Create a Sumologic log endpoint
* [DeleteLogSumologic](#deletelogsumologic) - Delete a Sumologic log endpoint
* [GetLogSumologic](#getlogsumologic) - Get a Sumologic log endpoint
* [ListLogSumologic](#listlogsumologic) - List Sumologic log endpoints
* [UpdateLogSumologic](#updatelogsumologic) - Update a Sumologic log endpoint

## CreateLogSumologic

Create a Sumologic for a particular service and version.

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
    operationSecurity := operations.CreateLogSumologicSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSumologic.CreateLogSumologic(ctx, operations.CreateLogSumologicRequest{
        LoggingSumologic2: &shared.LoggingSumologic2{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSumologicFormatVersionTwo.ToPointer(),
            MessageType: shared.LoggingMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingSumologicPlacementWafDebug.ToPointer(),
            ResponseCondition: sdk.String("null"),
            URL: sdk.String("reiciendis"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSumologicResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateLogSumologicRequest](../../models/operations/createlogsumologicrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.CreateLogSumologicSecurity](../../models/operations/createlogsumologicsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.CreateLogSumologicResponse](../../models/operations/createlogsumologicresponse.md), error**


## DeleteLogSumologic

Delete the Sumologic for a particular service and version.

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
    operationSecurity := operations.DeleteLogSumologicSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSumologic.DeleteLogSumologic(ctx, operations.DeleteLogSumologicRequest{
        LoggingSumologicName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogSumologic200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteLogSumologicRequest](../../models/operations/deletelogsumologicrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.DeleteLogSumologicSecurity](../../models/operations/deletelogsumologicsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.DeleteLogSumologicResponse](../../models/operations/deletelogsumologicresponse.md), error**


## GetLogSumologic

Get the Sumologic for a particular service and version.

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
    operationSecurity := operations.GetLogSumologicSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSumologic.GetLogSumologic(ctx, operations.GetLogSumologicRequest{
        LoggingSumologicName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSumologicResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetLogSumologicRequest](../../models/operations/getlogsumologicrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.GetLogSumologicSecurity](../../models/operations/getlogsumologicsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.GetLogSumologicResponse](../../models/operations/getlogsumologicresponse.md), error**


## ListLogSumologic

List all of the Sumologics for a particular service and version.

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
    operationSecurity := operations.ListLogSumologicSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSumologic.ListLogSumologic(ctx, operations.ListLogSumologicRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSumologicResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListLogSumologicRequest](../../models/operations/listlogsumologicrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ListLogSumologicSecurity](../../models/operations/listlogsumologicsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ListLogSumologicResponse](../../models/operations/listlogsumologicresponse.md), error**


## UpdateLogSumologic

Update the Sumologic for a particular service and version.

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
    operationSecurity := operations.UpdateLogSumologicSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingSumologic.UpdateLogSumologic(ctx, operations.UpdateLogSumologicRequest{
        LoggingSumologic2: &shared.LoggingSumologic2{
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingSumologicFormatVersionTwo.ToPointer(),
            MessageType: shared.LoggingMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingSumologicPlacementLessThanNilGreaterThan.ToPointer(),
            ResponseCondition: sdk.String("null"),
            URL: sdk.String("necessitatibus"),
        },
        LoggingSumologicName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingSumologicResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateLogSumologicRequest](../../models/operations/updatelogsumologicrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.UpdateLogSumologicSecurity](../../models/operations/updatelogsumologicsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.UpdateLogSumologicResponse](../../models/operations/updatelogsumologicresponse.md), error**

