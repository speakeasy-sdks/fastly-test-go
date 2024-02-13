# ACLEntry
(*ACLEntry*)

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
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.ACLEntry.BulkUpdateACLEntries(ctx, operations.BulkUpdateACLEntriesRequest{
        ACLID: "6tUXdegLTf5BCig0zGFrU3",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.BulkUpdateACLEntriesRequest](../../models/operations/bulkupdateaclentriesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.BulkUpdateACLEntriesResponse](../../models/operations/bulkupdateaclentriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateACLEntry

Add an ACL entry to an ACL.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.ACLEntry.CreateACLEntry(ctx, operations.CreateACLEntryRequest{
        ACLID: "6tUXdegLTf5BCig0zGFrU3",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ACLEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateACLEntryRequest](../../models/operations/createaclentryrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateACLEntryResponse](../../models/operations/createaclentryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteACLEntry

Delete an ACL entry from a specified ACL.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.ACLEntry.DeleteACLEntry(ctx, operations.DeleteACLEntryRequest{
        ACLEntryID: "6yxNzlOpW1V7JfSwvLGtOc",
        ACLID: "6tUXdegLTf5BCig0zGFrU3",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteACLEntryRequest](../../models/operations/deleteaclentryrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteACLEntryResponse](../../models/operations/deleteaclentryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetACLEntry

Retrieve a single ACL entry.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.ACLEntry.GetACLEntry(ctx, operations.GetACLEntryRequest{
        ACLEntryID: "6yxNzlOpW1V7JfSwvLGtOc",
        ACLID: "6tUXdegLTf5BCig0zGFrU3",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ACLEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetACLEntryRequest](../../models/operations/getaclentryrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetACLEntryResponse](../../models/operations/getaclentryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListACLEntries

List ACL entries for a specified ACL.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.ACLEntry.ListACLEntries(ctx, operations.ListACLEntriesRequest{
        ACLID: "6tUXdegLTf5BCig0zGFrU3",
        Direction: components.DirectionAscend.ToPointer(),
        Page: fastlytestgo.Int64(1),
        PerPage: fastlytestgo.Int64(20),
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        Sort: fastlytestgo.String("created"),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListACLEntriesRequest](../../models/operations/listaclentriesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListACLEntriesResponse](../../models/operations/listaclentriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateACLEntry

Update an ACL entry for a specified ACL.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go/v2"
	"context"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"log"
)

func main() {
    s := fastlytestgo.New(
        fastlytestgo.WithSecurity("YOUR-API-KEY"),
    )

    ctx := context.Background()
    res, err := s.ACLEntry.UpdateACLEntry(ctx, operations.UpdateACLEntryRequest{
        ACLEntryID: "6yxNzlOpW1V7JfSwvLGtOc",
        ACLID: "6tUXdegLTf5BCig0zGFrU3",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ACLEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateACLEntryRequest](../../models/operations/updateaclentryrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateACLEntryResponse](../../models/operations/updateaclentryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
