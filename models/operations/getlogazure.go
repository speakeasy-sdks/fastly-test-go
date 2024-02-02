// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type GetLogAzureRequest struct {
	// The name for the real-time logging configuration.
	LoggingAzureblobName string `pathParam:"style=simple,explode=false,name=logging_azureblob_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetLogAzureRequest) GetLoggingAzureblobName() string {
	if o == nil {
		return ""
	}
	return o.LoggingAzureblobName
}

func (o *GetLogAzureRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetLogAzureRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetLogAzureResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingAzureblobResponse *components.LoggingAzureblobResponse
}

func (o *GetLogAzureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLogAzureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLogAzureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLogAzureResponse) GetLoggingAzureblobResponse() *components.LoggingAzureblobResponse {
	if o == nil {
		return nil
	}
	return o.LoggingAzureblobResponse
}
