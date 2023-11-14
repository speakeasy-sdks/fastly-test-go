// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"net/http"
)

type CreateLogHerokuRequest struct {
	LoggingHeroku *components.LoggingHeroku `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateLogHerokuRequest) GetLoggingHeroku() *components.LoggingHeroku {
	if o == nil {
		return nil
	}
	return o.LoggingHeroku
}

func (o *CreateLogHerokuRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateLogHerokuRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type CreateLogHerokuResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingHerokuResponse *components.LoggingHerokuResponse
}

func (o *CreateLogHerokuResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLogHerokuResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLogHerokuResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLogHerokuResponse) GetLoggingHerokuResponse() *components.LoggingHerokuResponse {
	if o == nil {
		return nil
	}
	return o.LoggingHerokuResponse
}
