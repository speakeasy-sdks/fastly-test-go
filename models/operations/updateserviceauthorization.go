// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateServiceAuthorizationRequest struct {
	ServiceAuthorization *components.ServiceAuthorization `request:"mediaType=application/json"`
	// Alphanumeric string identifying a service authorization.
	ServiceAuthorizationID string `pathParam:"style=simple,explode=false,name=service_authorization_id"`
}

func (o *UpdateServiceAuthorizationRequest) GetServiceAuthorization() *components.ServiceAuthorization {
	if o == nil {
		return nil
	}
	return o.ServiceAuthorization
}

func (o *UpdateServiceAuthorizationRequest) GetServiceAuthorizationID() string {
	if o == nil {
		return ""
	}
	return o.ServiceAuthorizationID
}

type UpdateServiceAuthorizationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ServiceAuthorizationResponse *components.ServiceAuthorizationResponse
}

func (o *UpdateServiceAuthorizationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateServiceAuthorizationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateServiceAuthorizationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateServiceAuthorizationResponse) GetServiceAuthorizationResponse() *components.ServiceAuthorizationResponse {
	if o == nil {
		return nil
	}
	return o.ServiceAuthorizationResponse
}
