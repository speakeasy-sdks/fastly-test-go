# ACLEntry

## Overview

An ACL entry holds an individual IP address or subnet range and is a member of an ACL. ACL entries are versionless, which means they can be created, modified, or deleted without activating a new version of your service.


<https://developer.fastly.com/reference/api/acls/acl-entry>
### Available Operations

* [BulkUpdateACLEntries](#bulkupdateaclentries) - Update multiple ACL entries
* [CreateACLEntry](#createaclentry) - Create an ACL entry
* [DeleteACLEntry](#deleteaclentry) - Delete an ACL entry
* [GetACLEntry](#getaclentry) - Describe an ACL entry
* [ListACLEntries](#listaclentries) - List ACL entries
* [UpdateACLEntry](#updateaclentry) - Update an ACL entry

## BulkUpdateACLEntries

Update multiple ACL entries on the same ACL. For faster updates to your service, group your changes into large batches. The maximum batch size is 1000 entries. [Contact support](https://support.fastly.com/) to discuss raising this limit.

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
    operationSecurity := operations.BulkUpdateACLEntriesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.ACLEntry.BulkUpdateACLEntries(ctx, operations.BulkUpdateACLEntriesRequest{
        ACLID: "6tUXdegLTf5BCig0zGFrU3",
        BulkUpdateACLEntriesRequestInput: &shared.BulkUpdateACLEntriesRequestInput{
            Entries: []shared.BulkUpdateACLEntryInput{
                shared.BulkUpdateACLEntryInput{
                    Comment: sdk.String("provident"),
                    IP: sdk.String("127.0.0.1"),
                    Negated: shared.BulkUpdateACLEntryNegatedOne.ToPointer(),
                    Op: shared.BulkUpdateACLEntryOpDelete.ToPointer(),
                    Subnet: sdk.Int64(8),
                },
                shared.BulkUpdateACLEntryInput{
                    Comment: sdk.String("unde"),
                    IP: sdk.String("127.0.0.1"),
                    Negated: shared.BulkUpdateACLEntryNegatedOne.ToPointer(),
                    Op: shared.BulkUpdateACLEntryOpUpdate.ToPointer(),
                    Subnet: sdk.Int64(8),
                },
                shared.BulkUpdateACLEntryInput{
                    Comment: sdk.String("illum"),
                    IP: sdk.String("127.0.0.1"),
                    Negated: shared.BulkUpdateACLEntryNegatedZero.ToPointer(),
                    Op: shared.BulkUpdateACLEntryOpUpdate.ToPointer(),
                    Subnet: sdk.Int64(8),
                },
            },
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.BulkUpdateACLEntries200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.BulkUpdateACLEntriesRequest](../../models/operations/bulkupdateaclentriesrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.BulkUpdateACLEntriesSecurity](../../models/operations/bulkupdateaclentriessecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.BulkUpdateACLEntriesResponse](../../models/operations/bulkupdateaclentriesresponse.md), error**


## CreateACLEntry

Add an ACL entry to an ACL.

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
    operationSecurity := operations.CreateACLEntrySecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.ACLEntry.CreateACLEntry(ctx, operations.CreateACLEntryRequest{
        ACLEntry: &shared.ACLEntry{
            Comment: sdk.String("deserunt"),
            IP: sdk.String("127.0.0.1"),
            Negated: shared.ACLEntryNegatedZero.ToPointer(),
            Subnet: sdk.Int64(8),
        },
        ACLID: "6tUXdegLTf5BCig0zGFrU3",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ACLEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateACLEntryRequest](../../models/operations/createaclentryrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.CreateACLEntrySecurity](../../models/operations/createaclentrysecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.CreateACLEntryResponse](../../models/operations/createaclentryresponse.md), error**


## DeleteACLEntry

Delete an ACL entry from a specified ACL.

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
    operationSecurity := operations.DeleteACLEntrySecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.ACLEntry.DeleteACLEntry(ctx, operations.DeleteACLEntryRequest{
        ACLEntryID: "6yxNzlOpW1V7JfSwvLGtOc",
        ACLID: "6tUXdegLTf5BCig0zGFrU3",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteACLEntry200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteACLEntryRequest](../../models/operations/deleteaclentryrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.DeleteACLEntrySecurity](../../models/operations/deleteaclentrysecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.DeleteACLEntryResponse](../../models/operations/deleteaclentryresponse.md), error**


## GetACLEntry

Retrieve a single ACL entry.

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
    operationSecurity := operations.GetACLEntrySecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.ACLEntry.GetACLEntry(ctx, operations.GetACLEntryRequest{
        ACLEntryID: "6yxNzlOpW1V7JfSwvLGtOc",
        ACLID: "6tUXdegLTf5BCig0zGFrU3",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ACLEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetACLEntryRequest](../../models/operations/getaclentryrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.GetACLEntrySecurity](../../models/operations/getaclentrysecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.GetACLEntryResponse](../../models/operations/getaclentryresponse.md), error**


## ListACLEntries

List ACL entries for a specified ACL.

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
    operationSecurity := operations.ListACLEntriesSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.ACLEntry.ListACLEntries(ctx, operations.ListACLEntriesRequest{
        ACLID: "6tUXdegLTf5BCig0zGFrU3",
        Direction: shared.DirectionAscend.ToPointer(),
        Page: sdk.Int64(1),
        PerPage: sdk.Int64(20),
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        Sort: sdk.String("created"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ACLEntryResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListACLEntriesRequest](../../models/operations/listaclentriesrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.ListACLEntriesSecurity](../../models/operations/listaclentriessecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.ListACLEntriesResponse](../../models/operations/listaclentriesresponse.md), error**


## UpdateACLEntry

Update an ACL entry for a specified ACL.

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
    operationSecurity := operations.UpdateACLEntrySecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.ACLEntry.UpdateACLEntry(ctx, operations.UpdateACLEntryRequest{
        ACLEntry: &shared.ACLEntry{
            Comment: sdk.String("iure"),
            IP: sdk.String("127.0.0.1"),
            Negated: shared.ACLEntryNegatedZero.ToPointer(),
            Subnet: sdk.Int64(8),
        },
        ACLEntryID: "6yxNzlOpW1V7JfSwvLGtOc",
        ACLID: "6tUXdegLTf5BCig0zGFrU3",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ACLEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateACLEntryRequest](../../models/operations/updateaclentryrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.UpdateACLEntrySecurity](../../models/operations/updateaclentrysecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.UpdateACLEntryResponse](../../models/operations/updateaclentryresponse.md), error**

