// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type CreateLogSyslogRequest struct {
	LoggingSyslog *components.LoggingSyslog `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateLogSyslogRequest) GetLoggingSyslog() *components.LoggingSyslog {
	if o == nil {
		return nil
	}
	return o.LoggingSyslog
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
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingSyslogResponse *components.LoggingSyslogResponse
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

func (o *CreateLogSyslogResponse) GetLoggingSyslogResponse() *components.LoggingSyslogResponse {
	if o == nil {
		return nil
	}
	return o.LoggingSyslogResponse
}
