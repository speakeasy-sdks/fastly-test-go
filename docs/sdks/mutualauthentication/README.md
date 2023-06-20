# MutualAuthentication

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
	"context"
	"log"
	"Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.MutualAuthentication.CreateMutualTLSAuthentication(ctx, shared.MutualAuthenticationInput{
        Data: &shared.MutualAuthenticationDataInput{
            Attributes: &shared.MutualAuthenticationDataAttributes{
                CertBundle: sdk.String("in"),
                Enforced: sdk.Bool(false),
                Name: sdk.String("Suzanne Torphy"),
            },
            Relationships: &shared.RelationshipTLSActivationsInput{
                TLSActivations: &shared.RelationshipTLSActivationsTLSActivationsInput{
                    Data: []shared.RelationshipMemberTLSActivationInput{
                        shared.RelationshipMemberTLSActivationInput{
                            Type: shared.TypeTLSActivationTLSActivation.ToPointer(),
                        },
                    },
                },
            },
            Type: shared.TypeMutualAuthenticationMutualAuthentication.ToPointer(),
        },
    }, operations.CreateMutualTLSAuthenticationSecurity{
        Token: "",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [shared.MutualAuthenticationInput](../../models/shared/mutualauthenticationinput.md)                                 | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.CreateMutualTLSAuthenticationSecurity](../../models/operations/createmutualtlsauthenticationsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.CreateMutualTLSAuthenticationResponse](../../models/operations/createmutualtlsauthenticationresponse.md), error**


## DeleteMutualTLS

Remove a Mutual TLS authentication

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
    res, err := s.MutualAuthentication.DeleteMutualTLS(ctx, operations.DeleteMutualTLSRequest{
        MutualAuthenticationID: "SEAwSOsP7dEpTgGZdP7ZFw",
    }, operations.DeleteMutualTLSSecurity{
        Token: "",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteMutualTLSRequest](../../models/operations/deletemutualtlsrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.DeleteMutualTLSSecurity](../../models/operations/deletemutualtlssecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.DeleteMutualTLSResponse](../../models/operations/deletemutualtlsresponse.md), error**


## GetMutualAuthentication

Show a Mutual Authentication.

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
    res, err := s.MutualAuthentication.GetMutualAuthentication(ctx, operations.GetMutualAuthenticationRequest{
        Include: sdk.String("cum"),
        MutualAuthenticationID: "SEAwSOsP7dEpTgGZdP7ZFw",
    }, operations.GetMutualAuthenticationSecurity{
        Token: "",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetMutualAuthenticationRequest](../../models/operations/getmutualauthenticationrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.GetMutualAuthenticationSecurity](../../models/operations/getmutualauthenticationsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.GetMutualAuthenticationResponse](../../models/operations/getmutualauthenticationresponse.md), error**


## ListMutualAuthentications

List all mutual authentications.

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
    res, err := s.MutualAuthentication.ListMutualAuthentications(ctx, operations.ListMutualAuthenticationsRequest{
        Include: sdk.String("blanditiis"),
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
    }, operations.ListMutualAuthenticationsSecurity{
        Token: "",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListMutualAuthenticationsRequest](../../models/operations/listmutualauthenticationsrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.ListMutualAuthenticationsSecurity](../../models/operations/listmutualauthenticationssecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.ListMutualAuthenticationsResponse](../../models/operations/listmutualauthenticationsresponse.md), error**


## PatchMutualAuthentication

Update a Mutual Authentication.

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
    res, err := s.MutualAuthentication.PatchMutualAuthentication(ctx, operations.PatchMutualAuthenticationRequest{
        MutualAuthenticationInput: &shared.MutualAuthenticationInput{
            Data: &shared.MutualAuthenticationDataInput{
                Attributes: &shared.MutualAuthenticationDataAttributes{
                    CertBundle: sdk.String("quas"),
                    Enforced: sdk.Bool(false),
                    Name: sdk.String("Danny Nader"),
                },
                Relationships: &shared.RelationshipTLSActivationsInput{
                    TLSActivations: &shared.RelationshipTLSActivationsTLSActivationsInput{
                        Data: []shared.RelationshipMemberTLSActivationInput{
                            shared.RelationshipMemberTLSActivationInput{
                                Type: shared.TypeTLSActivationTLSActivation.ToPointer(),
                            },
                            shared.RelationshipMemberTLSActivationInput{
                                Type: shared.TypeTLSActivationTLSActivation.ToPointer(),
                            },
                            shared.RelationshipMemberTLSActivationInput{
                                Type: shared.TypeTLSActivationTLSActivation.ToPointer(),
                            },
                        },
                    },
                },
                Type: shared.TypeMutualAuthenticationMutualAuthentication.ToPointer(),
            },
        },
        MutualAuthenticationID: "SEAwSOsP7dEpTgGZdP7ZFw",
    }, operations.PatchMutualAuthenticationSecurity{
        Token: "",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchMutualAuthenticationRequest](../../models/operations/patchmutualauthenticationrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.PatchMutualAuthenticationSecurity](../../models/operations/patchmutualauthenticationsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.PatchMutualAuthenticationResponse](../../models/operations/patchmutualauthenticationresponse.md), error**

