// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateLogAzureRequest struct {
	LoggingAzureblobInput *shared.LoggingAzureblobInput `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateLogAzureRequest) GetLoggingAzureblobInput() *shared.LoggingAzureblobInput {
	if o == nil {
		return nil
	}
	return o.LoggingAzureblobInput
}

func (o *CreateLogAzureRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateLogAzureRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type CreateLogAzureResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingAzureblobResponse *shared.LoggingAzureblobResponse
}

func (o *CreateLogAzureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLogAzureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLogAzureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLogAzureResponse) GetLoggingAzureblobResponse() *shared.LoggingAzureblobResponse {
	if o == nil {
		return nil
	}
	return o.LoggingAzureblobResponse
}
