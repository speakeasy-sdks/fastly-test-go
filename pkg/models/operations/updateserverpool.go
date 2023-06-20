// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateServerPoolSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type UpdateServerPoolRequest struct {
	Pool2 *shared.Pool2 `request:"mediaType=application/x-www-form-urlencoded"`
	// Name for the Pool.
	PoolName string `pathParam:"style=simple,explode=false,name=pool_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

type UpdateServerPoolResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PoolResponse *shared.PoolResponse
}
