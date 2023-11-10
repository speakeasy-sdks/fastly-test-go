// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type UpdateTLSCertRequest struct {
	TLSCertificate *components.TLSCertificate `request:"mediaType=application/vnd.api+json"`
	// Alphanumeric string identifying a TLS certificate.
	TLSCertificateID string `pathParam:"style=simple,explode=false,name=tls_certificate_id"`
}

func (o *UpdateTLSCertRequest) GetTLSCertificate() *components.TLSCertificate {
	if o == nil {
		return nil
	}
	return o.TLSCertificate
}

func (o *UpdateTLSCertRequest) GetTLSCertificateID() string {
	if o == nil {
		return ""
	}
	return o.TLSCertificateID
}

type UpdateTLSCertResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TLSCertificateResponse *components.TLSCertificateResponse
}

func (o *UpdateTLSCertResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTLSCertResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTLSCertResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTLSCertResponse) GetTLSCertificateResponse() *components.TLSCertificateResponse {
	if o == nil {
		return nil
	}
	return o.TLSCertificateResponse
}
