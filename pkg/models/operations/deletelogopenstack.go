// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogOpenstackSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteLogOpenstackRequest struct {
	// The name for the real-time logging configuration.
	LoggingOpenstackName string `pathParam:"style=simple,explode=false,name=logging_openstack_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

// DeleteLogOpenstack200ApplicationJSON - OK
type DeleteLogOpenstack200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteLogOpenstackResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteLogOpenstack200ApplicationJSONObject *DeleteLogOpenstack200ApplicationJSON
}
