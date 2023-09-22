// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteDirectorRequest struct {
	// Name for the Director.
	DirectorName string `pathParam:"style=simple,explode=false,name=director_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteDirectorRequest) GetDirectorName() string {
	if o == nil {
		return ""
	}
	return o.DirectorName
}

func (o *DeleteDirectorRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteDirectorRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteDirector200ApplicationJSON - OK
type DeleteDirector200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteDirector200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteDirectorResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteDirector200ApplicationJSONObject *DeleteDirector200ApplicationJSON
}

func (o *DeleteDirectorResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteDirectorResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteDirectorResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteDirectorResponse) GetDeleteDirector200ApplicationJSONObject() *DeleteDirector200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteDirector200ApplicationJSONObject
}
