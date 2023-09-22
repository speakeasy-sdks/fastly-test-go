// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RequestSettingsAction - Allows you to terminate request handling and immediately perform an action.
type RequestSettingsAction string

const (
	RequestSettingsActionLookup RequestSettingsAction = "lookup"
	RequestSettingsActionPass   RequestSettingsAction = "pass"
)

func (e RequestSettingsAction) ToPointer() *RequestSettingsAction {
	return &e
}

func (e *RequestSettingsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "lookup":
		fallthrough
	case "pass":
		*e = RequestSettingsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestSettingsAction: %v", v)
	}
}

// RequestSettingsXff - Short for X-Forwarded-For.
type RequestSettingsXff string

const (
	RequestSettingsXffClear     RequestSettingsXff = "clear"
	RequestSettingsXffLeave     RequestSettingsXff = "leave"
	RequestSettingsXffAppend    RequestSettingsXff = "append"
	RequestSettingsXffAppendAll RequestSettingsXff = "append_all"
	RequestSettingsXffOverwrite RequestSettingsXff = "overwrite"
)

func (e RequestSettingsXff) ToPointer() *RequestSettingsXff {
	return &e
}

func (e *RequestSettingsXff) UnmarshalJSON(data []byte) error {
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
		*e = RequestSettingsXff(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestSettingsXff: %v", v)
	}
}

type RequestSettings2 struct {
	// Allows you to terminate request handling and immediately perform an action.
	Action *RequestSettingsAction `form:"name=action"`
	// Disable collapsed forwarding, so you don't wait for other objects to origin.
	BypassBusyWait *int64 `form:"name=bypass_busy_wait"`
	// Sets the host header.
	DefaultHost *string `form:"name=default_host"`
	// Allows you to force a cache miss for the request. Replaces the item in the cache if the content is cacheable.
	ForceMiss *int64 `form:"name=force_miss"`
	// Forces the request use SSL (redirects a non-SSL to SSL).
	ForceSsl *int64 `form:"name=force_ssl"`
	// Injects Fastly-Geo-Country, Fastly-Geo-City, and Fastly-Geo-Region into the request headers.
	GeoHeaders *int64 `form:"name=geo_headers"`
	// Comma separated list of varnish request object fields that should be in the hash key.
	HashKeys *string `form:"name=hash_keys"`
	// How old an object is allowed to be to serve stale-if-error or stale-while-revalidate.
	MaxStaleAge *int64 `form:"name=max_stale_age"`
	// Name for the request settings.
	Name *string `form:"name=name"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition *string `form:"name=request_condition"`
	// Injects the X-Timer info into the request for viewing origin fetch durations.
	TimerSupport *int64 `form:"name=timer_support"`
	// Short for X-Forwarded-For.
	Xff *RequestSettingsXff `form:"name=xff"`
}

func (o *RequestSettings2) GetAction() *RequestSettingsAction {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *RequestSettings2) GetBypassBusyWait() *int64 {
	if o == nil {
		return nil
	}
	return o.BypassBusyWait
}

func (o *RequestSettings2) GetDefaultHost() *string {
	if o == nil {
		return nil
	}
	return o.DefaultHost
}

func (o *RequestSettings2) GetForceMiss() *int64 {
	if o == nil {
		return nil
	}
	return o.ForceMiss
}

func (o *RequestSettings2) GetForceSsl() *int64 {
	if o == nil {
		return nil
	}
	return o.ForceSsl
}

func (o *RequestSettings2) GetGeoHeaders() *int64 {
	if o == nil {
		return nil
	}
	return o.GeoHeaders
}

func (o *RequestSettings2) GetHashKeys() *string {
	if o == nil {
		return nil
	}
	return o.HashKeys
}

func (o *RequestSettings2) GetMaxStaleAge() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxStaleAge
}

func (o *RequestSettings2) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *RequestSettings2) GetRequestCondition() *string {
	if o == nil {
		return nil
	}
	return o.RequestCondition
}

func (o *RequestSettings2) GetTimerSupport() *int64 {
	if o == nil {
		return nil
	}
	return o.TimerSupport
}

func (o *RequestSettings2) GetXff() *RequestSettingsXff {
	if o == nil {
		return nil
	}
	return o.Xff
}
