// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogSplunkRequest struct {
	// The name for the real-time logging configuration.
	LoggingSplunkName string `pathParam:"style=simple,explode=false,name=logging_splunk_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteLogSplunkRequest) GetLoggingSplunkName() string {
	if o == nil {
		return ""
	}
	return o.LoggingSplunkName
}

func (o *DeleteLogSplunkRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteLogSplunkRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteLogSplunkResponseBody - OK
type DeleteLogSplunkResponseBody struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteLogSplunkResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteLogSplunkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *DeleteLogSplunkResponseBody
}

func (o *DeleteLogSplunkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLogSplunkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLogSplunkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteLogSplunkResponse) GetObject() *DeleteLogSplunkResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
