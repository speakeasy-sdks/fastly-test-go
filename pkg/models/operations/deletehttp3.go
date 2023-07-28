// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteHttp3Security struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *DeleteHttp3Security) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type DeleteHttp3Request struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteHttp3Request) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteHttp3Request) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteHttp3200ApplicationJSON - OK
type DeleteHttp3200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteHttp3200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteHttp3Response struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteHttp3200ApplicationJSONObject *DeleteHttp3200ApplicationJSON
}

func (o *DeleteHttp3Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteHttp3Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteHttp3Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteHttp3Response) GetDeleteHttp3200ApplicationJSONObject() *DeleteHttp3200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteHttp3200ApplicationJSONObject
}
