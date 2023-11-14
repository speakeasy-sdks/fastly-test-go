// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"net/http"
)

type GetLogSftpRequest struct {
	// The name for the real-time logging configuration.
	LoggingSftpName string `pathParam:"style=simple,explode=false,name=logging_sftp_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetLogSftpRequest) GetLoggingSftpName() string {
	if o == nil {
		return ""
	}
	return o.LoggingSftpName
}

func (o *GetLogSftpRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetLogSftpRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetLogSftpResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingSftpResponse *components.LoggingSftpResponse
}

func (o *GetLogSftpResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLogSftpResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLogSftpResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLogSftpResponse) GetLoggingSftpResponse() *components.LoggingSftpResponse {
	if o == nil {
		return nil
	}
	return o.LoggingSftpResponse
}
