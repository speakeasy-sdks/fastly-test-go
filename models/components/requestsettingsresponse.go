// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"encoding/json"
	"fmt"
	"time"
)

// RequestSettingsResponseAction - Allows you to terminate request handling and immediately perform an action.
type RequestSettingsResponseAction string

const (
	RequestSettingsResponseActionLookup RequestSettingsResponseAction = "lookup"
	RequestSettingsResponseActionPass   RequestSettingsResponseAction = "pass"
)

func (e RequestSettingsResponseAction) ToPointer() *RequestSettingsResponseAction {
	return &e
}

func (e *RequestSettingsResponseAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "lookup":
		fallthrough
	case "pass":
		*e = RequestSettingsResponseAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestSettingsResponseAction: %v", v)
	}
}

// RequestSettingsResponseXff - Short for X-Forwarded-For.
type RequestSettingsResponseXff string

const (
	RequestSettingsResponseXffClear     RequestSettingsResponseXff = "clear"
	RequestSettingsResponseXffLeave     RequestSettingsResponseXff = "leave"
	RequestSettingsResponseXffAppend    RequestSettingsResponseXff = "append"
	RequestSettingsResponseXffAppendAll RequestSettingsResponseXff = "append_all"
	RequestSettingsResponseXffOverwrite RequestSettingsResponseXff = "overwrite"
)

func (e RequestSettingsResponseXff) ToPointer() *RequestSettingsResponseXff {
	return &e
}

func (e *RequestSettingsResponseXff) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "clear":
		fallthrough
	case "leave":
		fallthrough
	case "append":
		fallthrough
	case "append_all":
		fallthrough
	case "overwrite":
		*e = RequestSettingsResponseXff(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestSettingsResponseXff: %v", v)
	}
}

type RequestSettingsResponse struct {
	// Allows you to terminate request handling and immediately perform an action.
	Action *RequestSettingsResponseAction `json:"action,omitempty"`
	// Disable collapsed forwarding, so you don't wait for other objects to origin.
	BypassBusyWait *int64 `json:"bypass_busy_wait,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Sets the host header.
	DefaultHost *string `json:"default_host,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Allows you to force a cache miss for the request. Replaces the item in the cache if the content is cacheable.
	ForceMiss *int64 `json:"force_miss,omitempty"`
	// Forces the request use SSL (redirects a non-SSL to SSL).
	ForceSsl *int64 `json:"force_ssl,omitempty"`
	// Injects Fastly-Geo-Country, Fastly-Geo-City, and Fastly-Geo-Region into the request headers.
	GeoHeaders *int64 `json:"geo_headers,omitempty"`
	// Comma separated list of varnish request object fields that should be in the hash key.
	HashKeys *string `json:"hash_keys,omitempty"`
	// How old an object is allowed to be to serve stale-if-error or stale-while-revalidate.
	MaxStaleAge *int64 `json:"max_stale_age,omitempty"`
	// Name for the request settings.
	Name *string `json:"name,omitempty"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition *string `json:"request_condition,omitempty"`
	ServiceID        *string `json:"service_id,omitempty"`
	// Injects the X-Timer info into the request for viewing origin fetch durations.
	TimerSupport *int64 `json:"timer_support,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
	// Short for X-Forwarded-For.
	Xff *RequestSettingsResponseXff `json:"xff,omitempty"`
}

func (r RequestSettingsResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestSettingsResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestSettingsResponse) GetAction() *RequestSettingsResponseAction {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *RequestSettingsResponse) GetBypassBusyWait() *int64 {
	if o == nil {
		return nil
	}
	return o.BypassBusyWait
}

func (o *RequestSettingsResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RequestSettingsResponse) GetDefaultHost() *string {
	if o == nil {
		return nil
	}
	return o.DefaultHost
}

func (o *RequestSettingsResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *RequestSettingsResponse) GetForceMiss() *int64 {
	if o == nil {
		return nil
	}
	return o.ForceMiss
}

func (o *RequestSettingsResponse) GetForceSsl() *int64 {
	if o == nil {
		return nil
	}
	return o.ForceSsl
}

func (o *RequestSettingsResponse) GetGeoHeaders() *int64 {
	if o == nil {
		return nil
	}
	return o.GeoHeaders
}

func (o *RequestSettingsResponse) GetHashKeys() *string {
	if o == nil {
		return nil
	}
	return o.HashKeys
}

func (o *RequestSettingsResponse) GetMaxStaleAge() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxStaleAge
}

func (o *RequestSettingsResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *RequestSettingsResponse) GetRequestCondition() *string {
	if o == nil {
		return nil
	}
	return o.RequestCondition
}

func (o *RequestSettingsResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *RequestSettingsResponse) GetTimerSupport() *int64 {
	if o == nil {
		return nil
	}
	return o.TimerSupport
}

func (o *RequestSettingsResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RequestSettingsResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *RequestSettingsResponse) GetXff() *RequestSettingsResponseXff {
	if o == nil {
		return nil
	}
	return o.Xff
}
