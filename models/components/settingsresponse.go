// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type SettingsResponse struct {
	// The default host name for the version.
	GeneralDefaultHost *string `json:"general.default_host,omitempty"`
	// The default time-to-live (TTL) for the version.
	GeneralDefaultTTL *int64 `json:"general.default_ttl,omitempty"`
	// Enables serving a stale object if there is an error.
	GeneralStaleIfError *bool `default:"false" json:"general.stale_if_error"`
	// The default time-to-live (TTL) for serving the stale object for the version.
	GeneralStaleIfErrorTTL *int64  `default:"43200" json:"general.stale_if_error_ttl"`
	ServiceID              *string `json:"service_id,omitempty"`
	Version                *int64  `json:"version,omitempty"`
}

func (s SettingsResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SettingsResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SettingsResponse) GetGeneralDefaultHost() *string {
	if o == nil {
		return nil
	}
	return o.GeneralDefaultHost
}

func (o *SettingsResponse) GetGeneralDefaultTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.GeneralDefaultTTL
}

func (o *SettingsResponse) GetGeneralStaleIfError() *bool {
	if o == nil {
		return nil
	}
	return o.GeneralStaleIfError
}

func (o *SettingsResponse) GetGeneralStaleIfErrorTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.GeneralStaleIfErrorTTL
}

func (o *SettingsResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *SettingsResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
