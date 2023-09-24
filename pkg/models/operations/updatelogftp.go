// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateLogFtpRequest struct {
	LoggingFtpInput *shared.LoggingFtpInput `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingFtpName string `pathParam:"style=simple,explode=false,name=logging_ftp_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogFtpRequest) GetLoggingFtpInput() *shared.LoggingFtpInput {
	if o == nil {
		return nil
	}
	return o.LoggingFtpInput
}

func (o *UpdateLogFtpRequest) GetLoggingFtpName() string {
	if o == nil {
		return ""
	}
	return o.LoggingFtpName
}

func (o *UpdateLogFtpRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogFtpRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogFtpResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingFtpResponse *shared.LoggingFtpResponse
}

func (o *UpdateLogFtpResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogFtpResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogFtpResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogFtpResponse) GetLoggingFtpResponse() *shared.LoggingFtpResponse {
	if o == nil {
		return nil
	}
	return o.LoggingFtpResponse
}
