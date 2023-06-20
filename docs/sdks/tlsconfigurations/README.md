# TLSConfigurations

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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.TLSConfigurations.GetTLSConfig(ctx, operations.GetTLSConfigRequest{
        Include: sdk.String("dns_records"),
        TLSConfigurationID: "t7CguUGZzb2W9Euo5FoKa",
    }, operations.GetTLSConfigSecurity{
        Token: "",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetTLSConfigRequest](../../models/operations/gettlsconfigrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.GetTLSConfigSecurity](../../models/operations/gettlsconfigsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.TLSConfigurations.ListTLSConfigs(ctx, operations.ListTLSConfigsRequest{
        FilterBulk: sdk.String("eum"),
        Include: sdk.String("dns_records"),
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
    }, operations.ListTLSConfigsSecurity{
        Token: "",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListTLSConfigsRequest](../../models/operations/listtlsconfigsrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.ListTLSConfigsSecurity](../../models/operations/listtlsconfigssecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.TLSConfigurations.UpdateTLSConfig(ctx, operations.UpdateTLSConfigRequest{
        TLSConfigurationInput: &shared.TLSConfigurationInput{
            Data: &shared.TLSConfigurationDataInput{
                Attributes: &shared.TLSConfigurationDataAttributes{
                    Name: sdk.String("Deanna Swaniawski"),
                },
                Relationships: &shared.TLSConfigurationDataRelationships2Input{
                    DNSRecords: &shared.TLSConfigurationDataRelationships2DNSRecordsInput{
                        Data: []shared.RelationshipMemberTLSDNSRecordInput{
                            shared.RelationshipMemberTLSDNSRecordInput{
                                Type: shared.TypeTLSDNSRecordDNSRecord.ToPointer(),
                            },
                            shared.RelationshipMemberTLSDNSRecordInput{
                                Type: shared.TypeTLSDNSRecordDNSRecord.ToPointer(),
                            },
                        },
                    },
                },
                Type: shared.TypeTLSConfigurationTLSConfiguration.ToPointer(),
            },
        },
        TLSConfigurationID: "t7CguUGZzb2W9Euo5FoKa",
    }, operations.UpdateTLSConfigSecurity{
        Token: "",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateTLSConfigRequest](../../models/operations/updatetlsconfigrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.UpdateTLSConfigSecurity](../../models/operations/updatetlsconfigsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.UpdateTLSConfigResponse](../../models/operations/updatetlsconfigresponse.md), error**

