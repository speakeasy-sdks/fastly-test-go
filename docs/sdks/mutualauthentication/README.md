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
    res, err := s.MutualAuthentication.CreateMutualTLSAuthentication(ctx, shared.MutualAuthentication{
        Data: &shared.MutualAuthenticationData{
            Attributes: &shared.MutualAuthenticationDataAttributes{
                CertBundle: fastly.String("Manat yahoo"),
                Enforced: fastly.Bool(false),
                Name: fastly.String("architect female Markets"),
            },
            Relationships: &shared.RelationshipsForMutualAuthenticationInput{},
            Type: fastly.String("Architect Recycled"),
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.MutualAuthentication](../../models/shared/mutualauthentication.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


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


## GetMutualAuthentication

Show a Mutual Authentication.

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
    res, err := s.MutualAuthentication.GetMutualAuthentication(ctx, operations.GetMutualAuthenticationRequest{
        Include: fastly.String("huzzah"),
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


## ListMutualAuthentications

List all mutual authentications.

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
    res, err := s.MutualAuthentication.ListMutualAuthentications(ctx, operations.ListMutualAuthenticationsRequest{
        Include: fastly.String("visionary"),
        PageNumber: fastly.Int64(1),
        PageSize: fastly.Int64(20),
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


## PatchMutualAuthentication

Update a Mutual Authentication.

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
    res, err := s.MutualAuthentication.PatchMutualAuthentication(ctx, operations.PatchMutualAuthenticationRequest{
        MutualAuthentication: &shared.MutualAuthentication{
            Data: &shared.MutualAuthenticationData{
                Attributes: &shared.MutualAuthenticationDataAttributes{
                    CertBundle: fastly.String("New"),
                    Enforced: fastly.Bool(false),
                    Name: fastly.String("Convertible Gasoline Vista"),
                },
                Relationships: &shared.RelationshipsForMutualAuthenticationInput{},
                Type: fastly.String("Applications abaft web"),
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

