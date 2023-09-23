// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateTLSActivationResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Created
	TLSActivationResponse *shared.TLSActivationResponse
}

func (o *CreateTLSActivationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTLSActivationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTLSActivationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTLSActivationResponse) GetTLSActivationResponse() *shared.TLSActivationResponse {
	if o == nil {
		return nil
	}
	return o.TLSActivationResponse
}
