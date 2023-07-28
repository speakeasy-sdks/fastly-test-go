# IamPermissions

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
	"Fastly"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()
    operationSecurity := operations.ListPermissionsSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.IamPermissions.ListPermissions(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPermissions200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `security`                                                                               | [operations.ListPermissionsSecurity](../../models/operations/listpermissionssecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ListPermissionsResponse](../../models/operations/listpermissionsresponse.md), error**

