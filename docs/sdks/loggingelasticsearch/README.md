# LoggingElasticsearch
(*LoggingElasticsearch*)

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
    res, err := s.LoggingElasticsearch.CreateLogElasticsearch(ctx, operations.CreateLogElasticsearchRequest{
        LoggingElasticsearch: &shared.LoggingElasticsearch{
            Format: fastly.String("Sedan"),
            FormatVersion: shared.LoggingElasticsearchFormatVersionTwo.ToPointer(),
            Index: fastly.String("Mouse"),
            Name: fastly.String("test-log-endpoint"),
            Password: fastly.String("L0LtOVDJ_cyv_Gf"),
            Pipeline: fastly.String("leverage CSS"),
            Placement: shared.LoggingElasticsearchPlacementWafDebug.ToPointer(),
            RequestMaxBytes: fastly.Int64(733399),
            RequestMaxEntries: fastly.Int64(237272),
            ResponseCondition: fastly.String("null"),
            TLSCaCert: fastly.String("thistle purple"),
            TLSClientCert: fastly.String("superstructure"),
            TLSClientKey: fastly.String("Cayman Court Bicycle"),
            TLSHostname: fastly.String("mobile Croatian paradigms"),
            URL: fastly.String("http://trained-dusk.info"),
            User: fastly.String("Raul_Ankunding"),
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
    res, err := s.LoggingElasticsearch.DeleteLogElasticsearch(ctx, operations.DeleteLogElasticsearchRequest{
        LoggingElasticsearchName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogElasticsearch200ApplicationJSONObject != nil {
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
    res, err := s.LoggingElasticsearch.ListLogElasticsearch(ctx, operations.ListLogElasticsearchRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingElasticsearchResponses != nil {
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
    res, err := s.LoggingElasticsearch.UpdateLogElasticsearch(ctx, operations.UpdateLogElasticsearchRequest{
        LoggingElasticsearch: &shared.LoggingElasticsearch{
            Format: fastly.String("grenade yellow"),
            FormatVersion: shared.LoggingElasticsearchFormatVersionOne.ToPointer(),
            Index: fastly.String("HDD"),
            Name: fastly.String("test-log-endpoint"),
            Password: fastly.String("vvCrseYZg_P55Bi"),
            Pipeline: fastly.String("Handmade"),
            Placement: shared.LoggingElasticsearchPlacementNone.ToPointer(),
            RequestMaxBytes: fastly.Int64(534662),
            RequestMaxEntries: fastly.Int64(978553),
            ResponseCondition: fastly.String("null"),
            TLSCaCert: fastly.String("reciprocal PCI Wagon"),
            TLSClientCert: fastly.String("snare experiences"),
            TLSClientKey: fastly.String("Seychelles muffled Freeda"),
            TLSHostname: fastly.String("bypass"),
            URL: fastly.String("https://well-worn-maker.com"),
            User: fastly.String("Jamil_Bauch53"),
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

