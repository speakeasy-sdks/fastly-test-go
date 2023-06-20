# BillingAddress

## Overview

A billing address is used to calculate your bill correctly.

<https://developer.fastly.com/reference/api/account/billing-address>
### Available Operations

* [AddBillingAddr](#addbillingaddr) - Add a billing address to a customer
* [DeleteBillingAddr](#deletebillingaddr) - Delete a billing address
* [GetBillingAddr](#getbillingaddr) - Get a billing address
* [UpdateBillingAddr](#updatebillingaddr) - Update a billing address

## AddBillingAddr

Add a billing address to a customer.

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
    res, err := s.BillingAddress.AddBillingAddr(ctx, operations.AddBillingAddrRequest{
        BillingAddressRequestInput: &shared.BillingAddressRequestInput{
            Data: &shared.BillingAddressRequestDataInput{
                Attributes: &shared.BillingAddressAttributesInput{
                    Address1: sdk.String("80719 Dorothea Mountain"),
                    Address2: sdk.String("Apt. 652"),
                    City: sdk.String("New Rasheedville"),
                    Country: sdk.String("US"),
                    Locality: sdk.String("New Castle"),
                    PostalCode: sdk.String("53538-5902"),
                    State: sdk.String("DE"),
                },
                Type: shared.TypeBillingAddressBillingAddress.ToPointer(),
            },
            SkipVerification: sdk.Bool(false),
        },
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
    }, operations.AddBillingAddrSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillingAddressResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.AddBillingAddrRequest](../../models/operations/addbillingaddrrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.AddBillingAddrSecurity](../../models/operations/addbillingaddrsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.AddBillingAddrResponse](../../models/operations/addbillingaddrresponse.md), error**


## DeleteBillingAddr

Delete a customer's billing address.

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
    res, err := s.BillingAddress.DeleteBillingAddr(ctx, operations.DeleteBillingAddrRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
    }, operations.DeleteBillingAddrSecurity{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteBillingAddrRequest](../../models/operations/deletebillingaddrrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.DeleteBillingAddrSecurity](../../models/operations/deletebillingaddrsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.DeleteBillingAddrResponse](../../models/operations/deletebillingaddrresponse.md), error**


## GetBillingAddr

Get a customer's billing address.

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
    res, err := s.BillingAddress.GetBillingAddr(ctx, operations.GetBillingAddrRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
    }, operations.GetBillingAddrSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillingAddressResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetBillingAddrRequest](../../models/operations/getbillingaddrrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetBillingAddrSecurity](../../models/operations/getbillingaddrsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetBillingAddrResponse](../../models/operations/getbillingaddrresponse.md), error**


## UpdateBillingAddr

Update a customer's billing address. You may update only part of the customer's billing address.

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
    res, err := s.BillingAddress.UpdateBillingAddr(ctx, operations.UpdateBillingAddrRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
        UpdateBillingAddressRequestInput: &shared.UpdateBillingAddressRequestInput{
            Data: &shared.UpdateBillingAddressRequestDataInput{
                Attributes: &shared.BillingAddressAttributesInput{
                    Address1: sdk.String("80719 Dorothea Mountain"),
                    Address2: sdk.String("Apt. 652"),
                    City: sdk.String("New Rasheedville"),
                    Country: sdk.String("US"),
                    Locality: sdk.String("New Castle"),
                    PostalCode: sdk.String("53538-5902"),
                    State: sdk.String("DE"),
                },
                Type: shared.TypeBillingAddressBillingAddress.ToPointer(),
            },
            SkipVerification: sdk.Bool(false),
        },
    }, operations.UpdateBillingAddrSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillingAddressResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateBillingAddrRequest](../../models/operations/updatebillingaddrrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UpdateBillingAddrSecurity](../../models/operations/updatebillingaddrsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.UpdateBillingAddrResponse](../../models/operations/updatebillingaddrresponse.md), error**

