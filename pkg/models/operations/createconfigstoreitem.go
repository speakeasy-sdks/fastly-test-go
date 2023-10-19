// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateConfigStoreItemRequest struct {
	// An alphanumeric string identifying the config store.
	ConfigStoreID   string                  `pathParam:"style=simple,explode=false,name=config_store_id"`
	ConfigStoreItem *shared.ConfigStoreItem `request:"mediaType=application/x-www-form-urlencoded"`
}

func (o *CreateConfigStoreItemRequest) GetConfigStoreID() string {
	if o == nil {
		return ""
	}
	return o.ConfigStoreID
}

func (o *CreateConfigStoreItemRequest) GetConfigStoreItem() *shared.ConfigStoreItem {
	if o == nil {
		return nil
	}
	return o.ConfigStoreItem
}

type CreateConfigStoreItemResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ConfigStoreItemResponse *shared.ConfigStoreItemResponse
}

func (o *CreateConfigStoreItemResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateConfigStoreItemResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateConfigStoreItemResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateConfigStoreItemResponse) GetConfigStoreItemResponse() *shared.ConfigStoreItemResponse {
	if o == nil {
		return nil
	}
	return o.ConfigStoreItemResponse
}
