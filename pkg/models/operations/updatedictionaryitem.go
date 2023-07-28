// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateDictionaryItemSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *UpdateDictionaryItemSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type UpdateDictionaryItemRequest struct {
	// Alphanumeric string identifying a Dictionary.
	DictionaryID   string                 `pathParam:"style=simple,explode=false,name=dictionary_id"`
	DictionaryItem *shared.DictionaryItem `request:"mediaType=application/x-www-form-urlencoded"`
	// Item key, maximum 256 characters.
	DictionaryItemKey string `pathParam:"style=simple,explode=false,name=dictionary_item_key"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *UpdateDictionaryItemRequest) GetDictionaryID() string {
	if o == nil {
		return ""
	}
	return o.DictionaryID
}

func (o *UpdateDictionaryItemRequest) GetDictionaryItem() *shared.DictionaryItem {
	if o == nil {
		return nil
	}
	return o.DictionaryItem
}

func (o *UpdateDictionaryItemRequest) GetDictionaryItemKey() string {
	if o == nil {
		return ""
	}
	return o.DictionaryItemKey
}

func (o *UpdateDictionaryItemRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type UpdateDictionaryItemResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DictionaryItemResponse *shared.DictionaryItemResponse
}

func (o *UpdateDictionaryItemResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDictionaryItemResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDictionaryItemResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateDictionaryItemResponse) GetDictionaryItemResponse() *shared.DictionaryItemResponse {
	if o == nil {
		return nil
	}
	return o.DictionaryItemResponse
}
