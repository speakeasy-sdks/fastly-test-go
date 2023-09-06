# TLSSubscriptions

## Overview

The TLS subscriptions API allows you to programmatically generate TLS certificates that are procured and renewed by Fastly. Once a subscription is created for a given hostname or wildcard domain, DNS records are checked to ensure that the domain on the subscription is owned by the subscription creator. Provided DNS records are maintained, TLS certificates will automatically renew. If Fastly is unable to issue a certificate, we will retry to issue the certificate for 7 days past subscription creation or the latest certificate's not_after date, whichever is later. If after 7 days Fastly is unable to issue a certificate, the subscription state will change to `failed` and Fastly will stop retrying.

<https://developer.fastly.com/reference/api/tls/subs>
### Available Operations

* [CreateGlobalsignEmailChallenge](#createglobalsignemailchallenge) - Creates a GlobalSign email challenge.
* [CreateTLSSub](#createtlssub) - Create a TLS subscription
* [DeleteGlobalsignEmailChallenge](#deleteglobalsignemailchallenge) - Delete a GlobalSign email challenge
* [DeleteTLSSub](#deletetlssub) - Delete a TLS subscription
* [GetTLSSub](#gettlssub) - Get a TLS subscription
* [ListTLSSubs](#listtlssubs) - List TLS subscriptions
* [PatchTLSSub](#patchtlssub) - Update a TLS subscription

## CreateGlobalsignEmailChallenge

Creates an email challenge for a domain on a GlobalSign subscription. An email challenge will generate an email that can be used to validate domain ownership. If this challenge is created, then the domain can only be validated using email for the given subscription.


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
    operationSecurity := operations.CreateGlobalsignEmailChallengeSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSSubscriptions.CreateGlobalsignEmailChallenge(ctx, operations.CreateGlobalsignEmailChallengeRequest{
        RequestBody: map[string]interface{}{
            "in": "commodi",
        },
        TLSAuthorizationID: "aU3guUGZzb2W9Euo4Mo0r",
        TLSSubscriptionID: "sU3guUGZzb2W9Euo4Mo0r",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateGlobalsignEmailChallenge201ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.CreateGlobalsignEmailChallengeRequest](../../models/operations/createglobalsignemailchallengerequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.CreateGlobalsignEmailChallengeSecurity](../../models/operations/createglobalsignemailchallengesecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


### Response

**[*operations.CreateGlobalsignEmailChallengeResponse](../../models/operations/createglobalsignemailchallengeresponse.md), error**


## CreateTLSSub

Create a new TLS subscription. This response includes a list of possible challenges to verify domain ownership.

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
    operationSecurity := operations.CreateTLSSubSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSSubscriptions.CreateTLSSub(ctx, operations.CreateTLSSubRequest{
        Force: sdk.Bool(true),
        TLSSubscriptionInput: &shared.TLSSubscriptionInput{
            Data: &shared.TLSSubscriptionDataInput{
                Attributes: &shared.TLSSubscriptionDataAttributes{
                    CertificateAuthority: shared.TLSSubscriptionDataAttributesCertificateAuthorityGlobalsign.ToPointer(),
                },
                Relationships: &shared.TLSSubscriptionDataRelationships1Input{
                    TLSConfiguration: &shared.TLSSubscriptionDataRelationships1TLSConfigurationInput{
                        Data: &shared.RelationshipMemberTLSConfigurationInput{
                            Type: shared.TypeTLSConfigurationTLSConfiguration.ToPointer(),
                        },
                    },
                },
                Type: shared.TypeTLSSubscriptionTLSSubscription.ToPointer(),
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateTLSSubRequest](../../models/operations/createtlssubrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.CreateTLSSubSecurity](../../models/operations/createtlssubsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.CreateTLSSubResponse](../../models/operations/createtlssubresponse.md), error**


## DeleteGlobalsignEmailChallenge

Deletes a GlobalSign email challenge. After a GlobalSign email challenge is deleted, the domain can use HTTP and DNS validation methods again.

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
    operationSecurity := operations.DeleteGlobalsignEmailChallengeSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSSubscriptions.DeleteGlobalsignEmailChallenge(ctx, operations.DeleteGlobalsignEmailChallengeRequest{
        GlobalsignEmailChallengeID: "voluptas",
        TLSAuthorizationID: "unde",
        TLSSubscriptionID: "architecto",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.DeleteGlobalsignEmailChallengeRequest](../../models/operations/deleteglobalsignemailchallengerequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.DeleteGlobalsignEmailChallengeSecurity](../../models/operations/deleteglobalsignemailchallengesecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


### Response

**[*operations.DeleteGlobalsignEmailChallengeResponse](../../models/operations/deleteglobalsignemailchallengeresponse.md), error**


## DeleteTLSSub

Destroy a TLS subscription. A subscription cannot be destroyed if there are domains in the TLS enabled state.

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
    operationSecurity := operations.DeleteTLSSubSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSSubscriptions.DeleteTLSSub(ctx, operations.DeleteTLSSubRequest{
        TLSSubscriptionID: "sU3guUGZzb2W9Euo4Mo0r",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteTLSSubRequest](../../models/operations/deletetlssubrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.DeleteTLSSubSecurity](../../models/operations/deletetlssubsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.DeleteTLSSubResponse](../../models/operations/deletetlssubresponse.md), error**


## GetTLSSub

Show a TLS subscription.

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
    operationSecurity := operations.GetTLSSubSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSSubscriptions.GetTLSSub(ctx, operations.GetTLSSubRequest{
        Include: sdk.String("tls_authorizations"),
        TLSSubscriptionID: "sU3guUGZzb2W9Euo4Mo0r",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetTLSSubRequest](../../models/operations/gettlssubrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.GetTLSSubSecurity](../../models/operations/gettlssubsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.GetTLSSubResponse](../../models/operations/gettlssubresponse.md), error**


## ListTLSSubs

List all TLS subscriptions.

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
    operationSecurity := operations.ListTLSSubsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSSubscriptions.ListTLSSubs(ctx, operations.ListTLSSubsRequest{
        FilterHasActiveOrder: sdk.Bool(false),
        FilterState: sdk.String("suscipit"),
        FilterTLSDomainsID: sdk.String("sapiente"),
        Include: sdk.String("tls_authorizations"),
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
        Sort: shared.SortMinusCreatedAt.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSSubscriptionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListTLSSubsRequest](../../models/operations/listtlssubsrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.ListTLSSubsSecurity](../../models/operations/listtlssubssecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.ListTLSSubsResponse](../../models/operations/listtlssubsresponse.md), error**


## PatchTLSSub

Change the TLS domains or common name associated with this subscription, update the TLS configuration for this set of domains, or retry a subscription with state `failed` by setting the state to `retry`.

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
    operationSecurity := operations.PatchTLSSubSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.TLSSubscriptions.PatchTLSSub(ctx, operations.PatchTLSSubRequest{
        Force: sdk.Bool(true),
        TLSSubscriptionInput: &shared.TLSSubscriptionInput{
            Data: &shared.TLSSubscriptionDataInput{
                Attributes: &shared.TLSSubscriptionDataAttributes{
                    CertificateAuthority: shared.TLSSubscriptionDataAttributesCertificateAuthorityLetsEncrypt.ToPointer(),
                },
                Relationships: &shared.RelationshipTLSCertificatesInput{
                    TLSCertificates: &shared.RelationshipTLSCertificatesTLSCertificatesInput{
                        Data: []shared.RelationshipMemberTLSCertificateInput{
                            shared.RelationshipMemberTLSCertificateInput{
                                Type: shared.TypeTLSCertificateTLSCertificate.ToPointer(),
                            },
                        },
                    },
                },
                Type: shared.TypeTLSSubscriptionTLSSubscription.ToPointer(),
            },
        },
        TLSSubscriptionID: "sU3guUGZzb2W9Euo4Mo0r",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PatchTLSSubRequest](../../models/operations/patchtlssubrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.PatchTLSSubSecurity](../../models/operations/patchtlssubsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.PatchTLSSubResponse](../../models/operations/patchtlssubresponse.md), error**

