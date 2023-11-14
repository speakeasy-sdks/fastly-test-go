// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type GetPoolServerRequest struct {
	// Alphanumeric string identifying a Pool.
	PoolID string `pathParam:"style=simple,explode=false,name=pool_id"`
	// Alphanumeric string identifying a Server.
	ServerID string `pathParam:"style=simple,explode=false,name=server_id"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *GetPoolServerRequest) GetPoolID() string {
	if o == nil {
		return ""
	}
	return o.PoolID
}

func (o *GetPoolServerRequest) GetServerID() string {
	if o == nil {
		return ""
	}
	return o.ServerID
}

func (o *GetPoolServerRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type GetPoolServerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ServerResponse *components.ServerResponse
}

func (o *GetPoolServerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPoolServerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPoolServerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPoolServerResponse) GetServerResponse() *components.ServerResponse {
	if o == nil {
		return nil
	}
	return o.ServerResponse
}