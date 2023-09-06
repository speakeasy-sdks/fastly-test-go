# LoggingKinesis

## Overview

Fastly will publish log messages to an Amazon Kinesis stream in the format specified in the Amazon Kinesis Data Streams logging object.

<https://developer.fastly.com/reference/api/logging/kinesis>
### Available Operations

* [CreateLogKinesis](#createlogkinesis) - Create  an Amazon Kinesis log endpoint
* [DeleteLogKinesis](#deletelogkinesis) - Delete the Amazon Kinesis log endpoint
* [GetLogKinesis](#getlogkinesis) - Get an Amazon Kinesis log endpoint
* [ListLogKinesis](#listlogkinesis) - List Amazon Kinesis log endpoints

## CreateLogKinesis

Create an Amazon Kinesis Data Streams logging object for a particular service and version.

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
    operationSecurity := operations.CreateLogKinesisSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingKinesis.CreateLogKinesis(ctx, operations.CreateLogKinesisRequest{
        LoggingKinesis: &shared.LoggingKinesis{
            AccessKey: sdk.String("possimus"),
            Format: sdk.String("magnam"),
            FormatVersion: shared.LoggingFormatVersionOne.ToPointer(),
            IamRole: sdk.String("ex"),
            Name: sdk.String("test-log-endpoint"),
            Placement: shared.LoggingPlacementWafDebug.ToPointer(),
            Region: shared.AwsRegionUsWest1.ToPointer(),
            SecretKey: sdk.String("dolor"),
            Topic: sdk.String("maiores"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingKinesisResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateLogKinesisRequest](../../models/operations/createlogkinesisrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.CreateLogKinesisSecurity](../../models/operations/createlogkinesissecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.CreateLogKinesisResponse](../../models/operations/createlogkinesisresponse.md), error**


## DeleteLogKinesis

Delete an Amazon Kinesis Data Streams logging object for a particular service and version.

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
    operationSecurity := operations.DeleteLogKinesisSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingKinesis.DeleteLogKinesis(ctx, operations.DeleteLogKinesisRequest{
        LoggingKinesisName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogKinesis200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteLogKinesisRequest](../../models/operations/deletelogkinesisrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.DeleteLogKinesisSecurity](../../models/operations/deletelogkinesissecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.DeleteLogKinesisResponse](../../models/operations/deletelogkinesisresponse.md), error**


## GetLogKinesis

Get the details for an Amazon Kinesis Data Streams logging object for a particular service and version.

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
    operationSecurity := operations.GetLogKinesisSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingKinesis.GetLogKinesis(ctx, operations.GetLogKinesisRequest{
        LoggingKinesisName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingKinesisResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetLogKinesisRequest](../../models/operations/getlogkinesisrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.GetLogKinesisSecurity](../../models/operations/getlogkinesissecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.GetLogKinesisResponse](../../models/operations/getlogkinesisresponse.md), error**


## ListLogKinesis

List all of the Amazon Kinesis Data Streams logging objects for a particular service and version.

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
    operationSecurity := operations.ListLogKinesisSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingKinesis.ListLogKinesis(ctx, operations.ListLogKinesisRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingKinesisResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListLogKinesisRequest](../../models/operations/listlogkinesisrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.ListLogKinesisSecurity](../../models/operations/listlogkinesissecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.ListLogKinesisResponse](../../models/operations/listlogkinesisresponse.md), error**

