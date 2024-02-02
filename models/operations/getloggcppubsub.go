// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type GetLogGcpPubsubRequest struct {
	// The name for the real-time logging configuration.
	LoggingGooglePubsubName string `pathParam:"style=simple,explode=false,name=logging_google_pubsub_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetLogGcpPubsubRequest) GetLoggingGooglePubsubName() string {
	if o == nil {
		return ""
	}
	return o.LoggingGooglePubsubName
}

func (o *GetLogGcpPubsubRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetLogGcpPubsubRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetLogGcpPubsubResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingGooglePubsubResponse *components.LoggingGooglePubsubResponse
}

func (o *GetLogGcpPubsubResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLogGcpPubsubResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLogGcpPubsubResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLogGcpPubsubResponse) GetLoggingGooglePubsubResponse() *components.LoggingGooglePubsubResponse {
	if o == nil {
		return nil
	}
	return o.LoggingGooglePubsubResponse
}
