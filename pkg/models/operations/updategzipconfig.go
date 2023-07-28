// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateGzipConfigSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *UpdateGzipConfigSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type UpdateGzipConfigRequest struct {
	Gzip *shared.Gzip `request:"mediaType=application/x-www-form-urlencoded"`
	// Name of the gzip configuration.
	GzipName string `pathParam:"style=simple,explode=false,name=gzip_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateGzipConfigRequest) GetGzip() *shared.Gzip {
	if o == nil {
		return nil
	}
	return o.Gzip
}

func (o *UpdateGzipConfigRequest) GetGzipName() string {
	if o == nil {
		return ""
	}
	return o.GzipName
}

func (o *UpdateGzipConfigRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateGzipConfigRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateGzipConfigResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GzipResponse *shared.GzipResponse
}

func (o *UpdateGzipConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateGzipConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateGzipConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateGzipConfigResponse) GetGzipResponse() *shared.GzipResponse {
	if o == nil {
		return nil
	}
	return o.GzipResponse
}
