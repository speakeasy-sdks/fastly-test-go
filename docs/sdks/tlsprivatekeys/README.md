# TLSPrivateKeys

## Overview

A private key is used to sign a Certificate. A key can be used to sign multiple certificates.

<https://developer.fastly.com/reference/api/tls/custom-certs/private-keys>
### Available Operations

* [CreateTLSKey](#createtlskey) - Create a TLS private key
* [DeleteTLSKey](#deletetlskey) - Delete a TLS private key
* [GetTLSKey](#gettlskey) - Get a TLS private key
* [ListTLSKeys](#listtlskeys) - List TLS private keys

## CreateTLSKey

Create a TLS private key.

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
    res, err := s.TLSPrivateKeys.CreateTLSKey(ctx, shared.TLSPrivateKeyInput{
        Data: &shared.TLSPrivateKeyDataInput{
            Attributes: &shared.TLSPrivateKeyDataAttributes{
                Key: sdk.String("earum"),
                Name: sdk.String("Dr. Terrell Stanton"),
            },
            Relationships: &shared.TLSPrivateKeyDataRelationships2Input{
                TLSDomains: &shared.TLSPrivateKeyDataRelationships2TLSDomainsInput{
                    Data: []shared.RelationshipMemberTLSDomainInput{
                        shared.RelationshipMemberTLSDomainInput{
                            Type: shared.TypeTLSDomainTLSDomain.ToPointer(),
                        },
                        shared.RelationshipMemberTLSDomainInput{
                            Type: shared.TypeTLSDomainTLSDomain.ToPointer(),
                        },
                        shared.RelationshipMemberTLSDomainInput{
                            Type: shared.TypeTLSDomainTLSDomain.ToPointer(),
                        },
                        shared.RelationshipMemberTLSDomainInput{
                            Type: shared.TypeTLSDomainTLSDomain.ToPointer(),
                        },
                    },
                },
            },
            Type: shared.TypeTLSPrivateKeyTLSPrivateKey.ToPointer(),
        },
    }, operations.CreateTLSKeySecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSPrivateKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.TLSPrivateKeyInput](../../models/shared/tlsprivatekeyinput.md)             | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.CreateTLSKeySecurity](../../models/operations/createtlskeysecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.CreateTLSKeyResponse](../../models/operations/createtlskeyresponse.md), error**


## DeleteTLSKey

Destroy a TLS private key. Only private keys not already matched to any certificates can be deleted.

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
    res, err := s.TLSPrivateKeys.DeleteTLSKey(ctx, operations.DeleteTLSKeyRequest{
        TLSPrivateKeyID: "KeYguUGZzb2W9Euo4moOR",
    }, operations.DeleteTLSKeySecurity{
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteTLSKeyRequest](../../models/operations/deletetlskeyrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.DeleteTLSKeySecurity](../../models/operations/deletetlskeysecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.DeleteTLSKeyResponse](../../models/operations/deletetlskeyresponse.md), error**


## GetTLSKey

Show a TLS private key.

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
    res, err := s.TLSPrivateKeys.GetTLSKey(ctx, operations.GetTLSKeyRequest{
        TLSPrivateKeyID: "KeYguUGZzb2W9Euo4moOR",
    }, operations.GetTLSKeySecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSPrivateKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetTLSKeyRequest](../../models/operations/gettlskeyrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.GetTLSKeySecurity](../../models/operations/gettlskeysecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.GetTLSKeyResponse](../../models/operations/gettlskeyresponse.md), error**


## ListTLSKeys

List all TLS private keys.

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
    res, err := s.TLSPrivateKeys.ListTLSKeys(ctx, operations.ListTLSKeysRequest{
        FilterInUse: sdk.String("aliquid"),
        PageNumber: sdk.Int64(1),
        PageSize: sdk.Int64(20),
    }, operations.ListTLSKeysSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TLSPrivateKeysResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListTLSKeysRequest](../../models/operations/listtlskeysrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.ListTLSKeysSecurity](../../models/operations/listtlskeyssecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.ListTLSKeysResponse](../../models/operations/listtlskeysresponse.md), error**

