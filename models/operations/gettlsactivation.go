// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type GetTLSActivationRequest struct {
	// Include related objects. Optional, comma-separated values. Permitted values: `tls_certificate`, `tls_configuration`, and `tls_domain`.
	//
	Include *string `queryParam:"style=form,explode=false,name=include"`
	// Alphanumeric string identifying a TLS activation.
	TLSActivationID string `pathParam:"style=simple,explode=false,name=tls_activation_id"`
}

func (o *GetTLSActivationRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *GetTLSActivationRequest) GetTLSActivationID() string {
	if o == nil {
		return ""
	}
	return o.TLSActivationID
}

type GetTLSActivationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TLSActivationResponse *components.TLSActivationResponse
}

func (o *GetTLSActivationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTLSActivationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTLSActivationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTLSActivationResponse) GetTLSActivationResponse() *components.TLSActivationResponse {
	if o == nil {
		return nil
	}
	return o.TLSActivationResponse
}
