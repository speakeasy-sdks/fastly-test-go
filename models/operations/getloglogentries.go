// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type GetLogLogentriesRequest struct {
	// The name for the real-time logging configuration.
	LoggingLogentriesName string `pathParam:"style=simple,explode=false,name=logging_logentries_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetLogLogentriesRequest) GetLoggingLogentriesName() string {
	if o == nil {
		return ""
	}
	return o.LoggingLogentriesName
}

func (o *GetLogLogentriesRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetLogLogentriesRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetLogLogentriesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingLogentriesResponse *components.LoggingLogentriesResponse
}

func (o *GetLogLogentriesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLogLogentriesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLogLogentriesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLogLogentriesResponse) GetLoggingLogentriesResponse() *components.LoggingLogentriesResponse {
	if o == nil {
		return nil
	}
	return o.LoggingLogentriesResponse
}
