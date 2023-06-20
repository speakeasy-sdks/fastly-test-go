// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteTLSKeySecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteTLSKeyRequest struct {
	// Alphanumeric string identifying a private Key.
	TLSPrivateKeyID string `pathParam:"style=simple,explode=false,name=tls_private_key_id"`
}

type DeleteTLSKeyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
