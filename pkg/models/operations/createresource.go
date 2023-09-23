// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateResourceRequest struct {
	Resource *shared.Resource `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateResourceRequest) GetResource() *shared.Resource {
	if o == nil {
		return nil
	}
	return o.Resource
}

func (o *CreateResourceRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateResourceRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type CreateResourceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ResourceResponse *shared.ResourceResponse
}

func (o *CreateResourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateResourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateResourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateResourceResponse) GetResourceResponse() *shared.ResourceResponse {
	if o == nil {
		return nil
	}
	return o.ResourceResponse
}
