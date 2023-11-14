// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type UpdateLogGcsRequest struct {
	LoggingGcs *components.LoggingGcs `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingGcsName string `pathParam:"style=simple,explode=false,name=logging_gcs_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogGcsRequest) GetLoggingGcs() *components.LoggingGcs {
	if o == nil {
		return nil
	}
	return o.LoggingGcs
}

func (o *UpdateLogGcsRequest) GetLoggingGcsName() string {
	if o == nil {
		return ""
	}
	return o.LoggingGcsName
}

func (o *UpdateLogGcsRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogGcsRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogGcsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingGcsResponse *components.LoggingGcsResponse
}

func (o *UpdateLogGcsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogGcsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogGcsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogGcsResponse) GetLoggingGcsResponse() *components.LoggingGcsResponse {
	if o == nil {
		return nil
	}
	return o.LoggingGcsResponse
}