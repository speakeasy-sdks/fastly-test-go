// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteConfigStoreItemSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteConfigStoreItemRequest struct {
	// An alphanumeric string identifying the config store.
	ConfigStoreID string `pathParam:"style=simple,explode=false,name=config_store_id"`
	// Item key, maximum 256 characters.
	ConfigStoreItemKey string `pathParam:"style=simple,explode=false,name=config_store_item_key"`
}

// DeleteConfigStoreItem200ApplicationJSON - OK
type DeleteConfigStoreItem200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteConfigStoreItemResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteConfigStoreItem200ApplicationJSONObject *DeleteConfigStoreItem200ApplicationJSON
}
