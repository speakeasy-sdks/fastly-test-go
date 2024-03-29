// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateLogSumologicRequest struct {
	LoggingSumologic *components.LoggingSumologic `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingSumologicName string `pathParam:"style=simple,explode=false,name=logging_sumologic_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogSumologicRequest) GetLoggingSumologic() *components.LoggingSumologic {
	if o == nil {
		return nil
	}
	return o.LoggingSumologic
}

func (o *UpdateLogSumologicRequest) GetLoggingSumologicName() string {
	if o == nil {
		return ""
	}
	return o.LoggingSumologicName
}

func (o *UpdateLogSumologicRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogSumologicRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogSumologicResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingSumologicResponse *components.LoggingSumologicResponse
}

func (o *UpdateLogSumologicResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogSumologicResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogSumologicResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogSumologicResponse) GetLoggingSumologicResponse() *components.LoggingSumologicResponse {
	if o == nil {
		return nil
	}
	return o.LoggingSumologicResponse
}
