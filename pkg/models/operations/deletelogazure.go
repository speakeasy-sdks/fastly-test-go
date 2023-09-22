// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogAzureRequest struct {
	// The name for the real-time logging configuration.
	LoggingAzureblobName string `pathParam:"style=simple,explode=false,name=logging_azureblob_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteLogAzureRequest) GetLoggingAzureblobName() string {
	if o == nil {
		return ""
	}
	return o.LoggingAzureblobName
}

func (o *DeleteLogAzureRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteLogAzureRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteLogAzure200ApplicationJSON - OK
type DeleteLogAzure200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteLogAzure200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteLogAzureResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteLogAzure200ApplicationJSONObject *DeleteLogAzure200ApplicationJSON
}

func (o *DeleteLogAzureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLogAzureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLogAzureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteLogAzureResponse) GetDeleteLogAzure200ApplicationJSONObject() *DeleteLogAzure200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteLogAzure200ApplicationJSONObject
}
