// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type RevokeTokenCurrentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Token revocation error.
	GenericTokenError *shared.GenericTokenError
}

func (o *RevokeTokenCurrentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RevokeTokenCurrentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RevokeTokenCurrentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RevokeTokenCurrentResponse) GetGenericTokenError() *shared.GenericTokenError {
	if o == nil {
		return nil
	}
	return o.GenericTokenError
}
