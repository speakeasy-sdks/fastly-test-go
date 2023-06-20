// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogDigoceanSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteLogDigoceanRequest struct {
	// The name for the real-time logging configuration.
	LoggingDigitaloceanName string `pathParam:"style=simple,explode=false,name=logging_digitalocean_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

// DeleteLogDigocean200ApplicationJSON - OK
type DeleteLogDigocean200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteLogDigoceanResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteLogDigocean200ApplicationJSONObject *DeleteLogDigocean200ApplicationJSON
}
