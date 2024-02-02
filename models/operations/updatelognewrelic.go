// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateLogNewrelicRequest struct {
	LoggingNewrelic *components.LoggingNewrelic `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingNewrelicName string `pathParam:"style=simple,explode=false,name=logging_newrelic_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogNewrelicRequest) GetLoggingNewrelic() *components.LoggingNewrelic {
	if o == nil {
		return nil
	}
	return o.LoggingNewrelic
}

func (o *UpdateLogNewrelicRequest) GetLoggingNewrelicName() string {
	if o == nil {
		return ""
	}
	return o.LoggingNewrelicName
}

func (o *UpdateLogNewrelicRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogNewrelicRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogNewrelicResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingNewrelicResponse *components.LoggingNewrelicResponse
}

func (o *UpdateLogNewrelicResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogNewrelicResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogNewrelicResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogNewrelicResponse) GetLoggingNewrelicResponse() *components.LoggingNewrelicResponse {
	if o == nil {
		return nil
	}
	return o.LoggingNewrelicResponse
}
