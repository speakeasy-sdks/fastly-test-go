// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type CreateLogHoneycombRequest struct {
	LoggingHoneycomb *components.LoggingHoneycomb `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateLogHoneycombRequest) GetLoggingHoneycomb() *components.LoggingHoneycomb {
	if o == nil {
		return nil
	}
	return o.LoggingHoneycomb
}

func (o *CreateLogHoneycombRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateLogHoneycombRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type CreateLogHoneycombResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingHoneycomb *components.LoggingHoneycomb
}

func (o *CreateLogHoneycombResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLogHoneycombResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLogHoneycombResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLogHoneycombResponse) GetLoggingHoneycomb() *components.LoggingHoneycomb {
	if o == nil {
		return nil
	}
	return o.LoggingHoneycomb
}
