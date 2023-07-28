// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateLogLogshuttleSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *UpdateLogLogshuttleSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type UpdateLogLogshuttleRequest struct {
	LoggingLogshuttle2 *shared.LoggingLogshuttle2 `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingLogshuttleName string `pathParam:"style=simple,explode=false,name=logging_logshuttle_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogLogshuttleRequest) GetLoggingLogshuttle2() *shared.LoggingLogshuttle2 {
	if o == nil {
		return nil
	}
	return o.LoggingLogshuttle2
}

func (o *UpdateLogLogshuttleRequest) GetLoggingLogshuttleName() string {
	if o == nil {
		return ""
	}
	return o.LoggingLogshuttleName
}

func (o *UpdateLogLogshuttleRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogLogshuttleRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogLogshuttleResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingLogshuttleResponse *shared.LoggingLogshuttleResponse
}

func (o *UpdateLogLogshuttleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogLogshuttleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogLogshuttleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogLogshuttleResponse) GetLoggingLogshuttleResponse() *shared.LoggingLogshuttleResponse {
	if o == nil {
		return nil
	}
	return o.LoggingLogshuttleResponse
}
