// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type UpdateBillingAddrRequest struct {
	// Alphanumeric string identifying the customer.
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
	// One or more billing address attributes
	UpdateBillingAddressRequest *components.UpdateBillingAddressRequest `request:"mediaType=application/vnd.api+json"`
}

func (o *UpdateBillingAddrRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *UpdateBillingAddrRequest) GetUpdateBillingAddressRequest() *components.UpdateBillingAddressRequest {
	if o == nil {
		return nil
	}
	return o.UpdateBillingAddressRequest
}

type UpdateBillingAddrResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	BillingAddressResponse *components.BillingAddressResponse
}

func (o *UpdateBillingAddrResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateBillingAddrResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateBillingAddrResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateBillingAddrResponse) GetBillingAddressResponse() *components.BillingAddressResponse {
	if o == nil {
		return nil
	}
	return o.BillingAddressResponse
}