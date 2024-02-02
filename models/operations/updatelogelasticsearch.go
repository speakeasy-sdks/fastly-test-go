// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateLogElasticsearchRequest struct {
	LoggingElasticsearch *components.LoggingElasticsearch `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingElasticsearchName string `pathParam:"style=simple,explode=false,name=logging_elasticsearch_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogElasticsearchRequest) GetLoggingElasticsearch() *components.LoggingElasticsearch {
	if o == nil {
		return nil
	}
	return o.LoggingElasticsearch
}

func (o *UpdateLogElasticsearchRequest) GetLoggingElasticsearchName() string {
	if o == nil {
		return ""
	}
	return o.LoggingElasticsearchName
}

func (o *UpdateLogElasticsearchRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogElasticsearchRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogElasticsearchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingElasticsearchResponse *components.LoggingElasticsearchResponse
}

func (o *UpdateLogElasticsearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogElasticsearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogElasticsearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogElasticsearchResponse) GetLoggingElasticsearchResponse() *components.LoggingElasticsearchResponse {
	if o == nil {
		return nil
	}
	return o.LoggingElasticsearchResponse
}
