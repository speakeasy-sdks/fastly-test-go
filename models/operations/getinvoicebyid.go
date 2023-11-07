// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"io"
	"net/http"
)

type GetInvoiceByIDRequest struct {
	// Alphanumeric string identifying the customer.
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
	// Alphanumeric string identifying the invoice.
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoice_id"`
}

func (o *GetInvoiceByIDRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *GetInvoiceByIDRequest) GetInvoiceID() string {
	if o == nil {
		return ""
	}
	return o.InvoiceID
}

type GetInvoiceByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	BillingResponse *components.BillingResponse
	// OK
	Res *string
	// OK
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Stream io.ReadCloser
}

func (o *GetInvoiceByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetInvoiceByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetInvoiceByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetInvoiceByIDResponse) GetBillingResponse() *components.BillingResponse {
	if o == nil {
		return nil
	}
	return o.BillingResponse
}

func (o *GetInvoiceByIDResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}

func (o *GetInvoiceByIDResponse) GetStream() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Stream
}
