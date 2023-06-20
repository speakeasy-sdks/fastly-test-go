// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateCacheSettingsSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type UpdateCacheSettingsRequest struct {
	CacheSetting1 *shared.CacheSetting1 `request:"mediaType=application/x-www-form-urlencoded"`
	// Name for the cache settings object.
	CacheSettingsName string `pathParam:"style=simple,explode=false,name=cache_settings_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

type UpdateCacheSettingsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	CacheSettingResponse *shared.CacheSettingResponse
}
