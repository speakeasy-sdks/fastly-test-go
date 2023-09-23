// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteACLRequest struct {
	// Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace.
	ACLName string `pathParam:"style=simple,explode=false,name=acl_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteACLRequest) GetACLName() string {
	if o == nil {
		return ""
	}
	return o.ACLName
}

func (o *DeleteACLRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteACLRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteACL200ApplicationJSON - OK
type DeleteACL200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteACL200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteACLResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteACL200ApplicationJSONObject *DeleteACL200ApplicationJSON
}

func (o *DeleteACLResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteACLResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteACLResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteACLResponse) GetDeleteACL200ApplicationJSONObject() *DeleteACL200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteACL200ApplicationJSONObject
}
