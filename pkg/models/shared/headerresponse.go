// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// HeaderResponseAction - Accepts a string value.
type HeaderResponseAction string

const (
	HeaderResponseActionSet         HeaderResponseAction = "set"
	HeaderResponseActionAppend      HeaderResponseAction = "append"
	HeaderResponseActionDelete      HeaderResponseAction = "delete"
	HeaderResponseActionRegex       HeaderResponseAction = "regex"
	HeaderResponseActionRegexRepeat HeaderResponseAction = "regex_repeat"
)

func (e HeaderResponseAction) ToPointer() *HeaderResponseAction {
	return &e
}

func (e *HeaderResponseAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "set":
		fallthrough
	case "append":
		fallthrough
	case "delete":
		fallthrough
	case "regex":
		fallthrough
	case "regex_repeat":
		*e = HeaderResponseAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HeaderResponseAction: %v", v)
	}
}

// HeaderResponseType - Accepts a string value.
type HeaderResponseType string

const (
	HeaderResponseTypeRequest  HeaderResponseType = "request"
	HeaderResponseTypeCache    HeaderResponseType = "cache"
	HeaderResponseTypeResponse HeaderResponseType = "response"
)

func (e HeaderResponseType) ToPointer() *HeaderResponseType {
	return &e
}

func (e *HeaderResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "request":
		fallthrough
	case "cache":
		fallthrough
	case "response":
		*e = HeaderResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HeaderResponseType: %v", v)
	}
}

// HeaderResponse - OK
type HeaderResponse struct {
	// Accepts a string value.
	Action *HeaderResponseAction `json:"action,omitempty"`
	// Name of the cache condition controlling when this configuration applies.
	CacheCondition *string `json:"cache_condition,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Header to set.
	Dst *string `json:"dst,omitempty"`
	// Don't add the header if it is added already. Only applies to 'set' action.
	IgnoreIfSet *int64 `json:"ignore_if_set,omitempty"`
	// A handle to refer to this Header object.
	Name *string `json:"name,omitempty"`
	// Priority determines execution order. Lower numbers execute first.
	Priority *int64 `json:"priority,omitempty"`
	// Regular expression to use. Only applies to `regex` and `regex_repeat` actions.
	Regex *string `json:"regex,omitempty"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition *string `json:"request_condition,omitempty"`
	// Optional name of a response condition to apply.
	ResponseCondition *string `json:"response_condition,omitempty"`
	ServiceID         *string `json:"service_id,omitempty"`
	// Variable to be used as a source for the header content. Does not apply to `delete` action.
	Src *string `json:"src,omitempty"`
	// Value to substitute in place of regular expression. Only applies to `regex` and `regex_repeat` actions.
	Substitution *string `json:"substitution,omitempty"`
	// Accepts a string value.
	Type *HeaderResponseType `json:"type,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}

func (o *HeaderResponse) GetAction() *HeaderResponseAction {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *HeaderResponse) GetCacheCondition() *string {
	if o == nil {
		return nil
	}
	return o.CacheCondition
}

func (o *HeaderResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *HeaderResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *HeaderResponse) GetDst() *string {
	if o == nil {
		return nil
	}
	return o.Dst
}

func (o *HeaderResponse) GetIgnoreIfSet() *int64 {
	if o == nil {
		return nil
	}
	return o.IgnoreIfSet
}

func (o *HeaderResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *HeaderResponse) GetPriority() *int64 {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *HeaderResponse) GetRegex() *string {
	if o == nil {
		return nil
	}
	return o.Regex
}

func (o *HeaderResponse) GetRequestCondition() *string {
	if o == nil {
		return nil
	}
	return o.RequestCondition
}

func (o *HeaderResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *HeaderResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *HeaderResponse) GetSrc() *string {
	if o == nil {
		return nil
	}
	return o.Src
}

func (o *HeaderResponse) GetSubstitution() *string {
	if o == nil {
		return nil
	}
	return o.Substitution
}

func (o *HeaderResponse) GetType() *HeaderResponseType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *HeaderResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *HeaderResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
