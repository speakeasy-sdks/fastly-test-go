// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type SetCustomVclMainRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// The name of this VCL.
	VclName string `pathParam:"style=simple,explode=false,name=vcl_name"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *SetCustomVclMainRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *SetCustomVclMainRequest) GetVclName() string {
	if o == nil {
		return ""
	}
	return o.VclName
}

func (o *SetCustomVclMainRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type SetCustomVclMainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	VclResponse *components.VclResponse
}

func (o *SetCustomVclMainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SetCustomVclMainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SetCustomVclMainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SetCustomVclMainResponse) GetVclResponse() *components.VclResponse {
	if o == nil {
		return nil
	}
	return o.VclResponse
}
