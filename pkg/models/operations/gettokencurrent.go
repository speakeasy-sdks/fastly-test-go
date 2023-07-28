// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetTokenCurrentSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *GetTokenCurrentSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type GetTokenCurrentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Missing or expired token.
	GenericTokenError *shared.GenericTokenError
	// OK
	TokenResponse *shared.TokenResponse
}

func (o *GetTokenCurrentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTokenCurrentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTokenCurrentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTokenCurrentResponse) GetGenericTokenError() *shared.GenericTokenError {
	if o == nil {
		return nil
	}
	return o.GenericTokenError
}

func (o *GetTokenCurrentResponse) GetTokenResponse() *shared.TokenResponse {
	if o == nil {
		return nil
	}
	return o.TokenResponse
}
