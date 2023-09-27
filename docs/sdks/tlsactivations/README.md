# TLSActivations
(*TLSActivations*)

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
	fastly "Fastly"
	"Fastly/pkg/models/shared"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TLSActivations.CreateTLSActivation(ctx, shared.TLSActivationInput{
        Data: &shared.TLSActivationDataInput{
            Relationships: &shared.RelationshipsForTLSActivationInput{
                TLSCertificate: &shared.RelationshipsForTLSActivationTLSCertificateInput{
                    Data: &shared.RelationshipMemberTLSCertificateInput{
                        Type: fastly.String("adipisci"),
                    },
                },
                TLSDomain: &shared.RelationshipsForTLSActivationTLSDomainInput{
                    Data: &shared.RelationshipMemberTLSDomainInput{
                        Type: fastly.String("cum"),
                    },
                },
            },
            Type: fastly.String("blanditiis"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSActivationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.TLSActivationInput](../../models/shared/tlsactivationinput.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


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
    res, err := s.TLSActivations.DeleteTLSActivation(ctx, operations.DeleteTLSActivationRequest{
        TLSActivationID: "aCtguUGZzb2W9Euo4moOR",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteTLSActivationRequest](../../models/operations/deletetlsactivationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


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
    res, err := s.TLSActivations.GetTLSActivation(ctx, operations.GetTLSActivationRequest{
        Include: fastly.String("tls_certificate,tls_configuration,tls_domain"),
        TLSActivationID: "aCtguUGZzb2W9Euo4moOR",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSActivationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetTLSActivationRequest](../../models/operations/gettlsactivationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
    res, err := s.TLSActivations.ListTLSActivations(ctx, operations.ListTLSActivationsRequest{
        FilterTLSCertificateID: fastly.String("quas"),
        FilterTLSConfigurationID: fastly.String("hic"),
        FilterTLSDomainID: fastly.String("nesciunt"),
        Include: fastly.String("tls_certificate,tls_configuration,tls_domain"),
        PageNumber: fastly.Int64(1),
        PageSize: fastly.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSActivationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListTLSActivationsRequest](../../models/operations/listtlsactivationsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


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
    res, err := s.TLSActivations.UpdateTLSActivation(ctx, operations.UpdateTLSActivationRequest{
        TLSActivationInput: &shared.TLSActivationInput{
            Data: &shared.TLSActivationDataInput{
                Relationships: &shared.RelationshipsForTLSActivationInput{
                    TLSCertificate: &shared.RelationshipsForTLSActivationTLSCertificateInput{
                        Data: &shared.RelationshipMemberTLSCertificateInput{
                            Type: fastly.String("culpa"),
                        },
                    },
                    TLSDomain: &shared.RelationshipsForTLSActivationTLSDomainInput{
                        Data: &shared.RelationshipMemberTLSDomainInput{
                            Type: fastly.String("corrupti"),
                        },
                    },
                },
                Type: fastly.String("pariatur"),
            },
        },
        TLSActivationID: "aCtguUGZzb2W9Euo4moOR",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSActivationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateTLSActivationRequest](../../models/operations/updatetlsactivationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdateTLSActivationResponse](../../models/operations/updatetlsactivationresponse.md), error**

