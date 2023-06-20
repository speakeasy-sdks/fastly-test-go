# Customer

## Overview

A Customer is the base object that owns your Users and Services. Some information may be limited depending on access level.

<https://developer.fastly.com/reference/api/account/customer>
### Available Operations

* [DeleteCustomer](#deletecustomer) - Delete a customer
* [GetCustomer](#getcustomer) - Get a customer
* [GetLoggedInCustomer](#getloggedincustomer) - Get the logged in customer
* [ListUsers](#listusers) - List users
* [UpdateCustomer](#updatecustomer) - Update a customer

## DeleteCustomer

Delete a customer.

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
    res, err := s.Customer.DeleteCustomer(ctx, operations.DeleteCustomerRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
    }, operations.DeleteCustomerSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteCustomer200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteCustomerRequest](../../models/operations/deletecustomerrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.DeleteCustomerSecurity](../../models/operations/deletecustomersecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.DeleteCustomerResponse](../../models/operations/deletecustomerresponse.md), error**


## GetCustomer

Get a specific customer.

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
    res, err := s.Customer.GetCustomer(ctx, operations.GetCustomerRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
    }, operations.GetCustomerSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetCustomerRequest](../../models/operations/getcustomerrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.GetCustomerSecurity](../../models/operations/getcustomersecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.GetCustomerResponse](../../models/operations/getcustomerresponse.md), error**


## GetLoggedInCustomer

Get the logged in customer.

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
    res, err := s.Customer.GetLoggedInCustomer(ctx, operations.GetLoggedInCustomerSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `security`                                                                                       | [operations.GetLoggedInCustomerSecurity](../../models/operations/getloggedincustomersecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.GetLoggedInCustomerResponse](../../models/operations/getloggedincustomerresponse.md), error**


## ListUsers

List all users from a specified customer id.

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
    res, err := s.Customer.ListUsers(ctx, operations.ListUsersRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
    }, operations.ListUsersSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SchemasUserResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListUsersRequest](../../models/operations/listusersrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.ListUsersSecurity](../../models/operations/listuserssecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.ListUsersResponse](../../models/operations/listusersresponse.md), error**


## UpdateCustomer

Update a customer.

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
    res, err := s.Customer.UpdateCustomer(ctx, operations.UpdateCustomerRequest{
        CustomerInput: &shared.CustomerInput{
            BillingContactID: sdk.String("enim"),
            BillingNetworkType: shared.CustomerBillingNetworkTypePublic.ToPointer(),
            BillingRef: sdk.String("est"),
            Force2fa: sdk.Bool(false),
            ForceSso: sdk.Bool(false),
            HasAccountPanel: sdk.Bool(false),
            HasImprovedEvents: sdk.Bool(false),
            HasOpenstackLogging: sdk.Bool(false),
            HasPci: sdk.Bool(false),
            IPWhitelist: sdk.String("quibusdam"),
            LegalContactID: sdk.String("explicabo"),
            Name: sdk.String("Rudy Spencer"),
            OwnerID: sdk.String("qui"),
            PhoneNumber: sdk.String("aliquid"),
            PostalAddress: sdk.String("cupiditate"),
            PricingPlan: sdk.String("quos"),
            PricingPlanID: sdk.String("perferendis"),
            SecurityContactID: sdk.String("magni"),
            TechnicalContactID: sdk.String("assumenda"),
        },
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
    }, operations.UpdateCustomerSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateCustomerRequest](../../models/operations/updatecustomerrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.UpdateCustomerSecurity](../../models/operations/updatecustomersecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.UpdateCustomerResponse](../../models/operations/updatecustomerresponse.md), error**

