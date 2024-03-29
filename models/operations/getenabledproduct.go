// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type GetEnabledProductRequest struct {
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *GetEnabledProductRequest) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

func (o *GetEnabledProductRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type GetEnabledProductResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	EnabledProductResponse *components.EnabledProductResponse
}

func (o *GetEnabledProductResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEnabledProductResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEnabledProductResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetEnabledProductResponse) GetEnabledProductResponse() *components.EnabledProductResponse {
	if o == nil {
		return nil
	}
	return o.EnabledProductResponse
}
