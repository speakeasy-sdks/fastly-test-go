# LoggingBigquery

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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.LoggingBigquery.CreateLogBigquery(ctx, operations.CreateLogBigqueryRequest{
        LoggingBigquery2: &shared.LoggingBigquery2{
            AccountName: sdk.String("test-user@test-project-id.iam.gserviceaccount.com"),
            Dataset: sdk.String("voluptatibus"),
            Format: sdk.String("quisquam"),
            FormatVersion: shared.LoggingBigqueryFormatVersionTwo.ToPointer(),
            Name: sdk.String("Tim Erdman"),
            Placement: shared.LoggingBigqueryPlacementNone.ToPointer(),
            ProjectID: sdk.String("test-project-id"),
            ResponseCondition: sdk.String("null"),
            SecretKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            Table: sdk.String("vero"),
            TemplateSuffix: sdk.String("tenetur"),
            User: sdk.String("test-user@test-project-id.iam.gserviceaccount.com"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateLogBigquerySecurity{
        Token: "",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateLogBigqueryRequest](../../models/operations/createlogbigqueryrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.CreateLogBigquerySecurity](../../models/operations/createlogbigquerysecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.LoggingBigquery.DeleteLogBigquery(ctx, operations.DeleteLogBigqueryRequest{
        LoggingBigqueryName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteLogBigquerySecurity{
        Token: "",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteLogBigqueryRequest](../../models/operations/deletelogbigqueryrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.DeleteLogBigquerySecurity](../../models/operations/deletelogbigquerysecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.LoggingBigquery.GetLogBigquery(ctx, operations.GetLogBigqueryRequest{
        LoggingBigqueryName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetLogBigquerySecurity{
        Token: "",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetLogBigqueryRequest](../../models/operations/getlogbigqueryrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetLogBigquerySecurity](../../models/operations/getlogbigquerysecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.LoggingBigquery.ListLogBigquery(ctx, operations.ListLogBigqueryRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListLogBigquerySecurity{
        Token: "",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListLogBigqueryRequest](../../models/operations/listlogbigqueryrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ListLogBigquerySecurity](../../models/operations/listlogbigquerysecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.LoggingBigquery.UpdateLogBigquery(ctx, operations.UpdateLogBigqueryRequest{
        LoggingBigquery2: &shared.LoggingBigquery2{
            AccountName: sdk.String("test-user@test-project-id.iam.gserviceaccount.com"),
            Dataset: sdk.String("dignissimos"),
            Format: sdk.String("hic"),
            FormatVersion: shared.LoggingBigqueryFormatVersionTwo.ToPointer(),
            Name: sdk.String("Lonnie Murray"),
            Placement: shared.LoggingBigqueryPlacementWafDebug.ToPointer(),
            ProjectID: sdk.String("test-project-id"),
            ResponseCondition: sdk.String("null"),
            SecretKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            Table: sdk.String("dolore"),
            TemplateSuffix: sdk.String("quibusdam"),
            User: sdk.String("test-user@test-project-id.iam.gserviceaccount.com"),
        },
        LoggingBigqueryName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.UpdateLogBigquerySecurity{
        Token: "",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateLogBigqueryRequest](../../models/operations/updatelogbigqueryrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UpdateLogBigquerySecurity](../../models/operations/updatelogbigquerysecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.UpdateLogBigqueryResponse](../../models/operations/updatelogbigqueryresponse.md), error**
