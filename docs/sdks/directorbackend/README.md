# DirectorBackend

## Overview

Maps and relates backends as belonging to directors. Backends can belong to any number of directors but directors can only hold one reference to a specific backend.

<https://developer.fastly.com/reference/api/load-balancing/directors/backend>
### Available Operations

* [CreateDirectorBackend](#createdirectorbackend) - Create a director-backend relationship
* [DeleteDirectorBackend](#deletedirectorbackend) - Delete a director-backend relationship
* [GetDirectorBackend](#getdirectorbackend) - Get a director-backend relationship

## CreateDirectorBackend

Establishes a relationship between a Backend and a Director. The Backend is then considered a member of the Director and can be used to balance traffic onto.

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
    res, err := s.DirectorBackend.CreateDirectorBackend(ctx, operations.CreateDirectorBackendRequest{
        BackendName: "test-backend",
        DirectorName: "test-director",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.CreateDirectorBackendSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectorBackend != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateDirectorBackendRequest](../../models/operations/createdirectorbackendrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.CreateDirectorBackendSecurity](../../models/operations/createdirectorbackendsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.CreateDirectorBackendResponse](../../models/operations/createdirectorbackendresponse.md), error**


## DeleteDirectorBackend

Deletes the relationship between a Backend and a Director. The Backend is no longer considered a member of the Director and thus will not have traffic balanced onto it from this Director.

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
    res, err := s.DirectorBackend.DeleteDirectorBackend(ctx, operations.DeleteDirectorBackendRequest{
        BackendName: "test-backend",
        DirectorName: "test-director",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.DeleteDirectorBackendSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDirectorBackend200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteDirectorBackendRequest](../../models/operations/deletedirectorbackendrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.DeleteDirectorBackendSecurity](../../models/operations/deletedirectorbackendsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.DeleteDirectorBackendResponse](../../models/operations/deletedirectorbackendresponse.md), error**


## GetDirectorBackend

Returns the relationship between a Backend and a Director. If the Backend has been associated with the Director, it returns a simple record indicating this. Otherwise, returns a 404.

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
    res, err := s.DirectorBackend.GetDirectorBackend(ctx, operations.GetDirectorBackendRequest{
        BackendName: "test-backend",
        DirectorName: "test-director",
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operations.GetDirectorBackendSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectorBackend != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetDirectorBackendRequest](../../models/operations/getdirectorbackendrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.GetDirectorBackendSecurity](../../models/operations/getdirectorbackendsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.GetDirectorBackendResponse](../../models/operations/getdirectorbackendresponse.md), error**

