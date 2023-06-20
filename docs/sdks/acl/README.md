# ACL

## Overview

An access control list or "ACL" specifies individual IP addresses or subnet ranges and can be accessed and used from Fastly VCL.

<https://developer.fastly.com/reference/api/acls/acl>
### Available Operations

* [CreateACL](#createacl) - Create a new ACL
* [DeleteACL](#deleteacl) - Delete an ACL
* [GetACL](#getacl) - Describe an ACL
* [ListAcls](#listacls) - List ACLs
* [UpdateACL](#updateacl) - Update an ACL

## CreateACL

Create a new ACL attached to the specified service version. A new, empty ACL must be attached to a draft version of a service. The version associated with the ACL must be activated to be used.

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
    res, err := s.ACL.CreateACL(ctx, operations.CreateACLRequest{
        ACL: &shared.ACL{
            Name: sdk.String("test-acl"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateACLSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ACLResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.CreateACLRequest](../../models/operations/createaclrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.CreateACLSecurity](../../models/operations/createaclsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.CreateACLResponse](../../models/operations/createaclresponse.md), error**


## DeleteACL

Delete an ACL from the specified service version. To remove an ACL from use, the ACL must be deleted from a draft version and the version without the ACL must be activated.

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
    res, err := s.ACL.DeleteACL(ctx, operations.DeleteACLRequest{
        ACLName: "test-acl",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteACLSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteACL200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.DeleteACLRequest](../../models/operations/deleteaclrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.DeleteACLSecurity](../../models/operations/deleteaclsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.DeleteACLResponse](../../models/operations/deleteaclresponse.md), error**


## GetACL

Retrieve a single ACL by name for the version and service.

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
    res, err := s.ACL.GetACL(ctx, operations.GetACLRequest{
        ACLName: "test-acl",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetACLSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ACLResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetACLRequest](../../models/operations/getaclrequest.md)   | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `security`                                                             | [operations.GetACLSecurity](../../models/operations/getaclsecurity.md) | :heavy_check_mark:                                                     | The security requirements to use for the request.                      |


### Response

**[*operations.GetACLResponse](../../models/operations/getaclresponse.md), error**


## ListAcls

List ACLs.

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
    res, err := s.ACL.ListAcls(ctx, operations.ListAclsRequest{
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.ListAclsSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ACLResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListAclsRequest](../../models/operations/listaclsrequest.md)   | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.ListAclsSecurity](../../models/operations/listaclssecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.ListAclsResponse](../../models/operations/listaclsresponse.md), error**


## UpdateACL

Update an ACL for a particular service and version.

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
    res, err := s.ACL.UpdateACL(ctx, operations.UpdateACLRequest{
        ACL: &shared.ACL{
            Name: sdk.String("test-acl"),
        },
        ACLName: "test-acl",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.UpdateACLSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ACLResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.UpdateACLRequest](../../models/operations/updateaclrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.UpdateACLSecurity](../../models/operations/updateaclsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.UpdateACLResponse](../../models/operations/updateaclresponse.md), error**
