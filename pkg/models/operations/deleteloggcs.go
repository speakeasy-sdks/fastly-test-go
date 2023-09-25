// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogGcsRequest struct {
	// The name for the real-time logging configuration.
	LoggingGcsName string `pathParam:"style=simple,explode=false,name=logging_gcs_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteLogGcsRequest) GetLoggingGcsName() string {
	if o == nil {
		return ""
	}
	return o.LoggingGcsName
}

func (o *DeleteLogGcsRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteLogGcsRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteLogGcs200ApplicationJSON - OK
type DeleteLogGcs200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteLogGcs200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteLogGcsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteLogGcs200ApplicationJSONObject *DeleteLogGcs200ApplicationJSON
}

func (o *DeleteLogGcsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLogGcsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLogGcsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteLogGcsResponse) GetDeleteLogGcs200ApplicationJSONObject() *DeleteLogGcs200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteLogGcs200ApplicationJSONObject
}
