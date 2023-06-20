// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateServiceSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type UpdateServiceRequest struct {
	Service *shared.Service `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

type UpdateServiceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ServiceResponse *shared.ServiceResponse
}
