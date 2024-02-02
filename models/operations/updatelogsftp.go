// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateLogSftpRequest struct {
	LoggingSftp *components.LoggingSftp `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingSftpName string `pathParam:"style=simple,explode=false,name=logging_sftp_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogSftpRequest) GetLoggingSftp() *components.LoggingSftp {
	if o == nil {
		return nil
	}
	return o.LoggingSftp
}

func (o *UpdateLogSftpRequest) GetLoggingSftpName() string {
	if o == nil {
		return ""
	}
	return o.LoggingSftpName
}

func (o *UpdateLogSftpRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogSftpRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogSftpResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingSftpResponse *components.LoggingSftpResponse
}

func (o *UpdateLogSftpResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogSftpResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogSftpResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogSftpResponse) GetLoggingSftpResponse() *components.LoggingSftpResponse {
	if o == nil {
		return nil
	}
	return o.LoggingSftpResponse
}
