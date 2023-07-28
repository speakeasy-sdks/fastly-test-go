// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogCloudfilesSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *DeleteLogCloudfilesSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type DeleteLogCloudfilesRequest struct {
	// The name for the real-time logging configuration.
	LoggingCloudfilesName string `pathParam:"style=simple,explode=false,name=logging_cloudfiles_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteLogCloudfilesRequest) GetLoggingCloudfilesName() string {
	if o == nil {
		return ""
	}
	return o.LoggingCloudfilesName
}

func (o *DeleteLogCloudfilesRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteLogCloudfilesRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteLogCloudfiles200ApplicationJSON - OK
type DeleteLogCloudfiles200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteLogCloudfiles200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteLogCloudfilesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteLogCloudfiles200ApplicationJSONObject *DeleteLogCloudfiles200ApplicationJSON
}

func (o *DeleteLogCloudfilesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLogCloudfilesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLogCloudfilesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteLogCloudfilesResponse) GetDeleteLogCloudfiles200ApplicationJSONObject() *DeleteLogCloudfiles200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteLogCloudfiles200ApplicationJSONObject
}
