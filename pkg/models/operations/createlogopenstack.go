// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateLogOpenstackSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type CreateLogOpenstackRequest struct {
	LoggingOpenstackInput *shared.LoggingOpenstackInput `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

type CreateLogOpenstackResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingOpenstackResponse *shared.LoggingOpenstackResponse
}
