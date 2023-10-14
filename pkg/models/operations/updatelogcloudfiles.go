// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateLogCloudfilesRequest struct {
	LoggingCloudfilesInput *shared.LoggingCloudfilesInput `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingCloudfilesName string `pathParam:"style=simple,explode=false,name=logging_cloudfiles_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogCloudfilesRequest) GetLoggingCloudfilesInput() *shared.LoggingCloudfilesInput {
	if o == nil {
		return nil
	}
	return o.LoggingCloudfilesInput
}

func (o *UpdateLogCloudfilesRequest) GetLoggingCloudfilesName() string {
	if o == nil {
		return ""
	}
	return o.LoggingCloudfilesName
}

func (o *UpdateLogCloudfilesRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogCloudfilesRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogCloudfilesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingCloudfilesResponse *shared.LoggingCloudfilesResponse
}

func (o *UpdateLogCloudfilesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogCloudfilesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogCloudfilesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogCloudfilesResponse) GetLoggingCloudfilesResponse() *shared.LoggingCloudfilesResponse {
	if o == nil {
		return nil
	}
	return o.LoggingCloudfilesResponse
}
