// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateBulkTLSCertRequest struct {
	// Alphanumeric string identifying a TLS bulk certificate.
	CertificateID      string                         `pathParam:"style=simple,explode=false,name=certificate_id"`
	TLSBulkCertificate *components.TLSBulkCertificate `request:"mediaType=application/vnd.api+json"`
}

func (o *UpdateBulkTLSCertRequest) GetCertificateID() string {
	if o == nil {
		return ""
	}
	return o.CertificateID
}

func (o *UpdateBulkTLSCertRequest) GetTLSBulkCertificate() *components.TLSBulkCertificate {
	if o == nil {
		return nil
	}
	return o.TLSBulkCertificate
}

type UpdateBulkTLSCertResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TLSBulkCertificateResponse *components.TLSBulkCertificateResponse
}

func (o *UpdateBulkTLSCertResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateBulkTLSCertResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateBulkTLSCertResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateBulkTLSCertResponse) GetTLSBulkCertificateResponse() *components.TLSBulkCertificateResponse {
	if o == nil {
		return nil
	}
	return o.TLSBulkCertificateResponse
}
