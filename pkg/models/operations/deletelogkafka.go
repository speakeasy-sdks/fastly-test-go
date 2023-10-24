// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogKafkaRequest struct {
	// The name for the real-time logging configuration.
	LoggingKafkaName string `pathParam:"style=simple,explode=false,name=logging_kafka_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteLogKafkaRequest) GetLoggingKafkaName() string {
	if o == nil {
		return ""
	}
	return o.LoggingKafkaName
}

func (o *DeleteLogKafkaRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteLogKafkaRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteLogKafka200ApplicationJSON - OK
type DeleteLogKafka200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteLogKafka200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteLogKafkaResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteLogKafka200ApplicationJSONObject *DeleteLogKafka200ApplicationJSON
}

func (o *DeleteLogKafkaResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLogKafkaResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLogKafkaResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteLogKafkaResponse) GetDeleteLogKafka200ApplicationJSONObject() *DeleteLogKafka200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteLogKafka200ApplicationJSONObject
}
