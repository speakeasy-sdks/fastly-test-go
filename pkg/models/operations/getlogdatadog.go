// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetLogDatadogSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *GetLogDatadogSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type GetLogDatadogRequest struct {
	// The name for the real-time logging configuration.
	LoggingDatadogName string `pathParam:"style=simple,explode=false,name=logging_datadog_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetLogDatadogRequest) GetLoggingDatadogName() string {
	if o == nil {
		return ""
	}
	return o.LoggingDatadogName
}

func (o *GetLogDatadogRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetLogDatadogRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetLogDatadogResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingDatadogResponse *shared.LoggingDatadogResponse
}

func (o *GetLogDatadogResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLogDatadogResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLogDatadogResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLogDatadogResponse) GetLoggingDatadogResponse() *shared.LoggingDatadogResponse {
	if o == nil {
		return nil
	}
	return o.LoggingDatadogResponse
}
