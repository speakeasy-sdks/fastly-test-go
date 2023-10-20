// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateConfigStoreRequest struct {
	ConfigStore *shared.ConfigStore `request:"mediaType=application/x-www-form-urlencoded"`
	// An alphanumeric string identifying the config store.
	ConfigStoreID string `pathParam:"style=simple,explode=false,name=config_store_id"`
}

func (o *UpdateConfigStoreRequest) GetConfigStore() *shared.ConfigStore {
	if o == nil {
		return nil
	}
	return o.ConfigStore
}

func (o *UpdateConfigStoreRequest) GetConfigStoreID() string {
	if o == nil {
		return ""
	}
	return o.ConfigStoreID
}

type UpdateConfigStoreResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ConfigStoreResponse *shared.ConfigStoreResponse
}

func (o *UpdateConfigStoreResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConfigStoreResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConfigStoreResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateConfigStoreResponse) GetConfigStoreResponse() *shared.ConfigStoreResponse {
	if o == nil {
		return nil
	}
	return o.ConfigStoreResponse
}
