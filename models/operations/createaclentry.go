// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"net/http"
)

type CreateACLEntryRequest struct {
	ACLEntry *components.ACLEntry `request:"mediaType=application/json"`
	// Alphanumeric string identifying a ACL.
	ACLID string `pathParam:"style=simple,explode=false,name=acl_id"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *CreateACLEntryRequest) GetACLEntry() *components.ACLEntry {
	if o == nil {
		return nil
	}
	return o.ACLEntry
}

func (o *CreateACLEntryRequest) GetACLID() string {
	if o == nil {
		return ""
	}
	return o.ACLID
}

func (o *CreateACLEntryRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type CreateACLEntryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ACLEntryResponse *components.ACLEntryResponse
}

func (o *CreateACLEntryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateACLEntryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateACLEntryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateACLEntryResponse) GetACLEntryResponse() *components.ACLEntryResponse {
	if o == nil {
		return nil
	}
	return o.ACLEntryResponse
}
