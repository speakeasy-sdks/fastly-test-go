// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogLogentriesSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteLogLogentriesRequest struct {
	// The name for the real-time logging configuration.
	LoggingLogentriesName string `pathParam:"style=simple,explode=false,name=logging_logentries_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

// DeleteLogLogentries200ApplicationJSON - OK
type DeleteLogLogentries200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteLogLogentriesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteLogLogentries200ApplicationJSONObject *DeleteLogLogentries200ApplicationJSON
}
