// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetConfigStoreInfoRequest struct {
	// An alphanumeric string identifying the config store.
	ConfigStoreID string `pathParam:"style=simple,explode=false,name=config_store_id"`
}

func (o *GetConfigStoreInfoRequest) GetConfigStoreID() string {
	if o == nil {
		return ""
	}
	return o.ConfigStoreID
}

type GetConfigStoreInfoResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ConfigStoreInfoResponse *shared.ConfigStoreInfoResponse
}

func (o *GetConfigStoreInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConfigStoreInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConfigStoreInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetConfigStoreInfoResponse) GetConfigStoreInfoResponse() *shared.ConfigStoreInfoResponse {
	if o == nil {
		return nil
	}
	return o.ConfigStoreInfoResponse
}
