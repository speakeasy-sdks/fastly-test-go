// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateLogDigoceanRequest struct {
	LoggingDigitaloceanInput *shared.LoggingDigitaloceanInput `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateLogDigoceanRequest) GetLoggingDigitaloceanInput() *shared.LoggingDigitaloceanInput {
	if o == nil {
		return nil
	}
	return o.LoggingDigitaloceanInput
}

func (o *CreateLogDigoceanRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateLogDigoceanRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type CreateLogDigoceanResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingDigitaloceanResponse *shared.LoggingDigitaloceanResponse
}

func (o *CreateLogDigoceanResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLogDigoceanResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLogDigoceanResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLogDigoceanResponse) GetLoggingDigitaloceanResponse() *shared.LoggingDigitaloceanResponse {
	if o == nil {
		return nil
	}
	return o.LoggingDigitaloceanResponse
}
