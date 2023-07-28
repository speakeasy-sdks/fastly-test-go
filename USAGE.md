<!-- Start SDK Example Usage -->


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
    operationSecurity := operations.CreateACLSecurity{
            Token: "",
        }

    ctx := context.Background()
    res, err := s.ACL.CreateACL(ctx, operations.CreateACLRequest{
        ACL: &shared.ACL{
            Name: sdk.String("test-acl"),
        },
        ServiceID: "SU1Z0isxPaozGVKXdv0eY",
        VersionID: 1,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ACLResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->