// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type UpdateLogDatadogRequest struct {
	LoggingDatadog *components.LoggingDatadog `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingDatadogName string `pathParam:"style=simple,explode=false,name=logging_datadog_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogDatadogRequest) GetLoggingDatadog() *components.LoggingDatadog {
	if o == nil {
		return nil
	}
	return o.LoggingDatadog
}

func (o *UpdateLogDatadogRequest) GetLoggingDatadogName() string {
	if o == nil {
		return ""
	}
	return o.LoggingDatadogName
}

func (o *UpdateLogDatadogRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogDatadogRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogDatadogResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingDatadogResponse *components.LoggingDatadogResponse
}

func (o *UpdateLogDatadogResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogDatadogResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogDatadogResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogDatadogResponse) GetLoggingDatadogResponse() *components.LoggingDatadogResponse {
	if o == nil {
		return nil
	}
	return o.LoggingDatadogResponse
}