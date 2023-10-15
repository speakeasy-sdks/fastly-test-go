<!-- Start SDK Example Usage -->


```go
package main

import (
	fastly "Fastly"
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
	"context"
	"log"
)

func main() {
	s := fastly.New(
		fastly.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.ACL.CreateACL(ctx, operations.CreateACLRequest{
		ACL: &shared.ACL{
			Name: fastly.String("test-acl"),
		},
		ServiceID: "SU1Z0isxPaozGVKXdv0eY",
		VersionID: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.ACLResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->