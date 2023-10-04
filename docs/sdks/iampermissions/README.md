# IamPermissions
(*IamPermissions*)

## Overview

A list of available permissions that can be assigned to a custom role.

<https://developer.fastly.com/reference/api/account/permissions>
### Available Operations

* [ListPermissions](#listpermissions) - List permissions

## ListPermissions

List all permissions.

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
    res, err := s.IamPermissions.ListPermissions(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPermissions200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListPermissionsResponse](../../models/operations/listpermissionsresponse.md), error**

