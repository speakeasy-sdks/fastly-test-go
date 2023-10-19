// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetLogHoneycombRequest struct {
	// The name for the real-time logging configuration.
	LoggingHoneycombName string `pathParam:"style=simple,explode=false,name=logging_honeycomb_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetLogHoneycombRequest) GetLoggingHoneycombName() string {
	if o == nil {
		return ""
	}
	return o.LoggingHoneycombName
}

func (o *GetLogHoneycombRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetLogHoneycombRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetLogHoneycombResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingHoneycomb *shared.LoggingHoneycomb
}

func (o *GetLogHoneycombResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLogHoneycombResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLogHoneycombResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLogHoneycombResponse) GetLoggingHoneycomb() *shared.LoggingHoneycomb {
	if o == nil {
		return nil
	}
	return o.LoggingHoneycomb
}
