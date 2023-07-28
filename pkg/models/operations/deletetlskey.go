// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteTLSKeySecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *DeleteTLSKeySecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

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
	ContentType string
	StatusCode  int
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
