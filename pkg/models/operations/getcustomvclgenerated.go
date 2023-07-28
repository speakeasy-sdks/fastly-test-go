// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetCustomVclGeneratedSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *GetCustomVclGeneratedSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type GetCustomVclGeneratedRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetCustomVclGeneratedRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetCustomVclGeneratedRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetCustomVclGeneratedResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	VclResponse *shared.VclResponse
}

func (o *GetCustomVclGeneratedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCustomVclGeneratedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCustomVclGeneratedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCustomVclGeneratedResponse) GetVclResponse() *shared.VclResponse {
	if o == nil {
		return nil
	}
	return o.VclResponse
}
