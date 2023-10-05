// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateLogLogshuttleRequest struct {
	LoggingLogshuttle *shared.LoggingLogshuttle `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateLogLogshuttleRequest) GetLoggingLogshuttle() *shared.LoggingLogshuttle {
	if o == nil {
		return nil
	}
	return o.LoggingLogshuttle
}

func (o *CreateLogLogshuttleRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateLogLogshuttleRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type CreateLogLogshuttleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingLogshuttleResponse *shared.LoggingLogshuttleResponse
}

func (o *CreateLogLogshuttleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLogLogshuttleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLogLogshuttleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLogLogshuttleResponse) GetLoggingLogshuttleResponse() *shared.LoggingLogshuttleResponse {
	if o == nil {
		return nil
	}
	return o.LoggingLogshuttleResponse
}
