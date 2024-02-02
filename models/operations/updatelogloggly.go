// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateLogLogglyRequest struct {
	LoggingLoggly *components.LoggingLoggly `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingLogglyName string `pathParam:"style=simple,explode=false,name=logging_loggly_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogLogglyRequest) GetLoggingLoggly() *components.LoggingLoggly {
	if o == nil {
		return nil
	}
	return o.LoggingLoggly
}

func (o *UpdateLogLogglyRequest) GetLoggingLogglyName() string {
	if o == nil {
		return ""
	}
	return o.LoggingLogglyName
}

func (o *UpdateLogLogglyRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogLogglyRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogLogglyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingLogglyResponse *components.LoggingLogglyResponse
}

func (o *UpdateLogLogglyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogLogglyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogLogglyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogLogglyResponse) GetLoggingLogglyResponse() *components.LoggingLogglyResponse {
	if o == nil {
		return nil
	}
	return o.LoggingLogglyResponse
}
