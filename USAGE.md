<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"log"
)

func main() {
	s := fastlytestgo.New(
		fastlytestgo.WithSecurity("YOUR-API-KEY"),
	)

	ctx := context.Background()
	res, err := s.User.CreateUser(ctx, &components.User{
		Role: components.RoleUserUser.ToPointer(),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.UserResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->