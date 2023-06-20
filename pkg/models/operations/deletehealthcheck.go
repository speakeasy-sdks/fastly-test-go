// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteHealthcheckSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteHealthcheckRequest struct {
	// The name of the health check.
	HealthcheckName string `pathParam:"style=simple,explode=false,name=healthcheck_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

// DeleteHealthcheck200ApplicationJSON - OK
type DeleteHealthcheck200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteHealthcheckResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteHealthcheck200ApplicationJSONObject *DeleteHealthcheck200ApplicationJSON
}
