// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeletePoolServerRequest struct {
	// Alphanumeric string identifying a Pool.
	PoolID string `pathParam:"style=simple,explode=false,name=pool_id"`
	// Alphanumeric string identifying a Server.
	ServerID string `pathParam:"style=simple,explode=false,name=server_id"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *DeletePoolServerRequest) GetPoolID() string {
	if o == nil {
		return ""
	}
	return o.PoolID
}

func (o *DeletePoolServerRequest) GetServerID() string {
	if o == nil {
		return ""
	}
	return o.ServerID
}

func (o *DeletePoolServerRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

// DeletePoolServer200ApplicationJSON - OK
type DeletePoolServer200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeletePoolServer200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeletePoolServerResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeletePoolServer200ApplicationJSONObject *DeletePoolServer200ApplicationJSON
}

func (o *DeletePoolServerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeletePoolServerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeletePoolServerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeletePoolServerResponse) GetDeletePoolServer200ApplicationJSONObject() *DeletePoolServer200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeletePoolServer200ApplicationJSONObject
}
