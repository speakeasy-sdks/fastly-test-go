// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type CreateLogFtpRequest struct {
	LoggingFtp *components.LoggingFtp `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateLogFtpRequest) GetLoggingFtp() *components.LoggingFtp {
	if o == nil {
		return nil
	}
	return o.LoggingFtp
}

func (o *CreateLogFtpRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateLogFtpRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type CreateLogFtpResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingFtpResponse *components.LoggingFtpResponse
}

func (o *CreateLogFtpResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLogFtpResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLogFtpResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLogFtpResponse) GetLoggingFtpResponse() *components.LoggingFtpResponse {
	if o == nil {
		return nil
	}
	return o.LoggingFtpResponse
}
