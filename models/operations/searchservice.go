// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type SearchServiceRequest struct {
	// The name of the service.
	Name string `queryParam:"style=form,explode=true,name=name"`
}

func (o *SearchServiceRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type SearchServiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ServiceResponse *components.ServiceResponse
}

func (o *SearchServiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SearchServiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SearchServiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SearchServiceResponse) GetServiceResponse() *components.ServiceResponse {
	if o == nil {
		return nil
	}
	return o.ServiceResponse
}
