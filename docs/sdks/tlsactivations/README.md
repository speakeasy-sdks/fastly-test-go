# TLSActivations

## Overview

TLS activations.

<https://developer.fastly.com/reference/api/tls/custom-certs/activations>
### Available Operations

* [CreateTLSActivation](#createtlsactivation) - Enable TLS for a domain using a custom certificate
* [DeleteTLSActivation](#deletetlsactivation) - Disable TLS on a domain
* [GetTLSActivation](#gettlsactivation) - Get a TLS activation
* [ListTLSActivations](#listtlsactivations) - List TLS activations
* [UpdateTLSActivation](#updatetlsactivation) - Update a certificate

## CreateTLSActivation

Enable TLS for a particular TLS domain and certificate combination. These relationships must be specified to create the TLS activation.

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
    operationSecurity := operations.CreateTLSActivationSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSActivations.CreateTLSActivation(ctx, shared.TLSActivationInput{
        Data: &shared.TLSActivationDataInput{
            Relationships: &shared.RelationshipsForTLSActivationInput{
                TLSCertificate: &shared.RelationshipsForTLSActivationTLSCertificateInput{
                    Data: &shared.RelationshipMemberTLSCertificateInput{
                        Type: shared.TypeTLSCertificateTLSCertificate.ToPointer(),
                    },
                },
                TLSDomain: &shared.RelationshipsForTLSActivationTLSDomainInput{
                    Data: &shared.RelationshipMemberTLSDomainInput{
                        Type: shared.TypeTLSDomainTLSDomain.ToPointer(),
                    },
                },
            },
            Type: shared.TypeTLSActivationTLSActivation.ToPointer(),
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSActivationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [shared.TLSActivationInput](../../models/shared/tlsactivationinput.md)                           | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.CreateTLSActivationSecurity](../../models/operations/createtlsactivationsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.CreateTLSActivationResponse](../../models/operations/createtlsactivationresponse.md), error**


## DeleteTLSActivation

Disable TLS on the domain associated with this TLS activation.

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
    operationSecurity := operations.DeleteTLSActivationSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSActivations.DeleteTLSActivation(ctx, operations.DeleteTLSActivationRequest{
        TLSActivationID: "aCtguUGZzb2W9Euo4moOR",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteTLSActivationRequest](../../models/operations/deletetlsactivationrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.DeleteTLSActivationSecurity](../../models/operations/deletetlsactivationsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.DeleteTLSActivationResponse](../../models/operations/deletetlsactivationresponse.md), error**


## GetTLSActivation

Show a TLS activation.

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
    operationSecurity := operations.GetTLSActivationSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSActivations.GetTLSActivation(ctx, operations.GetTLSActivationRequest{
        Include: sdk.String("tls_certificate,tls_configuration,tls_domain"),
        TLSActivationID: "aCtguUGZzb2W9Euo4moOR",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSActivationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetTLSActivationRequest](../../models/operations/gettlsactivationrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.GetTLSActivationSecurity](../../models/operations/gettlsactivationsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.GetTLSActivationResponse](../../models/operations/gettlsactivationresponse.md), error**


## ListTLSActivations

List all TLS activations.

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
    operationSecurity := operations.ListTLSActivationsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSActivations.ListTLSActivations(ctx, operations.ListTLSActivationsRequest{
        FilterTLSCertificateID: sdk.String("libero"),
        FilterTLSConfigurationID: sdk.String("vitae"),
        FilterTLSDomainID: sdk.String("accusamus"),
        Include: sdk.String("tls_certificate,tls_configuration,tls_domain"),
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSActivationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListTLSActivationsRequest](../../models/operations/listtlsactivationsrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ListTLSActivationsSecurity](../../models/operations/listtlsactivationssecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ListTLSActivationsResponse](../../models/operations/listtlsactivationsresponse.md), error**


## UpdateTLSActivation

Update the certificate used to terminate TLS traffic for the domain associated with this TLS activation.

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
    operationSecurity := operations.UpdateTLSActivationSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSActivations.UpdateTLSActivation(ctx, operations.UpdateTLSActivationRequest{
        TLSActivationInput: &shared.TLSActivationInput{
            Data: &shared.TLSActivationDataInput{
                Relationships: &shared.RelationshipsForTLSActivationInput{
                    TLSCertificate: &shared.RelationshipsForTLSActivationTLSCertificateInput{
                        Data: &shared.RelationshipMemberTLSCertificateInput{
                            Type: shared.TypeTLSCertificateTLSCertificate.ToPointer(),
                        },
                    },
                    TLSDomain: &shared.RelationshipsForTLSActivationTLSDomainInput{
                        Data: &shared.RelationshipMemberTLSDomainInput{
                            Type: shared.TypeTLSDomainTLSDomain.ToPointer(),
                        },
                    },
                },
                Type: shared.TypeTLSActivationTLSActivation.ToPointer(),
            },
        },
        TLSActivationID: "aCtguUGZzb2W9Euo4moOR",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSActivationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateTLSActivationRequest](../../models/operations/updatetlsactivationrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.UpdateTLSActivationSecurity](../../models/operations/updatetlsactivationsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.UpdateTLSActivationResponse](../../models/operations/updatetlsactivationresponse.md), error**

