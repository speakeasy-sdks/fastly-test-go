// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type GetStoreRequest struct {
	StoreID string `pathParam:"style=simple,explode=false,name=store_id"`
}

func (o *GetStoreRequest) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

type GetStoreResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	StoreResponse *components.StoreResponse
}

func (o *GetStoreResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetStoreResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetStoreResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetStoreResponse) GetStoreResponse() *components.StoreResponse {
	if o == nil {
		return nil
	}
	return o.StoreResponse
}