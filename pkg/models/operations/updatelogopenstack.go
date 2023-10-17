// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateLogOpenstackRequest struct {
	LoggingOpenstackInput *shared.LoggingOpenstackInput `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingOpenstackName string `pathParam:"style=simple,explode=false,name=logging_openstack_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogOpenstackRequest) GetLoggingOpenstackInput() *shared.LoggingOpenstackInput {
	if o == nil {
		return nil
	}
	return o.LoggingOpenstackInput
}

func (o *UpdateLogOpenstackRequest) GetLoggingOpenstackName() string {
	if o == nil {
		return ""
	}
	return o.LoggingOpenstackName
}

func (o *UpdateLogOpenstackRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogOpenstackRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogOpenstackResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingOpenstackResponse *shared.LoggingOpenstackResponse
}

func (o *UpdateLogOpenstackResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogOpenstackResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogOpenstackResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogOpenstackResponse) GetLoggingOpenstackResponse() *shared.LoggingOpenstackResponse {
	if o == nil {
		return nil
	}
	return o.LoggingOpenstackResponse
}
