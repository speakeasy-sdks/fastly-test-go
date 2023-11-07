# Tokens
(*.Tokens*)

## Overview

An API Token is used to identify who is making the API call. Users can create multiple tokens to suit their needs.

<https://developer.fastly.com/reference/api/auth-tokens/user>
### Available Operations

* [GetTokenCurrent](#gettokencurrent) - Get the current token
* [ListTokensCustomer](#listtokenscustomer) - List tokens for a customer
* [ListTokensUser](#listtokensuser) - List tokens for the authenticated user
* [RevokeToken](#revoketoken) - Revoke a token
* [RevokeTokenCurrent](#revoketokencurrent) - Revoke the current token

## GetTokenCurrent

Get a single token based on the access_token used in the request.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Tokens.GetTokenCurrent(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.TokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetTokenCurrentResponse](../../models/operations/gettokencurrentresponse.md), error**


## ListTokensCustomer

List all tokens belonging to a specific customer.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Tokens.ListTokensCustomer(ctx, operations.ListTokensCustomerRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListTokensCustomerRequest](../../models/operations/listtokenscustomerrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListTokensCustomerResponse](../../models/operations/listtokenscustomerresponse.md), error**


## ListTokensUser

List all tokens belonging to the authenticated user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Tokens.ListTokensUser(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListTokensUserResponse](../../models/operations/listtokensuserresponse.md), error**


## RevokeToken

Revoke a specific token by its id.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Tokens.RevokeToken(ctx, operations.RevokeTokenRequest{
        TokenID: "5Yo3XXnrQpjc20u0ybrf2g",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.RevokeTokenRequest](../../models/operations/revoketokenrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.RevokeTokenResponse](../../models/operations/revoketokenresponse.md), error**


## RevokeTokenCurrent

Revoke a token that is used to authenticate the request.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/models/components"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Tokens.RevokeTokenCurrent(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.RevokeTokenCurrentResponse](../../models/operations/revoketokencurrentresponse.md), error**

