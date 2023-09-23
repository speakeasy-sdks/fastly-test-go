// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateLogHoneycombRequest struct {
	LoggingHoneycomb2 *shared.LoggingHoneycomb2 `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingHoneycombName string `pathParam:"style=simple,explode=false,name=logging_honeycomb_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogHoneycombRequest) GetLoggingHoneycomb2() *shared.LoggingHoneycomb2 {
	if o == nil {
		return nil
	}
	return o.LoggingHoneycomb2
}

func (o *UpdateLogHoneycombRequest) GetLoggingHoneycombName() string {
	if o == nil {
		return ""
	}
	return o.LoggingHoneycombName
}

func (o *UpdateLogHoneycombRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogHoneycombRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogHoneycombResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingHoneycombResponse *shared.LoggingHoneycombResponse
}

func (o *UpdateLogHoneycombResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogHoneycombResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogHoneycombResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogHoneycombResponse) GetLoggingHoneycombResponse() *shared.LoggingHoneycombResponse {
	if o == nil {
		return nil
	}
	return o.LoggingHoneycombResponse
}
