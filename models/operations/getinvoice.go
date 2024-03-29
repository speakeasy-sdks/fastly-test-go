// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"io"
	"net/http"
)

type GetInvoiceRequest struct {
	// 2-digit month.
	Month string `pathParam:"style=simple,explode=false,name=month"`
	// 4-digit year.
	Year string `pathParam:"style=simple,explode=false,name=year"`
}

func (o *GetInvoiceRequest) GetMonth() string {
	if o == nil {
		return ""
	}
	return o.Month
}

func (o *GetInvoiceRequest) GetYear() string {
	if o == nil {
		return ""
	}
	return o.Year
}

type GetInvoiceResponse struct {
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

func (o *GetInvoiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetInvoiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetInvoiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetInvoiceResponse) GetBillingResponse() *components.BillingResponse {
	if o == nil {
		return nil
	}
	return o.BillingResponse
}

func (o *GetInvoiceResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}

func (o *GetInvoiceResponse) GetStream() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Stream
}
