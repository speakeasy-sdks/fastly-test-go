// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteDictionaryItemRequest struct {
	// Alphanumeric string identifying a Dictionary.
	DictionaryID string `pathParam:"style=simple,explode=false,name=dictionary_id"`
	// Item key, maximum 256 characters.
	DictionaryItemKey string `pathParam:"style=simple,explode=false,name=dictionary_item_key"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *DeleteDictionaryItemRequest) GetDictionaryID() string {
	if o == nil {
		return ""
	}
	return o.DictionaryID
}

func (o *DeleteDictionaryItemRequest) GetDictionaryItemKey() string {
	if o == nil {
		return ""
	}
	return o.DictionaryItemKey
}

func (o *DeleteDictionaryItemRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

// DeleteDictionaryItemResponseBody - OK
type DeleteDictionaryItemResponseBody struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteDictionaryItemResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteDictionaryItemResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *DeleteDictionaryItemResponseBody
}

func (o *DeleteDictionaryItemResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteDictionaryItemResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteDictionaryItemResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteDictionaryItemResponse) GetObject() *DeleteDictionaryItemResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
