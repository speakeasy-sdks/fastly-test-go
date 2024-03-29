// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateRequestSettingsRequest struct {
	RequestSettings *components.RequestSettings `request:"mediaType=application/x-www-form-urlencoded"`
	// Name for the request settings.
	RequestSettingsName string `pathParam:"style=simple,explode=false,name=request_settings_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateRequestSettingsRequest) GetRequestSettings() *components.RequestSettings {
	if o == nil {
		return nil
	}
	return o.RequestSettings
}

func (o *UpdateRequestSettingsRequest) GetRequestSettingsName() string {
	if o == nil {
		return ""
	}
	return o.RequestSettingsName
}

func (o *UpdateRequestSettingsRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateRequestSettingsRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateRequestSettingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	RequestSettingsResponse *components.RequestSettingsResponse
}

func (o *UpdateRequestSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRequestSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRequestSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRequestSettingsResponse) GetRequestSettingsResponse() *components.RequestSettingsResponse {
	if o == nil {
		return nil
	}
	return o.RequestSettingsResponse
}
