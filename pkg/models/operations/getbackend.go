// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetBackendSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type GetBackendRequest struct {
	// The name of the backend.
	BackendName string `pathParam:"style=simple,explode=false,name=backend_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

type GetBackendResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	BackendResponse *shared.BackendResponse
}
