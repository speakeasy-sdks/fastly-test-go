// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PurgeAllRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *PurgeAllRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

// PurgeAll200ApplicationJSON - OK
type PurgeAll200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *PurgeAll200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type PurgeAllResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	PurgeAll200ApplicationJSONObject *PurgeAll200ApplicationJSON
}

func (o *PurgeAllResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PurgeAllResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PurgeAllResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PurgeAllResponse) GetPurgeAll200ApplicationJSONObject() *PurgeAll200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PurgeAll200ApplicationJSONObject
}
