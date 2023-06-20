// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteDirectorSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteDirectorRequest struct {
	// Name for the Director.
	DirectorName string `pathParam:"style=simple,explode=false,name=director_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

// DeleteDirector200ApplicationJSON - OK
type DeleteDirector200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteDirectorResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteDirector200ApplicationJSONObject *DeleteDirector200ApplicationJSON
}
