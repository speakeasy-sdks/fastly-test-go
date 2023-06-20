# LoggingElasticsearch

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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.LoggingElasticsearch.CreateLogElasticsearch(ctx, operations.CreateLogElasticsearchRequest{
        LoggingElasticsearch2: &shared.LoggingElasticsearch2{
            Format: sdk.String("veniam"),
            FormatVersion: shared.LoggingElasticsearchFormatVersionOne.ToPointer(),
            Index: sdk.String("inventore"),
            Name: sdk.String("test-log-endpoint"),
            Password: sdk.String("magnam"),
            Pipeline: sdk.String("ea"),
            Placement: shared.LoggingElasticsearchPlacementLessThanNilGreaterThan.ToPointer(),
            RequestMaxBytes: sdk.Int64(232234),
            RequestMaxEntries: sdk.Int64(926213),
            ResponseCondition: sdk.String("null"),
            TLSCaCert: sdk.String("aspernatur"),
            TLSClientCert: sdk.String("minima"),
            TLSClientKey: sdk.String("eaque"),
            TLSHostname: sdk.String("a"),
            URL: sdk.String("libero"),
            User: sdk.String("aut"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateLogElasticsearchSecurity{
        Token: "",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateLogElasticsearchRequest](../../models/operations/createlogelasticsearchrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.CreateLogElasticsearchSecurity](../../models/operations/createlogelasticsearchsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.LoggingElasticsearch.DeleteLogElasticsearch(ctx, operations.DeleteLogElasticsearchRequest{
        LoggingElasticsearchName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteLogElasticsearchSecurity{
        Token: "",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DeleteLogElasticsearchRequest](../../models/operations/deletelogelasticsearchrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.DeleteLogElasticsearchSecurity](../../models/operations/deletelogelasticsearchsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.LoggingElasticsearch.GetLogElasticsearch(ctx, operations.GetLogElasticsearchRequest{
        LoggingElasticsearchName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetLogElasticsearchSecurity{
        Token: "",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetLogElasticsearchRequest](../../models/operations/getlogelasticsearchrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.GetLogElasticsearchSecurity](../../models/operations/getlogelasticsearchsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.LoggingElasticsearch.ListLogElasticsearch(ctx, operations.ListLogElasticsearchRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListLogElasticsearchSecurity{
        Token: "",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListLogElasticsearchRequest](../../models/operations/listlogelasticsearchrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.ListLogElasticsearchSecurity](../../models/operations/listlogelasticsearchsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.LoggingElasticsearch.UpdateLogElasticsearch(ctx, operations.UpdateLogElasticsearchRequest{
        LoggingElasticsearch2: &shared.LoggingElasticsearch2{
            Format: sdk.String("aut"),
            FormatVersion: shared.LoggingElasticsearchFormatVersionTwo.ToPointer(),
            Index: sdk.String("impedit"),
            Name: sdk.String("test-log-endpoint"),
            Password: sdk.String("aliquam"),
            Pipeline: sdk.String("fugit"),
            Placement: shared.LoggingElasticsearchPlacementLessThanNilGreaterThan.ToPointer(),
            RequestMaxBytes: sdk.Int64(79522),
            RequestMaxEntries: sdk.Int64(250622),
            ResponseCondition: sdk.String("null"),
            TLSCaCert: sdk.String("et"),
            TLSClientCert: sdk.String("dolorum"),
            TLSClientKey: sdk.String("laborum"),
            TLSHostname: sdk.String("placeat"),
            URL: sdk.String("velit"),
            User: sdk.String("eum"),
        },
        LoggingElasticsearchName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.UpdateLogElasticsearchSecurity{
        Token: "",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateLogElasticsearchRequest](../../models/operations/updatelogelasticsearchrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.UpdateLogElasticsearchSecurity](../../models/operations/updatelogelasticsearchsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.UpdateLogElasticsearchResponse](../../models/operations/updatelogelasticsearchresponse.md), error**

