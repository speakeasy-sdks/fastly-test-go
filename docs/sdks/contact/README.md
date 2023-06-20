# Contact

## Overview

A Customer Contact is the base object that holds the different types of contact information Fastly uses to contact a customer.

<https://developer.fastly.com/reference/api/account/contact>
### Available Operations

* [DeleteContact](#deletecontact) - Delete a contact
* [ListContacts](#listcontacts) - List contacts

## DeleteContact

Delete a contact.

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
    res, err := s.Contact.DeleteContact(ctx, operations.DeleteContactRequest{
        ContactID: "x4xCwxxJxGCx123Rx5xTx",
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
    }, operations.DeleteContactSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteContact200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteContactRequest](../../models/operations/deletecontactrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.DeleteContactSecurity](../../models/operations/deletecontactsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.DeleteContactResponse](../../models/operations/deletecontactresponse.md), error**


## ListContacts

List all contacts from a specified customer ID.

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
    res, err := s.Contact.ListContacts(ctx, operations.ListContactsRequest{
        CustomerID: "x4xCwxxJxGCx123Rx5xTx",
    }, operations.ListContactsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SchemasContactResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListContactsRequest](../../models/operations/listcontactsrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.ListContactsSecurity](../../models/operations/listcontactssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.ListContactsResponse](../../models/operations/listcontactsresponse.md), error**

