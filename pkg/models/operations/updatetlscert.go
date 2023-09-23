// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateTLSCertRequest struct {
	TLSCertificateInput *shared.TLSCertificateInput `request:"mediaType=application/vnd.api+json"`
	// Alphanumeric string identifying a TLS certificate.
	TLSCertificateID string `pathParam:"style=simple,explode=false,name=tls_certificate_id"`
}

func (o *UpdateTLSCertRequest) GetTLSCertificateInput() *shared.TLSCertificateInput {
	if o == nil {
		return nil
	}
	return o.TLSCertificateInput
}

func (o *UpdateTLSCertRequest) GetTLSCertificateID() string {
	if o == nil {
		return ""
	}
	return o.TLSCertificateID
}

type UpdateTLSCertResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	TLSCertificateResponse *shared.TLSCertificateResponse
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

func (o *UpdateTLSCertResponse) GetTLSCertificateResponse() *shared.TLSCertificateResponse {
	if o == nil {
		return nil
	}
	return o.TLSCertificateResponse
}
