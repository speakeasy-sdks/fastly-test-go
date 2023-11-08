// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type UpdateHeaderObjectRequest struct {
	Header *components.Header `request:"mediaType=application/x-www-form-urlencoded"`
	// A handle to refer to this Header object.
	HeaderName string `pathParam:"style=simple,explode=false,name=header_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateHeaderObjectRequest) GetHeader() *components.Header {
	if o == nil {
		return nil
	}
	return o.Header
}

func (o *UpdateHeaderObjectRequest) GetHeaderName() string {
	if o == nil {
		return ""
	}
	return o.HeaderName
}

func (o *UpdateHeaderObjectRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateHeaderObjectRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateHeaderObjectResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	HeaderResponse *components.HeaderResponse
}

func (o *UpdateHeaderObjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateHeaderObjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateHeaderObjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateHeaderObjectResponse) GetHeaderResponse() *components.HeaderResponse {
	if o == nil {
		return nil
	}
	return o.HeaderResponse
}
