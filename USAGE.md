<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	fastlytestgo "github.com/speakeasy-sdks/fastly-test-go"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"log"
)

func main() {
	s := fastlytestgo.New(
		fastlytestgo.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.ApexRedirect.DeleteApexRedirect(ctx, operations.DeleteApexRedirectRequest{
		ApexRedirectID: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->