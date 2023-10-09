# LoggingKafka
(*LoggingKafka*)

## Overview

Fastly will upload log messages periodically to the server in the format specified in the Kafka object.

<https://developer.fastly.com/reference/api/logging/kafka>
### Available Operations

* [CreateLogKafka](#createlogkafka) - Create a Kafka log endpoint
* [DeleteLogKafka](#deletelogkafka) - Delete the Kafka log endpoint
* [GetLogKafka](#getlogkafka) - Get a Kafka log endpoint
* [ListLogKafka](#listlogkafka) - List Kafka log endpoints

## CreateLogKafka

Create a Kafka logging endpoint for a particular service and version.

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
    res, err := s.LoggingKafka.CreateLogKafka(ctx, operations.CreateLogKafkaRequest{
        LoggingKafka: &shared.LoggingKafka{
            Format: fastly.String("%h %l %u %t \"%r\" %&gt;s %b"),
            FormatVersion: shared.LoggingKafkaFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: shared.LoggingKafkaPlacementNone.ToPointer(),
            ResponseCondition: fastly.String("null"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingKafkaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateLogKafkaRequest](../../models/operations/createlogkafkarequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateLogKafkaResponse](../../models/operations/createlogkafkaresponse.md), error**


## DeleteLogKafka

Delete the Kafka logging endpoint for a particular service and version.

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
    res, err := s.LoggingKafka.DeleteLogKafka(ctx, operations.DeleteLogKafkaRequest{
        LoggingKafkaName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogKafka200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteLogKafkaRequest](../../models/operations/deletelogkafkarequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteLogKafkaResponse](../../models/operations/deletelogkafkaresponse.md), error**


## GetLogKafka

Get the Kafka logging endpoint for a particular service and version.

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
    res, err := s.LoggingKafka.GetLogKafka(ctx, operations.GetLogKafkaRequest{
        LoggingKafkaName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingKafkaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetLogKafkaRequest](../../models/operations/getlogkafkarequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetLogKafkaResponse](../../models/operations/getlogkafkaresponse.md), error**


## ListLogKafka

List all of the Kafka logging endpoints for a particular service and version.

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
    res, err := s.LoggingKafka.ListLogKafka(ctx, operations.ListLogKafkaRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingKafkaResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListLogKafkaRequest](../../models/operations/listlogkafkarequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListLogKafkaResponse](../../models/operations/listlogkafkaresponse.md), error**

