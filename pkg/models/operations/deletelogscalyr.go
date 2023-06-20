// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogScalyrSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteLogScalyrRequest struct {
	// The name for the real-time logging configuration.
	LoggingScalyrName string `pathParam:"style=simple,explode=false,name=logging_scalyr_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

// DeleteLogScalyr200ApplicationJSON - OK
type DeleteLogScalyr200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteLogScalyrResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteLogScalyr200ApplicationJSONObject *DeleteLogScalyr200ApplicationJSON
}