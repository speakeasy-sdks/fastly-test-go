// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteTLSKeyRequest struct {
	// Alphanumeric string identifying a private Key.
	TLSPrivateKeyID string `pathParam:"style=simple,explode=false,name=tls_private_key_id"`
}

func (o *DeleteTLSKeyRequest) GetTLSPrivateKeyID() string {
	if o == nil {
		return ""
	}
	return o.TLSPrivateKeyID
}

type DeleteTLSKeyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteTLSKeyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteTLSKeyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteTLSKeyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}