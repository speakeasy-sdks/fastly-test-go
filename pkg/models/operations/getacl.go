// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetACLRequest struct {
	// Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace.
	ACLName string `pathParam:"style=simple,explode=false,name=acl_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetACLRequest) GetACLName() string {
	if o == nil {
		return ""
	}
	return o.ACLName
}

func (o *GetACLRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetACLRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetACLResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ACLResponse *shared.ACLResponse
}

func (o *GetACLResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetACLResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetACLResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetACLResponse) GetACLResponse() *shared.ACLResponse {
	if o == nil {
		return nil
	}
	return o.ACLResponse
}
