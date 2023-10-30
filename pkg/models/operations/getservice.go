// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetServiceRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *GetServiceRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type GetServiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ServiceResponse *shared.ServiceResponse
}

func (o *GetServiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetServiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetServiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetServiceResponse) GetServiceResponse() *shared.ServiceResponse {
	if o == nil {
		return nil
	}
	return o.ServiceResponse
}
