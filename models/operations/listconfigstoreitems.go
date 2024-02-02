// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type ListConfigStoreItemsRequest struct {
	// An alphanumeric string identifying the config store.
	ConfigStoreID string `pathParam:"style=simple,explode=false,name=config_store_id"`
}

func (o *ListConfigStoreItemsRequest) GetConfigStoreID() string {
	if o == nil {
		return ""
	}
	return o.ConfigStoreID
}

type ListConfigStoreItemsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Classes []components.ConfigStoreItemResponse
}

func (o *ListConfigStoreItemsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListConfigStoreItemsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListConfigStoreItemsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListConfigStoreItemsResponse) GetClasses() []components.ConfigStoreItemResponse {
	if o == nil {
		return nil
	}
	return o.Classes
}
