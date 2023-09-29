// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetGzipConfigsRequest struct {
	// Name of the gzip configuration.
	GzipName string `pathParam:"style=simple,explode=false,name=gzip_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetGzipConfigsRequest) GetGzipName() string {
	if o == nil {
		return ""
	}
	return o.GzipName
}

func (o *GetGzipConfigsRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetGzipConfigsRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetGzipConfigsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GzipResponse *shared.GzipResponse
}

func (o *GetGzipConfigsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGzipConfigsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGzipConfigsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetGzipConfigsResponse) GetGzipResponse() *shared.GzipResponse {
	if o == nil {
		return nil
	}
	return o.GzipResponse
}
