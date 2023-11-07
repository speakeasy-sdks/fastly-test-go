// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type EnableProductRequest struct {
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *EnableProductRequest) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

func (o *EnableProductRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type EnableProductResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	EnabledProductResponse *components.EnabledProductResponse
}

func (o *EnableProductResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EnableProductResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EnableProductResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *EnableProductResponse) GetEnabledProductResponse() *components.EnabledProductResponse {
	if o == nil {
		return nil
	}
	return o.EnabledProductResponse
}
