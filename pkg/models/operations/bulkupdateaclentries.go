// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type BulkUpdateACLEntriesSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type BulkUpdateACLEntriesRequest struct {
	// Alphanumeric string identifying a ACL.
	ACLID                            string                                   `pathParam:"style=simple,explode=false,name=acl_id"`
	BulkUpdateACLEntriesRequestInput *shared.BulkUpdateACLEntriesRequestInput `request:"mediaType=application/json"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

// BulkUpdateACLEntries200ApplicationJSON - OK
type BulkUpdateACLEntries200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type BulkUpdateACLEntriesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	BulkUpdateACLEntries200ApplicationJSONObject *BulkUpdateACLEntries200ApplicationJSON
}