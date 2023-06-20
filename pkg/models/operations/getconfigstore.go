// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetConfigStoreSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type GetConfigStoreRequest struct {
	// An alphanumeric string identifying the config store.
	ConfigStoreID string `pathParam:"style=simple,explode=false,name=config_store_id"`
}

type GetConfigStoreResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ConfigStoreResponse *shared.ConfigStoreResponse
}
