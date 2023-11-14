// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type CreatePoolServerRequest struct {
	// Alphanumeric string identifying a Pool.
	PoolID string             `pathParam:"style=simple,explode=false,name=pool_id"`
	Server *components.Server `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *CreatePoolServerRequest) GetPoolID() string {
	if o == nil {
		return ""
	}
	return o.PoolID
}

func (o *CreatePoolServerRequest) GetServer() *components.Server {
	if o == nil {
		return nil
	}
	return o.Server
}

func (o *CreatePoolServerRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type CreatePoolServerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ServerResponse *components.ServerResponse
}

func (o *CreatePoolServerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePoolServerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePoolServerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreatePoolServerResponse) GetServerResponse() *components.ServerResponse {
	if o == nil {
		return nil
	}
	return o.ServerResponse
}