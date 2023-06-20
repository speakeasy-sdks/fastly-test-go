// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogHTTPSSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteLogHTTPSRequest struct {
	// The name for the real-time logging configuration.
	LoggingHTTPSName string `pathParam:"style=simple,explode=false,name=logging_https_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

// DeleteLogHTTPS200ApplicationJSON - OK
type DeleteLogHTTPS200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteLogHTTPSResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteLogHTTPS200ApplicationJSONObject *DeleteLogHTTPS200ApplicationJSON
}
