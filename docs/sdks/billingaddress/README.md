# BillingAddress
(*BillingAddress*)

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
    res, err := s.BillingAddress.AddBillingAddr(ctx, operations.AddBillingAddrRequest{
        BillingAddressRequest: &components.BillingAddressRequest{
            Data: &components.BillingAddressRequestData{
                Attributes: &components.BillingAddressAttributesInput{
                    Address1: fastlytestgo.String("80719 Dorothea Mountain"),
                    Address2: fastlytestgo.String("Apt. 652"),
                    City: fastlytestgo.String("New Rasheedville"),
                    Country: fastlytestgo.String("US"),
                    Locality: fastlytestgo.String("New Castle"),
                    PostalCode: fastlytestgo.String("53538-5902"),
                    State: fastlytestgo.String("DE"),
                },
            },
        },
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.AddBillingAddrRequest](../../models/operations/addbillingaddrrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.AddBillingAddrResponse](../../models/operations/addbillingaddrresponse.md), error**
| Error Object                                      | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| sdkerrors.BillingAddressVerificationErrorResponse | 400                                               | application/vnd.api+json                          |
| sdkerrors.SDKError                                | 4xx-5xx                                           | */*                                               |

## DeleteBillingAddr

Delete a customer's billing address.

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
    res, err := s.BillingAddress.DeleteBillingAddr(ctx, operations.DeleteBillingAddrRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
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
| `request`                                                                                  | [operations.DeleteBillingAddrRequest](../../models/operations/deletebillingaddrrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteBillingAddrResponse](../../models/operations/deletebillingaddrresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetBillingAddr

Get a customer's billing address.

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
    res, err := s.BillingAddress.GetBillingAddr(ctx, operations.GetBillingAddrRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetBillingAddrRequest](../../models/operations/getbillingaddrrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetBillingAddrResponse](../../models/operations/getbillingaddrresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateBillingAddr

Update a customer's billing address. You may update only part of the customer's billing address.

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
    res, err := s.BillingAddress.UpdateBillingAddr(ctx, operations.UpdateBillingAddrRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
        UpdateBillingAddressRequest: &components.UpdateBillingAddressRequest{
            Data: &components.UpdateBillingAddressRequestData{
                Attributes: &components.BillingAddressAttributesInput{
                    Address1: fastlytestgo.String("80719 Dorothea Mountain"),
                    Address2: fastlytestgo.String("Apt. 652"),
                    City: fastlytestgo.String("New Rasheedville"),
                    Country: fastlytestgo.String("US"),
                    Locality: fastlytestgo.String("New Castle"),
                    PostalCode: fastlytestgo.String("53538-5902"),
                    State: fastlytestgo.String("DE"),
                },
            },
        },
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateBillingAddrRequest](../../models/operations/updatebillingaddrrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateBillingAddrResponse](../../models/operations/updatebillingaddrresponse.md), error**
| Error Object                                      | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| sdkerrors.BillingAddressVerificationErrorResponse | 400                                               | application/vnd.api+json                          |
| sdkerrors.SDKError                                | 4xx-5xx                                           | */*                                               |
