// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"net/http"
)

type GetLogHTTPSRequest struct {
	// The name for the real-time logging configuration.
	LoggingHTTPSName string `pathParam:"style=simple,explode=false,name=logging_https_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetLogHTTPSRequest) GetLoggingHTTPSName() string {
	if o == nil {
		return ""
	}
	return o.LoggingHTTPSName
}

func (o *GetLogHTTPSRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetLogHTTPSRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetLogHTTPSResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingHTTPSResponse *components.LoggingHTTPSResponse
}

func (o *GetLogHTTPSResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLogHTTPSResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLogHTTPSResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLogHTTPSResponse) GetLoggingHTTPSResponse() *components.LoggingHTTPSResponse {
	if o == nil {
		return nil
	}
	return o.LoggingHTTPSResponse
}
