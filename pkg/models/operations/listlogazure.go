// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type ListLogAzureRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *ListLogAzureRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *ListLogAzureRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type ListLogAzureResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingAzureblobResponses []shared.LoggingAzureblobResponse
}

func (o *ListLogAzureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListLogAzureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListLogAzureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListLogAzureResponse) GetLoggingAzureblobResponses() []shared.LoggingAzureblobResponse {
	if o == nil {
		return nil
	}
	return o.LoggingAzureblobResponses
}
