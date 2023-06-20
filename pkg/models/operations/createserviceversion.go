// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateServiceVersionSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type CreateServiceVersionRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

type CreateServiceVersionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	VersionCreateResponse *shared.VersionCreateResponse
}
