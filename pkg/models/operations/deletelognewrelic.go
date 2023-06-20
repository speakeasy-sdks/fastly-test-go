// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogNewrelicSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteLogNewrelicRequest struct {
	// The name for the real-time logging configuration.
	LoggingNewrelicName string `pathParam:"style=simple,explode=false,name=logging_newrelic_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

// DeleteLogNewrelic200ApplicationJSON - OK
type DeleteLogNewrelic200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteLogNewrelicResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteLogNewrelic200ApplicationJSONObject *DeleteLogNewrelic200ApplicationJSON
}
