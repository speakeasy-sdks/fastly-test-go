// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateCustomVclRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string      `pathParam:"style=simple,explode=false,name=service_id"`
	Vcl       *shared.Vcl `request:"mediaType=application/x-www-form-urlencoded"`
	// The name of this VCL.
	VclName string `pathParam:"style=simple,explode=false,name=vcl_name"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateCustomVclRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateCustomVclRequest) GetVcl() *shared.Vcl {
	if o == nil {
		return nil
	}
	return o.Vcl
}

func (o *UpdateCustomVclRequest) GetVclName() string {
	if o == nil {
		return ""
	}
	return o.VclName
}

func (o *UpdateCustomVclRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateCustomVclResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	VclResponse *shared.VclResponse
}

func (o *UpdateCustomVclResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCustomVclResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCustomVclResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateCustomVclResponse) GetVclResponse() *shared.VclResponse {
	if o == nil {
		return nil
	}
	return o.VclResponse
}
