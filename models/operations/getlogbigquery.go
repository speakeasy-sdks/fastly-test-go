// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type GetLogBigqueryRequest struct {
	// The name for the real-time logging configuration.
	LoggingBigqueryName string `pathParam:"style=simple,explode=false,name=logging_bigquery_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetLogBigqueryRequest) GetLoggingBigqueryName() string {
	if o == nil {
		return ""
	}
	return o.LoggingBigqueryName
}

func (o *GetLogBigqueryRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetLogBigqueryRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetLogBigqueryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingBigqueryResponse *components.LoggingBigqueryResponse
}

func (o *GetLogBigqueryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLogBigqueryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLogBigqueryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLogBigqueryResponse) GetLoggingBigqueryResponse() *components.LoggingBigqueryResponse {
	if o == nil {
		return nil
	}
	return o.LoggingBigqueryResponse
}
