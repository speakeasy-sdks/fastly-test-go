// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"encoding/json"
	"fmt"
	"time"
)

// CacheSettingResponseAction - If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.
type CacheSettingResponseAction string

const (
	CacheSettingResponseActionPass    CacheSettingResponseAction = "pass"
	CacheSettingResponseActionCache   CacheSettingResponseAction = "cache"
	CacheSettingResponseActionRestart CacheSettingResponseAction = "restart"
)

func (e CacheSettingResponseAction) ToPointer() *CacheSettingResponseAction {
	return &e
}

func (e *CacheSettingResponseAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pass":
		fallthrough
	case "cache":
		fallthrough
	case "restart":
		*e = CacheSettingResponseAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CacheSettingResponseAction: %v", v)
	}
}

type CacheSettingResponse struct {
	// If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.
	//
	Action *CacheSettingResponseAction `json:"action,omitempty"`
	// Name of the cache condition controlling when this configuration applies.
	CacheCondition *string `json:"cache_condition,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Name for the cache settings object.
	Name      *string `json:"name,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	// Maximum time in seconds to continue to use a stale version of the object if future requests to your backend server fail (also known as 'stale if error').
	StaleTTL *int64 `json:"stale_ttl,omitempty"`
	// Maximum time to consider the object fresh in the cache (the cache 'time to live').
	TTL *int64 `json:"ttl,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}

func (c CacheSettingResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CacheSettingResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CacheSettingResponse) GetAction() *CacheSettingResponseAction {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *CacheSettingResponse) GetCacheCondition() *string {
	if o == nil {
		return nil
	}
	return o.CacheCondition
}

func (o *CacheSettingResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CacheSettingResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *CacheSettingResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CacheSettingResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *CacheSettingResponse) GetStaleTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.StaleTTL
}

func (o *CacheSettingResponse) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *CacheSettingResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CacheSettingResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
