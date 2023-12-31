// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogElasticsearchRequest struct {
	// The name for the real-time logging configuration.
	LoggingElasticsearchName string `pathParam:"style=simple,explode=false,name=logging_elasticsearch_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteLogElasticsearchRequest) GetLoggingElasticsearchName() string {
	if o == nil {
		return ""
	}
	return o.LoggingElasticsearchName
}

func (o *DeleteLogElasticsearchRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteLogElasticsearchRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteLogElasticsearchResponseBody - OK
type DeleteLogElasticsearchResponseBody struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteLogElasticsearchResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteLogElasticsearchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *DeleteLogElasticsearchResponseBody
}

func (o *DeleteLogElasticsearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLogElasticsearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLogElasticsearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteLogElasticsearchResponse) GetObject() *DeleteLogElasticsearchResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
