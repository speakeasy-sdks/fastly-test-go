// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type CreateStoreRequest struct {
	Location *string           `queryParam:"style=form,explode=true,name=location"`
	Store    *components.Store `request:"mediaType=application/json"`
}

func (o *CreateStoreRequest) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *CreateStoreRequest) GetStore() *components.Store {
	if o == nil {
		return nil
	}
	return o.Store
}

type CreateStoreResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	StoreResponse *components.StoreResponse
}

func (o *CreateStoreResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateStoreResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateStoreResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateStoreResponse) GetStoreResponse() *components.StoreResponse {
	if o == nil {
		return nil
	}
	return o.StoreResponse
}
