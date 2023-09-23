// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteACLEntryRequest struct {
	// Alphanumeric string identifying an ACL Entry.
	ACLEntryID string `pathParam:"style=simple,explode=false,name=acl_entry_id"`
	// Alphanumeric string identifying a ACL.
	ACLID string `pathParam:"style=simple,explode=false,name=acl_id"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *DeleteACLEntryRequest) GetACLEntryID() string {
	if o == nil {
		return ""
	}
	return o.ACLEntryID
}

func (o *DeleteACLEntryRequest) GetACLID() string {
	if o == nil {
		return ""
	}
	return o.ACLID
}

func (o *DeleteACLEntryRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

// DeleteACLEntry200ApplicationJSON - OK
type DeleteACLEntry200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteACLEntry200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteACLEntryResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteACLEntry200ApplicationJSONObject *DeleteACLEntry200ApplicationJSON
}

func (o *DeleteACLEntryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteACLEntryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteACLEntryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteACLEntryResponse) GetDeleteACLEntry200ApplicationJSONObject() *DeleteACLEntry200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteACLEntry200ApplicationJSONObject
}
