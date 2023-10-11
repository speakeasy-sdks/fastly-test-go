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
	"context"
	"log"
	fastly "Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.TLSConfigurations.GetTLSConfig(ctx, operations.GetTLSConfigRequest{
        Include: fastly.String("dns_records"),
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


## ListTLSConfigs

List all TLS configurations.

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.TLSConfigurations.ListTLSConfigs(ctx, operations.ListTLSConfigsRequest{
        Include: fastly.String("dns_records"),
        PageNumber: fastly.Int64(1),
        PageSize: fastly.Int64(20),
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


## UpdateTLSConfig

Update a TLS configuration.

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
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.TLSConfigurations.UpdateTLSConfig(ctx, operations.UpdateTLSConfigRequest{
        TLSConfiguration: &shared.TLSConfiguration{
            Data: &shared.TLSConfigurationData{
                Attributes: &shared.TLSConfigurationDataAttributes{},
                Relationships: shared.CreateRelationshipsForTLSConfigurationInputRelationshipsForTLSConfiguration1Input(
                        shared.RelationshipsForTLSConfiguration1Input{
                            Service: &shared.RelationshipMemberServiceInput{},
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

