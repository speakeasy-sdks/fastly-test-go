# User

## Overview

A user of the Fastly API and web interface. A user is always associated with a customer. Some information may be limited depending on access level.

<https://developer.fastly.com/reference/api/account/user>
### Available Operations

* [CreateUser](#createuser) - Create a user
* [DeleteUser](#deleteuser) - Delete a user
* [GetCurrentUser](#getcurrentuser) - Get the current user
* [GetUser](#getuser) - Get a user
* [RequestPasswordReset](#requestpasswordreset) - Request a password reset
* [UpdateUser](#updateuser) - Update a user
* [UpdateUserPassword](#updateuserpassword) - Update the user's password

## CreateUser

Create a user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.User.CreateUser(ctx, shared.UserInput{
        LimitServices: sdk.Bool(false),
        Locked: sdk.Bool(false),
        Name: sdk.String("Melba Stokes I"),
        RequireNewPassword: sdk.Bool(false),
        Role: shared.RoleUserUser.ToPointer(),
        TwoFactorAuthEnabled: sdk.Bool(false),
        TwoFactorSetupRequired: sdk.Bool(false),
    }, operations.CreateUserSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.UserInput](../../models/shared/userinput.md)                           | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.CreateUserSecurity](../../models/operations/createusersecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.CreateUserResponse](../../models/operations/createuserresponse.md), error**


## DeleteUser

Delete a user.

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
    res, err := s.User.DeleteUser(ctx, operations.DeleteUserRequest{
        UserID: "x9KzsrACXZv8tPwlEDsKb6",
    }, operations.DeleteUserSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteUser200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.DeleteUserRequest](../../models/operations/deleteuserrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.DeleteUserSecurity](../../models/operations/deleteusersecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.DeleteUserResponse](../../models/operations/deleteuserresponse.md), error**


## GetCurrentUser

Get the logged in user.

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
    res, err := s.User.GetCurrentUser(ctx, operations.GetCurrentUserSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `security`                                                                             | [operations.GetCurrentUserSecurity](../../models/operations/getcurrentusersecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetCurrentUserResponse](../../models/operations/getcurrentuserresponse.md), error**


## GetUser

Get a specific user.

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
    res, err := s.User.GetUser(ctx, operations.GetUserRequest{
        UserID: "x9KzsrACXZv8tPwlEDsKb6",
    }, operations.GetUserSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetUserRequest](../../models/operations/getuserrequest.md)   | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `security`                                                               | [operations.GetUserSecurity](../../models/operations/getusersecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |


### Response

**[*operations.GetUserResponse](../../models/operations/getuserresponse.md), error**


## RequestPasswordReset

Requests a password reset for the specified user.

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
    res, err := s.User.RequestPasswordReset(ctx, operations.RequestPasswordResetRequest{
        UserLogin: "krisowner@example.com",
    }, operations.RequestPasswordResetSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestPasswordReset200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RequestPasswordResetRequest](../../models/operations/requestpasswordresetrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.RequestPasswordResetSecurity](../../models/operations/requestpasswordresetsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.RequestPasswordResetResponse](../../models/operations/requestpasswordresetresponse.md), error**


## UpdateUser

Update a user. Only users with the role of `superuser` can make changes to other users on the account. Non-superusers may use this endpoint to make changes to their own account. Two-factor attributes are not editable via this endpoint.

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
    res, err := s.User.UpdateUser(ctx, operations.UpdateUserRequest{
        UserInput: &shared.UserInput{
            LimitServices: sdk.Bool(false),
            Locked: sdk.Bool(false),
            Name: sdk.String("Tanya Buckridge"),
            RequireNewPassword: sdk.Bool(false),
            Role: shared.RoleUserUser.ToPointer(),
            TwoFactorAuthEnabled: sdk.Bool(false),
            TwoFactorSetupRequired: sdk.Bool(false),
        },
        UserID: "x9KzsrACXZv8tPwlEDsKb6",
    }, operations.UpdateUserSecurity{
        Token: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UpdateUserRequest](../../models/operations/updateuserrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.UpdateUserSecurity](../../models/operations/updateusersecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.UpdateUserResponse](../../models/operations/updateuserresponse.md), error**


## UpdateUserPassword

Update the user's password to a new one.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Fastly"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.User.UpdateUserPassword(ctx, shared.PasswordChange{
        NewPassword: sdk.String("praesentium"),
        OldPassword: sdk.String("maiores"),
    }, operations.UpdateUserPasswordSecurity{
        Password: "",
        Username: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [shared.PasswordChange](../../models/shared/passwordchange.md)                                 | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.UpdateUserPasswordSecurity](../../models/operations/updateuserpasswordsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.UpdateUserPasswordResponse](../../models/operations/updateuserpasswordresponse.md), error**
