# Billing

## Overview

Get information on current and past bills.

<https://developer.fastly.com/reference/api/account/billing>
### Available Operations

* [GetInvoice](#getinvoice) - Get an invoice
* [GetInvoiceByID](#getinvoicebyid) - Get an invoice
* [GetInvoiceMtd](#getinvoicemtd) - Get month-to-date billing estimate

## GetInvoice

Get the invoice for a given year and month. Can be any month from when the Customer was created to the current month.

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
    res, err := s.Billing.GetInvoice(ctx, operations.GetInvoiceRequest{
        Month: "05",
        Year: "2020",
    }, operations.GetInvoiceSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetInvoiceRequest](../../models/operations/getinvoicerequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.GetInvoiceSecurity](../../models/operations/getinvoicesecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.GetInvoiceResponse](../../models/operations/getinvoiceresponse.md), error**


## GetInvoiceByID

Get the invoice for the given invoice_id.

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
    res, err := s.Billing.GetInvoiceByID(ctx, operations.GetInvoiceByIDRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
        InvoiceID: "7SlAESxcJ2zxHOV4gQ9y9X",
    }, operations.GetInvoiceByIDSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetInvoiceByIDRequest](../../models/operations/getinvoicebyidrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetInvoiceByIDSecurity](../../models/operations/getinvoicebyidsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetInvoiceByIDResponse](../../models/operations/getinvoicebyidresponse.md), error**


## GetInvoiceMtd

Get the current month-to-date estimate. This endpoint has two different responses. Under normal circumstances, it generally takes less than 5 seconds to generate but in certain cases can take up to 60 seconds. Once generated the month-to-date estimate is cached for 4 hours, and is available the next request will return the JSON representation of the month-to-date estimate. While a report is being generated in the background, this endpoint will return a `202 Accepted` response. The full format of which can be found in detail in our [billing calculation guide](https://docs.fastly.com/en/guides/how-we-calculate-your-bill). There are certain accounts for which we are unable to generate a month-to-date estimate. For example, accounts who have parent-pay are unable to generate an MTD estimate. The parent accounts are able to generate a month-to-date estimate but that estimate will not include the child accounts amounts at this time.

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
    res, err := s.Billing.GetInvoiceMtd(ctx, operations.GetInvoiceMtdRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
        Month: sdk.String("05"),
        Year: sdk.String("2020"),
    }, operations.GetInvoiceMtdSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillingEstimateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetInvoiceMtdRequest](../../models/operations/getinvoicemtdrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.GetInvoiceMtdSecurity](../../models/operations/getinvoicemtdsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.GetInvoiceMtdResponse](../../models/operations/getinvoicemtdresponse.md), error**

