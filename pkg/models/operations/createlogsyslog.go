// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateLogSyslogRequest struct {
	LoggingSyslog2 *shared.LoggingSyslog2 `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateLogSyslogRequest) GetLoggingSyslog2() *shared.LoggingSyslog2 {
	if o == nil {
		return nil
	}
	return o.LoggingSyslog2
}

func (o *CreateLogSyslogRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateLogSyslogRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type CreateLogSyslogResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingSyslogResponse *shared.LoggingSyslogResponse
}

func (o *CreateLogSyslogResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLogSyslogResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLogSyslogResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLogSyslogResponse) GetLoggingSyslogResponse() *shared.LoggingSyslogResponse {
	if o == nil {
		return nil
	}
	return o.LoggingSyslogResponse
}
