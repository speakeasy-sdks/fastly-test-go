// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateLogSplunkRequest struct {
	LoggingSplunk *components.LoggingSplunk `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingSplunkName string `pathParam:"style=simple,explode=false,name=logging_splunk_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogSplunkRequest) GetLoggingSplunk() *components.LoggingSplunk {
	if o == nil {
		return nil
	}
	return o.LoggingSplunk
}

func (o *UpdateLogSplunkRequest) GetLoggingSplunkName() string {
	if o == nil {
		return ""
	}
	return o.LoggingSplunkName
}

func (o *UpdateLogSplunkRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogSplunkRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogSplunkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingSplunkResponse *components.LoggingSplunkResponse
}

func (o *UpdateLogSplunkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogSplunkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogSplunkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogSplunkResponse) GetLoggingSplunkResponse() *components.LoggingSplunkResponse {
	if o == nil {
		return nil
	}
	return o.LoggingSplunkResponse
}
