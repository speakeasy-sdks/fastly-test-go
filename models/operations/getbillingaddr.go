// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type GetBillingAddrRequest struct {
	// Alphanumeric string identifying the customer.
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
}

func (o *GetBillingAddrRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

type GetBillingAddrResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	BillingAddressResponse *components.BillingAddressResponse
}

func (o *GetBillingAddrResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBillingAddrResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBillingAddrResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetBillingAddrResponse) GetBillingAddressResponse() *components.BillingAddressResponse {
	if o == nil {
		return nil
	}
	return o.BillingAddressResponse
}
