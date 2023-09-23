// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateLogGcpPubsubRequest struct {
	LoggingGooglePubsub2 *shared.LoggingGooglePubsub2 `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingGooglePubsubName string `pathParam:"style=simple,explode=false,name=logging_google_pubsub_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogGcpPubsubRequest) GetLoggingGooglePubsub2() *shared.LoggingGooglePubsub2 {
	if o == nil {
		return nil
	}
	return o.LoggingGooglePubsub2
}

func (o *UpdateLogGcpPubsubRequest) GetLoggingGooglePubsubName() string {
	if o == nil {
		return ""
	}
	return o.LoggingGooglePubsubName
}

func (o *UpdateLogGcpPubsubRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogGcpPubsubRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogGcpPubsubResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingGooglePubsubResponse *shared.LoggingGooglePubsubResponse
}

func (o *UpdateLogGcpPubsubResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogGcpPubsubResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogGcpPubsubResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogGcpPubsubResponse) GetLoggingGooglePubsubResponse() *shared.LoggingGooglePubsubResponse {
	if o == nil {
		return nil
	}
	return o.LoggingGooglePubsubResponse
}
