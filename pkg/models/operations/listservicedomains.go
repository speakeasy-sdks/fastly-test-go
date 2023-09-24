// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type ListServiceDomainsRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *ListServiceDomainsRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type ListServiceDomainsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DomainResponses []shared.DomainResponse
}

func (o *ListServiceDomainsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListServiceDomainsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListServiceDomainsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListServiceDomainsResponse) GetDomainResponses() []shared.DomainResponse {
	if o == nil {
		return nil
	}
	return o.DomainResponses
}
