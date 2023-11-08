// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type UpdateACLEntryRequest struct {
	ACLEntry *components.ACLEntry `request:"mediaType=application/json"`
	// Alphanumeric string identifying an ACL Entry.
	ACLEntryID string `pathParam:"style=simple,explode=false,name=acl_entry_id"`
	// Alphanumeric string identifying a ACL.
	ACLID string `pathParam:"style=simple,explode=false,name=acl_id"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *UpdateACLEntryRequest) GetACLEntry() *components.ACLEntry {
	if o == nil {
		return nil
	}
	return o.ACLEntry
}

func (o *UpdateACLEntryRequest) GetACLEntryID() string {
	if o == nil {
		return ""
	}
	return o.ACLEntryID
}

func (o *UpdateACLEntryRequest) GetACLID() string {
	if o == nil {
		return ""
	}
	return o.ACLID
}

func (o *UpdateACLEntryRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type UpdateACLEntryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ACLEntryResponse *components.ACLEntryResponse
}

func (o *UpdateACLEntryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateACLEntryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateACLEntryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateACLEntryResponse) GetACLEntryResponse() *components.ACLEntryResponse {
	if o == nil {
		return nil
	}
	return o.ACLEntryResponse
}
