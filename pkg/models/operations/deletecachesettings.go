// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteCacheSettingsSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteCacheSettingsRequest struct {
	// Name for the cache settings object.
	CacheSettingsName string `pathParam:"style=simple,explode=false,name=cache_settings_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

// DeleteCacheSettings200ApplicationJSON - OK
type DeleteCacheSettings200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteCacheSettingsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteCacheSettings200ApplicationJSONObject *DeleteCacheSettings200ApplicationJSON
}
