// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type CreateLogDatadogRequest struct {
	LoggingDatadog *components.LoggingDatadog `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateLogDatadogRequest) GetLoggingDatadog() *components.LoggingDatadog {
	if o == nil {
		return nil
	}
	return o.LoggingDatadog
}

func (o *CreateLogDatadogRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateLogDatadogRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type CreateLogDatadogResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingDatadogResponse *components.LoggingDatadogResponse
}

func (o *CreateLogDatadogResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLogDatadogResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLogDatadogResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLogDatadogResponse) GetLoggingDatadogResponse() *components.LoggingDatadogResponse {
	if o == nil {
		return nil
	}
	return o.LoggingDatadogResponse
}