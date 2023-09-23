// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateServerPoolRequest struct {
	Pool2 *shared.Pool2 `request:"mediaType=application/x-www-form-urlencoded"`
	// Name for the Pool.
	PoolName string `pathParam:"style=simple,explode=false,name=pool_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateServerPoolRequest) GetPool2() *shared.Pool2 {
	if o == nil {
		return nil
	}
	return o.Pool2
}

func (o *UpdateServerPoolRequest) GetPoolName() string {
	if o == nil {
		return ""
	}
	return o.PoolName
}

func (o *UpdateServerPoolRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateServerPoolRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateServerPoolResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PoolResponse *shared.PoolResponse
}

func (o *UpdateServerPoolResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateServerPoolResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateServerPoolResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateServerPoolResponse) GetPoolResponse() *shared.PoolResponse {
	if o == nil {
		return nil
	}
	return o.PoolResponse
}
