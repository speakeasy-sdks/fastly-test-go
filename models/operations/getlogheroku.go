// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type GetLogHerokuRequest struct {
	// The name for the real-time logging configuration.
	LoggingHerokuName string `pathParam:"style=simple,explode=false,name=logging_heroku_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetLogHerokuRequest) GetLoggingHerokuName() string {
	if o == nil {
		return ""
	}
	return o.LoggingHerokuName
}

func (o *GetLogHerokuRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetLogHerokuRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetLogHerokuResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingHerokuResponse *components.LoggingHerokuResponse
}

func (o *GetLogHerokuResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLogHerokuResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLogHerokuResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLogHerokuResponse) GetLoggingHerokuResponse() *components.LoggingHerokuResponse {
	if o == nil {
		return nil
	}
	return o.LoggingHerokuResponse
}
