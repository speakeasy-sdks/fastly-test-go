# TLSCertificates
(*TLSCertificates*)

## Overview

A TLS certificate is used to terminate TLS traffic for one or more of your [TLS domains](#tls_domains).

<https://developer.fastly.com/reference/api/tls/custom-certs/certificates>
### Available Operations

* [CreateTLSCert](#createtlscert) - Create a TLS certificate
* [DeleteTLSCert](#deletetlscert) - Delete a TLS certificate
* [GetTLSCert](#gettlscert) - Get a TLS certificate
* [ListTLSCerts](#listtlscerts) - List TLS certificates
* [UpdateTLSCert](#updatetlscert) - Update a TLS certificate

## CreateTLSCert

Create a TLS certificate.

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
    res, err := s.TLSCertificates.CreateTLSCert(ctx, shared.TLSCertificateInput{
        Data: &shared.TLSCertificateDataInput{
            Attributes: &shared.TLSCertificateDataAttributes{
                CertBlob: fastly.String("Gloves SQL Dubnium"),
                Name: fastly.String("Avon Electric"),
            },
            Relationships: &shared.RelationshipTLSDomainsInput{
                TLSDomains: &shared.RelationshipTLSDomainsTLSDomainsInput{
                    Data: []shared.RelationshipMemberTLSDomainInput{
                        shared.RelationshipMemberTLSDomainInput{
                            Type: shared.TypeTLSDomainTLSDomain.ToPointer(),
                        },
                    },
                },
            },
            Type: shared.TypeTLSCertificateTLSCertificate.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTLSCert201ApplicationVndAPIPlusJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.TLSCertificateInput](../../models/shared/tlscertificateinput.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CreateTLSCertResponse](../../models/operations/createtlscertresponse.md), error**


## DeleteTLSCert

Destroy a TLS certificate. TLS certificates already enabled for a domain cannot be destroyed.

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
    res, err := s.TLSCertificates.DeleteTLSCert(ctx, operations.DeleteTLSCertRequest{
        TLSCertificateID: "cRTguUGZzb2W9Euo4moOr",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteTLSCertRequest](../../models/operations/deletetlscertrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.DeleteTLSCertResponse](../../models/operations/deletetlscertresponse.md), error**


## GetTLSCert

Show a TLS certificate.

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
    res, err := s.TLSCertificates.GetTLSCert(ctx, operations.GetTLSCertRequest{
        TLSCertificateID: "cRTguUGZzb2W9Euo4moOr",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSCertificateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetTLSCertRequest](../../models/operations/gettlscertrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetTLSCertResponse](../../models/operations/gettlscertresponse.md), error**


## ListTLSCerts

List all TLS certificates.

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
    res, err := s.TLSCertificates.ListTLSCerts(ctx, operations.ListTLSCertsRequest{
        FilterInUse: fastly.String("Northwest overriding plum"),
        FilterNotAfter: fastly.String("Dollar"),
        FilterTLSDomainsID: fastly.String("Corporate Customer"),
        Include: fastly.String("Passage Blues"),
        PageNumber: fastly.Int64(1),
        PageSize: fastly.Int64(20),
        Sort: shared.SortCreatedAt.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSCertificatesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListTLSCertsRequest](../../models/operations/listtlscertsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListTLSCertsResponse](../../models/operations/listtlscertsresponse.md), error**


## UpdateTLSCert

Replace a TLS certificate with a newly reissued TLS certificate, or update a TLS certificate's name. If replacing a TLS certificate, the new TLS certificate must contain all SAN entries as the current TLS certificate. It must either have an exact matching list or contain a superset.

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
    res, err := s.TLSCertificates.UpdateTLSCert(ctx, operations.UpdateTLSCertRequest{
        TLSCertificateInput: &shared.TLSCertificateInput{
            Data: &shared.TLSCertificateDataInput{
                Attributes: &shared.TLSCertificateDataAttributes{
                    CertBlob: fastly.String("male Klocko North"),
                    Name: fastly.String("National"),
                },
                Relationships: &shared.RelationshipTLSDomainsInput{
                    TLSDomains: &shared.RelationshipTLSDomainsTLSDomainsInput{
                        Data: []shared.RelationshipMemberTLSDomainInput{
                            shared.RelationshipMemberTLSDomainInput{
                                Type: shared.TypeTLSDomainTLSDomain.ToPointer(),
                            },
                        },
                    },
                },
                Type: shared.TypeTLSCertificateTLSCertificate.ToPointer(),
            },
        },
        TLSCertificateID: "cRTguUGZzb2W9Euo4moOr",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSCertificateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateTLSCertRequest](../../models/operations/updatetlscertrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateTLSCertResponse](../../models/operations/updatetlscertresponse.md), error**

