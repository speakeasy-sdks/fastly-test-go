// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogLogentriesSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *DeleteLogLogentriesSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type DeleteLogLogentriesRequest struct {
	// The name for the real-time logging configuration.
	LoggingLogentriesName string `pathParam:"style=simple,explode=false,name=logging_logentries_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteLogLogentriesRequest) GetLoggingLogentriesName() string {
	if o == nil {
		return ""
	}
	return o.LoggingLogentriesName
}

func (o *DeleteLogLogentriesRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteLogLogentriesRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteLogLogentries200ApplicationJSON - OK
type DeleteLogLogentries200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteLogLogentries200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteLogLogentriesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteLogLogentries200ApplicationJSONObject *DeleteLogLogentries200ApplicationJSON
}

func (o *DeleteLogLogentriesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLogLogentriesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLogLogentriesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteLogLogentriesResponse) GetDeleteLogLogentries200ApplicationJSONObject() *DeleteLogLogentries200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteLogLogentries200ApplicationJSONObject
}
