// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"encoding/json"
	"fmt"
)

// HeaderAction - Accepts a string value.
type HeaderAction string

const (
	HeaderActionSet         HeaderAction = "set"
	HeaderActionAppend      HeaderAction = "append"
	HeaderActionDelete      HeaderAction = "delete"
	HeaderActionRegex       HeaderAction = "regex"
	HeaderActionRegexRepeat HeaderAction = "regex_repeat"
)

func (e HeaderAction) ToPointer() *HeaderAction {
	return &e
}

func (e *HeaderAction) UnmarshalJSON(data []byte) error {
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
		*e = HeaderAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HeaderAction: %v", v)
	}
}

// HeaderType - Accepts a string value.
type HeaderType string

const (
	HeaderTypeRequest  HeaderType = "request"
	HeaderTypeCache    HeaderType = "cache"
	HeaderTypeResponse HeaderType = "response"
)

func (e HeaderType) ToPointer() *HeaderType {
	return &e
}

func (e *HeaderType) UnmarshalJSON(data []byte) error {
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
		*e = HeaderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HeaderType: %v", v)
	}
}

type Header struct {
	// Accepts a string value.
	Action *HeaderAction `form:"name=action"`
	// Name of the cache condition controlling when this configuration applies.
	CacheCondition *string `form:"name=cache_condition"`
	// Header to set.
	Dst *string `form:"name=dst"`
	// Don't add the header if it is added already. Only applies to 'set' action.
	IgnoreIfSet *int64 `form:"name=ignore_if_set"`
	// A handle to refer to this Header object.
	Name *string `form:"name=name"`
	// Priority determines execution order. Lower numbers execute first.
	Priority *int64 `default:"100" form:"name=priority"`
	// Regular expression to use. Only applies to `regex` and `regex_repeat` actions.
	Regex *string `form:"name=regex"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition *string `form:"name=request_condition"`
	// Optional name of a response condition to apply.
	ResponseCondition *string `form:"name=response_condition"`
	// Variable to be used as a source for the header content. Does not apply to `delete` action.
	Src *string `form:"name=src"`
	// Value to substitute in place of regular expression. Only applies to `regex` and `regex_repeat` actions.
	Substitution *string `form:"name=substitution"`
	// Accepts a string value.
	Type *HeaderType `form:"name=type"`
}

func (h Header) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *Header) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Header) GetAction() *HeaderAction {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *Header) GetCacheCondition() *string {
	if o == nil {
		return nil
	}
	return o.CacheCondition
}

func (o *Header) GetDst() *string {
	if o == nil {
		return nil
	}
	return o.Dst
}

func (o *Header) GetIgnoreIfSet() *int64 {
	if o == nil {
		return nil
	}
	return o.IgnoreIfSet
}

func (o *Header) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Header) GetPriority() *int64 {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *Header) GetRegex() *string {
	if o == nil {
		return nil
	}
	return o.Regex
}

func (o *Header) GetRequestCondition() *string {
	if o == nil {
		return nil
	}
	return o.RequestCondition
}

func (o *Header) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *Header) GetSrc() *string {
	if o == nil {
		return nil
	}
	return o.Src
}

func (o *Header) GetSubstitution() *string {
	if o == nil {
		return nil
	}
	return o.Substitution
}

func (o *Header) GetType() *HeaderType {
	if o == nil {
		return nil
	}
	return o.Type
}