// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetTLSKeySecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type GetTLSKeyRequest struct {
	// Alphanumeric string identifying a private Key.
	TLSPrivateKeyID string `pathParam:"style=simple,explode=false,name=tls_private_key_id"`
}

type GetTLSKeyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	TLSPrivateKeyResponse *shared.TLSPrivateKeyResponse
}