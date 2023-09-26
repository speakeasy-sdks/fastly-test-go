# TLSBulkCertificates

## Overview

Available to Platform TLS customers, these endpoints streamline the upload, deployment and management of large numbers of TLS certificates. A certificate is used to terminate TLS traffic for one or more of your fully qualified domain names (domains). Uploading a new certificate automatically enables TLS for all domains listed as Subject Alternative Names (SAN entries) on the certificate.

<https://developer.fastly.com/reference/api/tls/platform>
### Available Operations

* [DeleteBulkTLSCert](#deletebulktlscert) - Delete a certificate
* [GetTLSBulkCert](#gettlsbulkcert) - Get a certificate
* [ListTLSBulkCerts](#listtlsbulkcerts) - List certificates
* [UpdateBulkTLSCert](#updatebulktlscert) - Update a certificate
* [UploadTLSBulkCert](#uploadtlsbulkcert) - Upload a certificate

## DeleteBulkTLSCert

Destroy a certificate. This disables TLS for all domains listed as SAN entries.

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
    res, err := s.TLSBulkCertificates.DeleteBulkTLSCert(ctx, operations.DeleteBulkTLSCertRequest{
        CertificateID: "cRTguUGZzb2W9Euo4moOr",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteBulkTLSCertRequest](../../models/operations/deletebulktlscertrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteBulkTLSCertResponse](../../models/operations/deletebulktlscertresponse.md), error**


## GetTLSBulkCert

Retrieve a single certificate.

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
    res, err := s.TLSBulkCertificates.GetTLSBulkCert(ctx, operations.GetTLSBulkCertRequest{
        CertificateID: "cRTguUGZzb2W9Euo4moOr",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSBulkCertificateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetTLSBulkCertRequest](../../models/operations/gettlsbulkcertrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetTLSBulkCertResponse](../../models/operations/gettlsbulkcertresponse.md), error**


## ListTLSBulkCerts

List all certificates.

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
    res, err := s.TLSBulkCertificates.ListTLSBulkCerts(ctx, operations.ListTLSBulkCertsRequest{
        FilterTLSDomainID: fastly.String("totam"),
        PageNumber: fastly.Int64(1),
        PageSize: fastly.Int64(20),
        Sort: shared.SortMinusCreatedAt.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSBulkCertificatesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListTLSBulkCertsRequest](../../models/operations/listtlsbulkcertsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListTLSBulkCertsResponse](../../models/operations/listtlsbulkcertsresponse.md), error**


## UpdateBulkTLSCert

Replace a certificate with a newly reissued certificate. By using this endpoint, the original certificate will cease to be used for future TLS handshakes. Thus, only SAN entries that appear in the replacement certificate will become TLS enabled. Any SAN entries that are missing in the replacement certificate will become disabled.

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
    res, err := s.TLSBulkCertificates.UpdateBulkTLSCert(ctx, operations.UpdateBulkTLSCertRequest{
        CertificateID: "cRTguUGZzb2W9Euo4moOr",
        TLSBulkCertificateInput: &shared.TLSBulkCertificateInput{
            Data: &shared.TLSBulkCertificateDataInput{
                Attributes: &shared.TLSBulkCertificateDataAttributes{
                    AllowUntrustedRoot: fastly.Bool(false),
                    CertBlob: fastly.String("exercitationem"),
                    IntermediatesBlob: fastly.String("nobis"),
                },
                Relationships: &shared.RelationshipTLSConfigurationsInput{
                    TLSConfigurations: &shared.RelationshipTLSConfigurationsTLSConfigurationsInput{
                        Data: []shared.RelationshipMemberTLSConfigurationInput{
                            shared.RelationshipMemberTLSConfigurationInput{
                                Type: fastly.String("sit"),
                            },
                        },
                    },
                },
                Type: fastly.String("rerum"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSBulkCertificateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateBulkTLSCertRequest](../../models/operations/updatebulktlscertrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateBulkTLSCertResponse](../../models/operations/updatebulktlscertresponse.md), error**


## UploadTLSBulkCert

Upload a new certificate. TLS domains are automatically enabled upon certificate creation. If a domain is already enabled on a previously uploaded certificate, that domain will be updated to use the new certificate for all future TLS handshake requests.

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
    res, err := s.TLSBulkCertificates.UploadTLSBulkCert(ctx, shared.TLSBulkCertificateDataInput{
        Attributes: &shared.TLSBulkCertificateDataAttributes{
            AllowUntrustedRoot: fastly.Bool(false),
            CertBlob: fastly.String("sed"),
            IntermediatesBlob: fastly.String("reiciendis"),
        },
        Relationships: &shared.RelationshipTLSConfigurationsInput{
            TLSConfigurations: &shared.RelationshipTLSConfigurationsTLSConfigurationsInput{
                Data: []shared.RelationshipMemberTLSConfigurationInput{
                    shared.RelationshipMemberTLSConfigurationInput{
                        Type: fastly.String("explicabo"),
                    },
                },
            },
        },
        Type: fastly.String("asperiores"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSBulkCertificateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.TLSBulkCertificateDataInput](../../models/shared/tlsbulkcertificatedatainput.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UploadTLSBulkCertResponse](../../models/operations/uploadtlsbulkcertresponse.md), error**

