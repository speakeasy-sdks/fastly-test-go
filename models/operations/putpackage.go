// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"net/http"
)

type PutPackageRequest struct {
	// We recommend using the Expect header because it may identify issues with the request based upon the headers alone instead of requiring you to wait until the entire binary package upload has completed.
	Expect        *string                   `header:"style=simple,explode=false,name=expect"`
	PackageUpload *components.PackageUpload `request:"mediaType=multipart/form-data"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *PutPackageRequest) GetExpect() *string {
	if o == nil {
		return nil
	}
	return o.Expect
}

func (o *PutPackageRequest) GetPackageUpload() *components.PackageUpload {
	if o == nil {
		return nil
	}
	return o.PackageUpload
}

func (o *PutPackageRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *PutPackageRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type PutPackageResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	PackageResponse *components.PackageResponse
}

func (o *PutPackageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutPackageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutPackageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutPackageResponse) GetPackageResponse() *components.PackageResponse {
	if o == nil {
		return nil
	}
	return o.PackageResponse
}
