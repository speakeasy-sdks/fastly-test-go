<!-- Start SDK Example Usage -->


```go
package main

import (
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
	"context"
	"log"
)

func main() {
	s := fastly.New(
		fastly.WithSecurity(""),
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