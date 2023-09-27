// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateLogElasticsearchRequest struct {
	LoggingElasticsearch2 *shared.LoggingElasticsearch2 `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateLogElasticsearchRequest) GetLoggingElasticsearch2() *shared.LoggingElasticsearch2 {
	if o == nil {
		return nil
	}
	return o.LoggingElasticsearch2
}

func (o *CreateLogElasticsearchRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateLogElasticsearchRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type CreateLogElasticsearchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingElasticsearchResponse *shared.LoggingElasticsearchResponse
}

func (o *CreateLogElasticsearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLogElasticsearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLogElasticsearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLogElasticsearchResponse) GetLoggingElasticsearchResponse() *shared.LoggingElasticsearchResponse {
	if o == nil {
		return nil
	}
	return o.LoggingElasticsearchResponse
}
