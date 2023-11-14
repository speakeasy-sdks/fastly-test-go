# TLSSubscriptions
(*TLSSubscriptions*)

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
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.TLSSubscriptions.CreateGlobalsignEmailChallenge(ctx, operations.CreateGlobalsignEmailChallengeRequest{
        RequestBody: map[string]interface{}{
            "key": "string",
        },
        TLSAuthorizationID: "aU3guUGZzb2W9Euo4Mo0r",
        TLSSubscriptionID: "sU3guUGZzb2W9Euo4Mo0r",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.CreateGlobalsignEmailChallengeRequest](../../models/operations/createglobalsignemailchallengerequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.CreateGlobalsignEmailChallengeResponse](../../models/operations/createglobalsignemailchallengeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateTLSSub

Create a new TLS subscription. This response includes a list of possible challenges to verify domain ownership.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.TLSSubscriptions.CreateTLSSub(ctx, operations.CreateTLSSubRequest{
        Force: fastlytestgo.Bool(true),
        TLSSubscription: &components.TLSSubscription{
            Data: &components.TLSSubscriptionData{
                Attributes: &components.TLSSubscriptionDataAttributes{},
                Relationships: components.CreateRelationshipsForTLSSubscriptionRelationshipCommonName(
                        components.RelationshipCommonName{
                            CommonName: &components.RelationshipMemberTLSDomainInput{},
                        },
                ),
            },
        },
    })
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
| `request`                                                                        | [operations.CreateTLSSubRequest](../../models/operations/createtlssubrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateTLSSubResponse](../../models/operations/createtlssubresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteGlobalsignEmailChallenge

Deletes a GlobalSign email challenge. After a GlobalSign email challenge is deleted, the domain can use HTTP and DNS validation methods again.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.TLSSubscriptions.DeleteGlobalsignEmailChallenge(ctx, operations.DeleteGlobalsignEmailChallengeRequest{
        GlobalsignEmailChallengeID: "string",
        TLSAuthorizationID: "string",
        TLSSubscriptionID: "string",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DeleteGlobalsignEmailChallengeRequest](../../models/operations/deleteglobalsignemailchallengerequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.DeleteGlobalsignEmailChallengeResponse](../../models/operations/deleteglobalsignemailchallengeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteTLSSub

Destroy a TLS subscription. A subscription cannot be destroyed if there are domains in the TLS enabled state.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.TLSSubscriptions.DeleteTLSSub(ctx, operations.DeleteTLSSubRequest{
        TLSSubscriptionID: "sU3guUGZzb2W9Euo4Mo0r",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteTLSSubRequest](../../models/operations/deletetlssubrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DeleteTLSSubResponse](../../models/operations/deletetlssubresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetTLSSub

Show a TLS subscription.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.TLSSubscriptions.GetTLSSub(ctx, operations.GetTLSSubRequest{
        Include: fastlytestgo.String("tls_authorizations"),
        TLSSubscriptionID: "sU3guUGZzb2W9Euo4Mo0r",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetTLSSubRequest](../../models/operations/gettlssubrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetTLSSubResponse](../../models/operations/gettlssubresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListTLSSubs

List all TLS subscriptions.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.TLSSubscriptions.ListTLSSubs(ctx, operations.ListTLSSubsRequest{
        Include: fastlytestgo.String("tls_authorizations"),
        PageNumber: fastlytestgo.Int64(1),
        PageSize: fastlytestgo.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSSubscriptionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListTLSSubsRequest](../../models/operations/listtlssubsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ListTLSSubsResponse](../../models/operations/listtlssubsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PatchTLSSub

Change the TLS domains or common name associated with this subscription, update the TLS configuration for this set of domains, or retry a subscription with state `failed` by setting the state to `retry`.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.TLSSubscriptions.PatchTLSSub(ctx, operations.PatchTLSSubRequest{
        Force: fastlytestgo.Bool(true),
        TLSSubscription: &components.TLSSubscription{
            Data: &components.TLSSubscriptionData{
                Attributes: &components.TLSSubscriptionDataAttributes{},
                Relationships: components.CreateRelationshipsForTLSSubscriptionRelationshipTLSDomainsInput(
                        components.RelationshipTLSDomainsInput{
                            TLSDomains: &components.RelationshipTLSDomainsTLSDomains{
                                Data: []components.RelationshipMemberTLSDomainInput{
                                    components.RelationshipMemberTLSDomainInput{},
                                },
                            },
                        },
                ),
            },
        },
        TLSSubscriptionID: "sU3guUGZzb2W9Euo4Mo0r",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.PatchTLSSubRequest](../../models/operations/patchtlssubrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.PatchTLSSubResponse](../../models/operations/patchtlssubresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
