// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type GetDictionaryItemRequest struct {
	// Alphanumeric string identifying a Dictionary.
	DictionaryID string `pathParam:"style=simple,explode=false,name=dictionary_id"`
	// Item key, maximum 256 characters.
	DictionaryItemKey string `pathParam:"style=simple,explode=false,name=dictionary_item_key"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *GetDictionaryItemRequest) GetDictionaryID() string {
	if o == nil {
		return ""
	}
	return o.DictionaryID
}

func (o *GetDictionaryItemRequest) GetDictionaryItemKey() string {
	if o == nil {
		return ""
	}
	return o.DictionaryItemKey
}

func (o *GetDictionaryItemRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type GetDictionaryItemResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DictionaryItemResponse *components.DictionaryItemResponse
}

func (o *GetDictionaryItemResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDictionaryItemResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDictionaryItemResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDictionaryItemResponse) GetDictionaryItemResponse() *components.DictionaryItemResponse {
	if o == nil {
		return nil
	}
	return o.DictionaryItemResponse
}
