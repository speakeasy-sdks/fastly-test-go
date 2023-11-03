// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogDatadogRequest struct {
	// The name for the real-time logging configuration.
	LoggingDatadogName string `pathParam:"style=simple,explode=false,name=logging_datadog_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteLogDatadogRequest) GetLoggingDatadogName() string {
	if o == nil {
		return ""
	}
	return o.LoggingDatadogName
}

func (o *DeleteLogDatadogRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteLogDatadogRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteLogDatadog200ApplicationJSON - OK
type DeleteLogDatadog200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteLogDatadog200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteLogDatadogResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteLogDatadog200ApplicationJSONObject *DeleteLogDatadog200ApplicationJSON
}

func (o *DeleteLogDatadogResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLogDatadogResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLogDatadogResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteLogDatadogResponse) GetDeleteLogDatadog200ApplicationJSONObject() *DeleteLogDatadog200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteLogDatadog200ApplicationJSONObject
}
