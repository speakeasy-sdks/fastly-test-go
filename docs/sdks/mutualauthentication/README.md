# MutualAuthentication
(*MutualAuthentication*)

## Overview

The Mutual TLS API allows for client-to-server authentication using client-side X.509 authentication. The main Mutual Authentication object represents the certificate bundle and other configurations which support Mutual TLS for your domains.

<https://developer.fastly.com/reference/api/tls/mutual-tls/authentication>
### Available Operations

* [CreateMutualTLSAuthentication](#createmutualtlsauthentication) - Create a Mutual Authentication
* [DeleteMutualTLS](#deletemutualtls) - Delete a Mutual TLS
* [GetMutualAuthentication](#getmutualauthentication) - Get a Mutual Authentication
* [ListMutualAuthentications](#listmutualauthentications) - List Mutual Authentications
* [PatchMutualAuthentication](#patchmutualauthentication) - Update a Mutual Authentication

## CreateMutualTLSAuthentication

Create a mutual authentication using a bundle of certificates to enable client-to-server mutual TLS.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.MutualAuthentication.CreateMutualTLSAuthentication(ctx, &components.MutualAuthentication{
        Data: &components.MutualAuthenticationData{
            Attributes: &components.MutualAuthenticationDataAttributes{},
            Relationships: components.CreateRelationshipsForMutualAuthenticationInputRelationshipTLSActivationsInput(
                    components.RelationshipTLSActivationsInput{
                        TLSActivations: &components.RelationshipTLSActivationsTLSActivations{
                            Data: []components.RelationshipMemberTLSActivationInput{
                                components.RelationshipMemberTLSActivationInput{},
                            },
                        },
                    },
            ),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MutualAuthenticationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.MutualAuthentication](../../models/components/mutualauthentication.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateMutualTLSAuthenticationResponse](../../models/operations/createmutualtlsauthenticationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteMutualTLS

Remove a Mutual TLS authentication

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
	"net/http"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.MutualAuthentication.DeleteMutualTLS(ctx, operations.DeleteMutualTLSRequest{
        MutualAuthenticationID: "SEAwSOsP7dEpTgGZdP7ZFw",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteMutualTLSRequest](../../models/operations/deletemutualtlsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DeleteMutualTLSResponse](../../models/operations/deletemutualtlsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetMutualAuthentication

Show a Mutual Authentication.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.MutualAuthentication.GetMutualAuthentication(ctx, operations.GetMutualAuthenticationRequest{
        MutualAuthenticationID: "SEAwSOsP7dEpTgGZdP7ZFw",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MutualAuthenticationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetMutualAuthenticationRequest](../../models/operations/getmutualauthenticationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetMutualAuthenticationResponse](../../models/operations/getmutualauthenticationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListMutualAuthentications

List all mutual authentications.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.MutualAuthentication.ListMutualAuthentications(ctx, operations.ListMutualAuthenticationsRequest{
        PageNumber: fastlytestgo.Int64(1),
        PageSize: fastlytestgo.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MutualAuthenticationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListMutualAuthenticationsRequest](../../models/operations/listmutualauthenticationsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.ListMutualAuthenticationsResponse](../../models/operations/listmutualauthenticationsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchMutualAuthentication

Update a Mutual Authentication.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.MutualAuthentication.PatchMutualAuthentication(ctx, operations.PatchMutualAuthenticationRequest{
        MutualAuthentication: &components.MutualAuthentication{
            Data: &components.MutualAuthenticationData{
                Attributes: &components.MutualAuthenticationDataAttributes{},
                Relationships: components.CreateRelationshipsForMutualAuthenticationInputRelationshipTLSActivationsInput(
                        components.RelationshipTLSActivationsInput{
                            TLSActivations: &components.RelationshipTLSActivationsTLSActivations{
                                Data: []components.RelationshipMemberTLSActivationInput{
                                    components.RelationshipMemberTLSActivationInput{},
                                },
                            },
                        },
                ),
            },
        },
        MutualAuthenticationID: "SEAwSOsP7dEpTgGZdP7ZFw",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MutualAuthenticationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PatchMutualAuthenticationRequest](../../models/operations/patchmutualauthenticationrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.PatchMutualAuthenticationResponse](../../models/operations/patchmutualauthenticationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
