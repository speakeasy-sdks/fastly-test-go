// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogKinesisRequest struct {
	// The name for the real-time logging configuration.
	LoggingKinesisName string `pathParam:"style=simple,explode=false,name=logging_kinesis_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteLogKinesisRequest) GetLoggingKinesisName() string {
	if o == nil {
		return ""
	}
	return o.LoggingKinesisName
}

func (o *DeleteLogKinesisRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteLogKinesisRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteLogKinesisResponseBody - OK
type DeleteLogKinesisResponseBody struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteLogKinesisResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteLogKinesisResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *DeleteLogKinesisResponseBody
}

func (o *DeleteLogKinesisResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLogKinesisResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLogKinesisResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteLogKinesisResponse) GetObject() *DeleteLogKinesisResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
