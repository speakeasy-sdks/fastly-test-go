# TLSConfigurations
(*TLSConfigurations*)

## Overview

Customers with access to multiple sets of IP pools are able to apply different configuration options to their TLS enabled domains.

<https://developer.fastly.com/reference/api/tls/configuration>
### Available Operations

* [GetTLSConfig](#gettlsconfig) - Get a TLS configuration
* [ListTLSConfigs](#listtlsconfigs) - List TLS configurations
* [UpdateTLSConfig](#updatetlsconfig) - Update a TLS configuration

## GetTLSConfig

Show a TLS configuration.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.TLSConfigurations.GetTLSConfig(ctx, operations.GetTLSConfigRequest{
        Include: fastlytestgo.String("dns_records"),
        TLSConfigurationID: "t7CguUGZzb2W9Euo5FoKa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSConfigurationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetTLSConfigRequest](../../models/operations/gettlsconfigrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetTLSConfigResponse](../../models/operations/gettlsconfigresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListTLSConfigs

List all TLS configurations.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.TLSConfigurations.ListTLSConfigs(ctx, operations.ListTLSConfigsRequest{
        Include: fastlytestgo.String("dns_records"),
        PageNumber: fastlytestgo.Int64(1),
        PageSize: fastlytestgo.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSConfigurationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListTLSConfigsRequest](../../models/operations/listtlsconfigsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListTLSConfigsResponse](../../models/operations/listtlsconfigsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateTLSConfig

Update a TLS configuration.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.TLSConfigurations.UpdateTLSConfig(ctx, operations.UpdateTLSConfigRequest{
        TLSConfiguration: &components.TLSConfiguration{
            Data: &components.TLSConfigurationData{
                Attributes: &components.TLSConfigurationDataAttributes{},
                Relationships: components.CreateRelationshipsForTLSConfigurationInputRelationshipsForTLSConfiguration1Input(
                        components.RelationshipsForTLSConfiguration1Input{
                            Service: &components.RelationshipMemberServiceInput{},
                        },
                ),
            },
        },
        TLSConfigurationID: "t7CguUGZzb2W9Euo5FoKa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSConfigurationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateTLSConfigRequest](../../models/operations/updatetlsconfigrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateTLSConfigResponse](../../models/operations/updatetlsconfigresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
