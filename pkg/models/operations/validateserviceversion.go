// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ValidateServiceVersionSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *ValidateServiceVersionSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type ValidateServiceVersionRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *ValidateServiceVersionRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *ValidateServiceVersionRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// ValidateServiceVersion200ApplicationJSON - OK
type ValidateServiceVersion200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *ValidateServiceVersion200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type ValidateServiceVersionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ValidateServiceVersion200ApplicationJSONObject *ValidateServiceVersion200ApplicationJSON
}

func (o *ValidateServiceVersionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ValidateServiceVersionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ValidateServiceVersionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ValidateServiceVersionResponse) GetValidateServiceVersion200ApplicationJSONObject() *ValidateServiceVersion200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ValidateServiceVersion200ApplicationJSONObject
}
