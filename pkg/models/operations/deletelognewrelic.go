// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogNewrelicRequest struct {
	// The name for the real-time logging configuration.
	LoggingNewrelicName string `pathParam:"style=simple,explode=false,name=logging_newrelic_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteLogNewrelicRequest) GetLoggingNewrelicName() string {
	if o == nil {
		return ""
	}
	return o.LoggingNewrelicName
}

func (o *DeleteLogNewrelicRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteLogNewrelicRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteLogNewrelic200ApplicationJSON - OK
type DeleteLogNewrelic200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteLogNewrelic200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteLogNewrelicResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteLogNewrelic200ApplicationJSONObject *DeleteLogNewrelic200ApplicationJSON
}

func (o *DeleteLogNewrelicResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLogNewrelicResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLogNewrelicResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteLogNewrelicResponse) GetDeleteLogNewrelic200ApplicationJSONObject() *DeleteLogNewrelic200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteLogNewrelic200ApplicationJSONObject
}
