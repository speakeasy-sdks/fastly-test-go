// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteHeaderObjectRequest struct {
	// A handle to refer to this Header object.
	HeaderName string `pathParam:"style=simple,explode=false,name=header_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteHeaderObjectRequest) GetHeaderName() string {
	if o == nil {
		return ""
	}
	return o.HeaderName
}

func (o *DeleteHeaderObjectRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteHeaderObjectRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteHeaderObject200ApplicationJSON - OK
type DeleteHeaderObject200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteHeaderObject200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteHeaderObjectResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteHeaderObject200ApplicationJSONObject *DeleteHeaderObject200ApplicationJSON
}

func (o *DeleteHeaderObjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteHeaderObjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteHeaderObjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteHeaderObjectResponse) GetDeleteHeaderObject200ApplicationJSONObject() *DeleteHeaderObject200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteHeaderObject200ApplicationJSONObject
}
