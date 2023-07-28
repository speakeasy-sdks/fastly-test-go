// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogSumologicSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *DeleteLogSumologicSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type DeleteLogSumologicRequest struct {
	// The name for the real-time logging configuration.
	LoggingSumologicName string `pathParam:"style=simple,explode=false,name=logging_sumologic_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteLogSumologicRequest) GetLoggingSumologicName() string {
	if o == nil {
		return ""
	}
	return o.LoggingSumologicName
}

func (o *DeleteLogSumologicRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteLogSumologicRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteLogSumologic200ApplicationJSON - OK
type DeleteLogSumologic200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteLogSumologic200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteLogSumologicResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteLogSumologic200ApplicationJSONObject *DeleteLogSumologic200ApplicationJSON
}

func (o *DeleteLogSumologicResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLogSumologicResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLogSumologicResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteLogSumologicResponse) GetDeleteLogSumologic200ApplicationJSONObject() *DeleteLogSumologic200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteLogSumologic200ApplicationJSONObject
}
