// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type ListLogGcsRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *ListLogGcsRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *ListLogGcsRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type ListLogGcsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingGcsResponses []shared.LoggingGcsResponse
}

func (o *ListLogGcsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListLogGcsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListLogGcsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListLogGcsResponse) GetLoggingGcsResponses() []shared.LoggingGcsResponse {
	if o == nil {
		return nil
	}
	return o.LoggingGcsResponses
}
