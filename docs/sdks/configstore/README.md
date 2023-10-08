# ConfigStore
(*ConfigStore*)

## Overview

A container that lets you store data in key-value pairs.

<https://developer.fastly.com/reference/api/services/resources/config-store>
### Available Operations

* [CreateConfigStore](#createconfigstore) - Create a config store
* [DeleteConfigStore](#deleteconfigstore) - Delete a config store
* [GetConfigStore](#getconfigstore) - Describe a config store
* [GetConfigStoreInfo](#getconfigstoreinfo) - Get config store metadata
* [ListConfigStoreServices](#listconfigstoreservices) - List linked services
* [ListConfigStores](#listconfigstores) - List config stores
* [UpdateConfigStore](#updateconfigstore) - Update a config store

## CreateConfigStore

Create a config store.

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
    res, err := s.ConfigStore.CreateConfigStore(ctx, &shared.ConfigStore{
        Name: fastly.String("test-config-store"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigStoreResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [shared.ConfigStore](../../models/shared/configstore.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


### Response

**[*operations.CreateConfigStoreResponse](../../models/operations/createconfigstoreresponse.md), error**


## DeleteConfigStore

Delete a config store.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ConfigStore.DeleteConfigStore(ctx, operations.DeleteConfigStoreRequest{
        ConfigStoreID: "7Lsb7Y76rChV9hSrv3KgFl",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteConfigStore200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteConfigStoreRequest](../../models/operations/deleteconfigstorerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteConfigStoreResponse](../../models/operations/deleteconfigstoreresponse.md), error**


## GetConfigStore

Describe a config store by its identifier.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ConfigStore.GetConfigStore(ctx, operations.GetConfigStoreRequest{
        ConfigStoreID: "7Lsb7Y76rChV9hSrv3KgFl",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigStoreResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetConfigStoreRequest](../../models/operations/getconfigstorerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetConfigStoreResponse](../../models/operations/getconfigstoreresponse.md), error**


## GetConfigStoreInfo

Retrieve metadata for a single config store.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ConfigStore.GetConfigStoreInfo(ctx, operations.GetConfigStoreInfoRequest{
        ConfigStoreID: "7Lsb7Y76rChV9hSrv3KgFl",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigStoreInfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetConfigStoreInfoRequest](../../models/operations/getconfigstoreinforequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetConfigStoreInfoResponse](../../models/operations/getconfigstoreinforesponse.md), error**


## ListConfigStoreServices

List services linked to a config store

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ConfigStore.ListConfigStoreServices(ctx, operations.ListConfigStoreServicesRequest{
        ConfigStoreID: "7Lsb7Y76rChV9hSrv3KgFl",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListConfigStoreServices200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListConfigStoreServicesRequest](../../models/operations/listconfigstoreservicesrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.ListConfigStoreServicesResponse](../../models/operations/listconfigstoreservicesresponse.md), error**


## ListConfigStores

List config stores.

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
    res, err := s.ConfigStore.ListConfigStores(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigStoreResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListConfigStoresResponse](../../models/operations/listconfigstoresresponse.md), error**


## UpdateConfigStore

Update a config store.

### Example Usage

```go
package main

import(
	"context"
	"log"
	fastly "Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := fastly.New(
        fastly.WithSecurity(shared.Security{
            Token: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ConfigStore.UpdateConfigStore(ctx, operations.UpdateConfigStoreRequest{
        ConfigStore: &shared.ConfigStore{
            Name: fastly.String("test-config-store"),
        },
        ConfigStoreID: "7Lsb7Y76rChV9hSrv3KgFl",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigStoreResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateConfigStoreRequest](../../models/operations/updateconfigstorerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateConfigStoreResponse](../../models/operations/updateconfigstoreresponse.md), error**

