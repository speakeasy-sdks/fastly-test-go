// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetTLSBulkCertSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *GetTLSBulkCertSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type GetTLSBulkCertRequest struct {
	// Alphanumeric string identifying a TLS bulk certificate.
	CertificateID string `pathParam:"style=simple,explode=false,name=certificate_id"`
}

func (o *GetTLSBulkCertRequest) GetCertificateID() string {
	if o == nil {
		return ""
	}
	return o.CertificateID
}

type GetTLSBulkCertResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	TLSBulkCertificateResponse *shared.TLSBulkCertificateResponse
}

func (o *GetTLSBulkCertResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTLSBulkCertResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTLSBulkCertResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTLSBulkCertResponse) GetTLSBulkCertificateResponse() *shared.TLSBulkCertificateResponse {
	if o == nil {
		return nil
	}
	return o.TLSBulkCertificateResponse
}
