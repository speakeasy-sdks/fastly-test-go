# LoggingElasticsearch
(*.LoggingElasticsearch*)

## Overview

Fastly will upload log messages periodically to the server in the format specified in the Elasticsearch object.

<https://developer.fastly.com/reference/api/logging/elasticsearch>
### Available Operations

* [CreateLogElasticsearch](#createlogelasticsearch) - Create an Elasticsearch log endpoint
* [DeleteLogElasticsearch](#deletelogelasticsearch) - Delete an Elasticsearch log endpoint
* [GetLogElasticsearch](#getlogelasticsearch) - Get an Elasticsearch log endpoint
* [ListLogElasticsearch](#listlogelasticsearch) - List Elasticsearch log endpoints
* [UpdateLogElasticsearch](#updatelogelasticsearch) - Update an Elasticsearch log endpoint

## CreateLogElasticsearch

Create a Elasticsearch logging endpoint for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingElasticsearch.CreateLogElasticsearch(ctx, operations.CreateLogElasticsearchRequest{
        LoggingElasticsearch: &components.LoggingElasticsearch{
            FormatVersion: components.LoggingElasticsearchFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: components.LoggingElasticsearchPlacementNone.ToPointer(),
            ResponseCondition: fastly.String("null"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingElasticsearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateLogElasticsearchRequest](../../models/operations/createlogelasticsearchrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.CreateLogElasticsearchResponse](../../models/operations/createlogelasticsearchresponse.md), error**


## DeleteLogElasticsearch

Delete the Elasticsearch logging endpoint for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingElasticsearch.DeleteLogElasticsearch(ctx, operations.DeleteLogElasticsearchRequest{
        LoggingElasticsearchName: "test-log-endpoint",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteLogElasticsearchRequest](../../models/operations/deletelogelasticsearchrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.DeleteLogElasticsearchResponse](../../models/operations/deletelogelasticsearchresponse.md), error**


## GetLogElasticsearch

Get the Elasticsearch logging endpoint for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingElasticsearch.GetLogElasticsearch(ctx, operations.GetLogElasticsearchRequest{
        LoggingElasticsearchName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingElasticsearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetLogElasticsearchRequest](../../models/operations/getlogelasticsearchrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetLogElasticsearchResponse](../../models/operations/getlogelasticsearchresponse.md), error**


## ListLogElasticsearch

List all of the Elasticsearch logging endpoints for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingElasticsearch.ListLogElasticsearch(ctx, operations.ListLogElasticsearchRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListLogElasticsearchRequest](../../models/operations/listlogelasticsearchrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ListLogElasticsearchResponse](../../models/operations/listlogelasticsearchresponse.md), error**


## UpdateLogElasticsearch

Update the Elasticsearch logging endpoint for a particular service and version.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LoggingElasticsearch.UpdateLogElasticsearch(ctx, operations.UpdateLogElasticsearchRequest{
        LoggingElasticsearch: &components.LoggingElasticsearch{
            FormatVersion: components.LoggingElasticsearchFormatVersionTwo.ToPointer(),
            Name: fastly.String("test-log-endpoint"),
            Placement: components.LoggingElasticsearchPlacementWafDebug.ToPointer(),
            ResponseCondition: fastly.String("null"),
        },
        LoggingElasticsearchName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingElasticsearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateLogElasticsearchRequest](../../models/operations/updatelogelasticsearchrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.UpdateLogElasticsearchResponse](../../models/operations/updatelogelasticsearchresponse.md), error**

