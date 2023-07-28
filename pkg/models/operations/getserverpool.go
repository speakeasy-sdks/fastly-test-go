// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetServerPoolSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *GetServerPoolSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type GetServerPoolRequest struct {
	// Name for the Pool.
	PoolName string `pathParam:"style=simple,explode=false,name=pool_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetServerPoolRequest) GetPoolName() string {
	if o == nil {
		return ""
	}
	return o.PoolName
}

func (o *GetServerPoolRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetServerPoolRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetServerPoolResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PoolResponse *shared.PoolResponse
}

func (o *GetServerPoolResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetServerPoolResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetServerPoolResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetServerPoolResponse) GetPoolResponse() *shared.PoolResponse {
	if o == nil {
		return nil
	}
	return o.PoolResponse
}
