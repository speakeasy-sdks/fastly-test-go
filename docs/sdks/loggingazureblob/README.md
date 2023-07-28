# LoggingAzureblob

## Overview

Fastly will upload log messages to the Azure Blob Storage container in the format specified in the Azure Blob object.

<https://developer.fastly.com/reference/api/logging/azureblob>
### Available Operations

* [CreateLogAzure](#createlogazure) - Create an Azure Blob Storage log endpoint
* [DeleteLogAzure](#deletelogazure) - Delete the Azure Blob Storage log endpoint
* [GetLogAzure](#getlogazure) - Get an Azure Blob Storage log endpoint
* [ListLogAzure](#listlogazure) - List Azure Blob Storage log endpoints
* [UpdateLogAzure](#updatelogazure) - Update an Azure Blob Storage log endpoint

## CreateLogAzure

Create an Azure Blob Storage logging endpoint for a particular service and version.

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
    operationSecurity := operations.CreateLogAzureSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingAzureblob.CreateLogAzure(ctx, operations.CreateLogAzureRequest{
        LoggingAzureblobInput: &shared.LoggingAzureblobInput{
            AccountName: sdk.String("aut"),
            CompressionCodec: shared.LoggingAzureblobCompressionCodecSnappy.ToPointer(),
            Container: sdk.String("itaque"),
            FileMaxBytes: sdk.Int64(9240),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingAzureblobFormatVersionTwo.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingAzureblobMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("repellendus"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingAzureblobPlacementLessThanNilGreaterThan.ToPointer(),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            SasToken: sdk.String("doloribus"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingAzureblobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateLogAzureRequest](../../models/operations/createlogazurerequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.CreateLogAzureSecurity](../../models/operations/createlogazuresecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.CreateLogAzureResponse](../../models/operations/createlogazureresponse.md), error**


## DeleteLogAzure

Delete the Azure Blob Storage logging endpoint for a particular service and version.

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
    operationSecurity := operations.DeleteLogAzureSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingAzureblob.DeleteLogAzure(ctx, operations.DeleteLogAzureRequest{
        LoggingAzureblobName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteLogAzure200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteLogAzureRequest](../../models/operations/deletelogazurerequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.DeleteLogAzureSecurity](../../models/operations/deletelogazuresecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.DeleteLogAzureResponse](../../models/operations/deletelogazureresponse.md), error**


## GetLogAzure

Get the Azure Blob Storage logging endpoint for a particular service and version.

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
    operationSecurity := operations.GetLogAzureSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingAzureblob.GetLogAzure(ctx, operations.GetLogAzureRequest{
        LoggingAzureblobName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingAzureblobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetLogAzureRequest](../../models/operations/getlogazurerequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.GetLogAzureSecurity](../../models/operations/getlogazuresecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.GetLogAzureResponse](../../models/operations/getlogazureresponse.md), error**


## ListLogAzure

List all of the Azure Blob Storage logging endpoints for a particular service and version.

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
    operationSecurity := operations.ListLogAzureSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingAzureblob.ListLogAzure(ctx, operations.ListLogAzureRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingAzureblobResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListLogAzureRequest](../../models/operations/listlogazurerequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.ListLogAzureSecurity](../../models/operations/listlogazuresecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.ListLogAzureResponse](../../models/operations/listlogazureresponse.md), error**


## UpdateLogAzure

Update the Azure Blob Storage logging endpoint for a particular service and version.

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
    operationSecurity := operations.UpdateLogAzureSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.LoggingAzureblob.UpdateLogAzure(ctx, operations.UpdateLogAzureRequest{
        LoggingAzureblobInput: &shared.LoggingAzureblobInput{
            AccountName: sdk.String("ut"),
            CompressionCodec: shared.LoggingAzureblobCompressionCodecGzip.ToPointer(),
            Container: sdk.String("cupiditate"),
            FileMaxBytes: sdk.Int64(181631),
            Format: sdk.String("%h %l %u %t "%r" %&gt;s %b"),
            FormatVersion: shared.LoggingAzureblobFormatVersionOne.ToPointer(),
            GzipLevel: sdk.Int64(0),
            MessageType: shared.LoggingAzureblobMessageTypeClassic.ToPointer(),
            Name: sdk.String("test-log-endpoint"),
            Path: sdk.String("laudantium"),
            Period: sdk.Int64(3600),
            Placement: shared.LoggingAzureblobPlacementWafDebug.ToPointer(),
            PublicKey: sdk.String("-----BEGIN PRIVATE KEY-----
        ...
        -----END PRIVATE KEY-----
        "),
            ResponseCondition: sdk.String("null"),
            SasToken: sdk.String("occaecati"),
        },
        LoggingAzureblobName: "test-log-endpoint",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggingAzureblobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateLogAzureRequest](../../models/operations/updatelogazurerequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.UpdateLogAzureSecurity](../../models/operations/updatelogazuresecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.UpdateLogAzureResponse](../../models/operations/updatelogazureresponse.md), error**

