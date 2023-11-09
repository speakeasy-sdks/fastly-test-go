// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type CreateServiceAuthorizationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ServiceAuthorizationResponse *components.ServiceAuthorizationResponse
}

func (o *CreateServiceAuthorizationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateServiceAuthorizationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateServiceAuthorizationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateServiceAuthorizationResponse) GetServiceAuthorizationResponse() *components.ServiceAuthorizationResponse {
	if o == nil {
		return nil
	}
	return o.ServiceAuthorizationResponse
}
