// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type GetCacheSettingsRequest struct {
	// Name for the cache settings object.
	CacheSettingsName string `pathParam:"style=simple,explode=false,name=cache_settings_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetCacheSettingsRequest) GetCacheSettingsName() string {
	if o == nil {
		return ""
	}
	return o.CacheSettingsName
}

func (o *GetCacheSettingsRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetCacheSettingsRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetCacheSettingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	CacheSettingResponse *components.CacheSettingResponse
}

func (o *GetCacheSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCacheSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCacheSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCacheSettingsResponse) GetCacheSettingResponse() *components.CacheSettingResponse {
	if o == nil {
		return nil
	}
	return o.CacheSettingResponse
}
