# LoggingBigquery
(*LoggingBigquery*)

## Overview

Fastly will upload log messages to the Google BigQuery dataset and table in the format specified in the BigQuery logging object.

<https://developer.fastly.com/reference/api/logging/bigquery>
### Available Operations

* [CreateLogBigquery](#createlogbigquery) - Create a BigQuery log endpoint
* [DeleteLogBigquery](#deletelogbigquery) - Delete a BigQuery log endpoint
* [GetLogBigquery](#getlogbigquery) - Get a BigQuery log endpoint
* [ListLogBigquery](#listlogbigquery) - List BigQuery log endpoints
* [UpdateLogBigquery](#updatelogbigquery) - Update a BigQuery log endpoint

## CreateLogBigquery

Create a BigQuery logging object for a particular service and version.

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
    res, err := s.LoggingBigquery.CreateLogBigquery(ctx, operations.CreateLogBigqueryRequest{
        LoggingBigquery2: &shared.LoggingBigquery2{
            AccountName: fastly.String("test-user@test-project-id.iam.gserviceaccount.com"),
            Dataset: fastly.String("qui"),
            Format: fastly.String("ipsum"),
            FormatVersion: shared.LoggingBigqueryFormatVersionTwo.ToPointer(),
            Name: fastly.String("Felipe Klein"),
            Placement: shared.LoggingBigqueryPlacementNone.ToPointer(),
            ProjectID: fastly.String("test-project-id"),
            ResponseCondition: fastly.String("null"),
            SecretKey: fastly.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            Table: fastly.String("dolorum"),
            TemplateSuffix: fastly.String("numquam"),
            User: fastly.String("test-user@test-project-id.iam.gserviceaccount.com"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingBigqueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateLogBigqueryRequest](../../models/operations/createlogbigqueryrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateLogBigqueryResponse](../../models/operations/createlogbigqueryresponse.md), error**


## DeleteLogBigquery

Delete a BigQuery logging object for a particular service and version.

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
    res, err := s.LoggingBigquery.DeleteLogBigquery(ctx, operations.DeleteLogBigqueryRequest{
        LoggingBigqueryName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogBigquery200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteLogBigqueryRequest](../../models/operations/deletelogbigqueryrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteLogBigqueryResponse](../../models/operations/deletelogbigqueryresponse.md), error**


## GetLogBigquery

Get the details for a BigQuery logging object for a particular service and version.

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
    res, err := s.LoggingBigquery.GetLogBigquery(ctx, operations.GetLogBigqueryRequest{
        LoggingBigqueryName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingBigqueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetLogBigqueryRequest](../../models/operations/getlogbigqueryrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetLogBigqueryResponse](../../models/operations/getlogbigqueryresponse.md), error**


## ListLogBigquery

List all of the BigQuery logging objects for a particular service and version.

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
    res, err := s.LoggingBigquery.ListLogBigquery(ctx, operations.ListLogBigqueryRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingBigqueryResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListLogBigqueryRequest](../../models/operations/listlogbigqueryrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListLogBigqueryResponse](../../models/operations/listlogbigqueryresponse.md), error**


## UpdateLogBigquery

Update a BigQuery logging object for a particular service and version.

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
    res, err := s.LoggingBigquery.UpdateLogBigquery(ctx, operations.UpdateLogBigqueryRequest{
        LoggingBigquery2: &shared.LoggingBigquery2{
            AccountName: fastly.String("test-user@test-project-id.iam.gserviceaccount.com"),
            Dataset: fastly.String("veritatis"),
            Format: fastly.String("ipsa"),
            FormatVersion: shared.LoggingBigqueryFormatVersionOne.ToPointer(),
            Name: fastly.String("Viola Hahn"),
            Placement: shared.LoggingBigqueryPlacementLessThanNilGreaterThan.ToPointer(),
            ProjectID: fastly.String("test-project-id"),
            ResponseCondition: fastly.String("null"),
            SecretKey: fastly.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            Table: fastly.String("voluptas"),
            TemplateSuffix: fastly.String("natus"),
            User: fastly.String("test-user@test-project-id.iam.gserviceaccount.com"),
        },
        LoggingBigqueryName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingBigqueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateLogBigqueryRequest](../../models/operations/updatelogbigqueryrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateLogBigqueryResponse](../../models/operations/updatelogbigqueryresponse.md), error**

