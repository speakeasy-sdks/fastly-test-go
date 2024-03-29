// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"time"
)

// ExclusionType - The type of exclusion.
type ExclusionType string

const (
	ExclusionTypeRule     ExclusionType = "rule"
	ExclusionTypeVariable ExclusionType = "variable"
	ExclusionTypeWaf      ExclusionType = "waf"
)

func (e ExclusionType) ToPointer() *ExclusionType {
	return &e
}

func (e *ExclusionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "rule":
		fallthrough
	case "variable":
		fallthrough
	case "waf":
		*e = ExclusionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExclusionType: %v", v)
	}
}

// Variable - The variable to exclude. An optional selector can be specified after the variable separated by a colon (`:`) to restrict the variable to a particular parameter. Required for `exclusion_type=variable`.
type Variable string

const (
	VariableReqCookies      Variable = "req.cookies"
	VariableReqHeaders      Variable = "req.headers"
	VariableReqPost         Variable = "req.post"
	VariableReqPostFilename Variable = "req.post_filename"
	VariableReqQs           Variable = "req.qs"
)

func (e Variable) ToPointer() *Variable {
	return &e
}

func (e *Variable) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "req.cookies":
		fallthrough
	case "req.headers":
		fallthrough
	case "req.post":
		fallthrough
	case "req.post_filename":
		fallthrough
	case "req.qs":
		*e = Variable(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Variable: %v", v)
	}
}

type WafExclusionResponseDataAttributes struct {
	// A conditional expression in VCL used to determine if the condition is met.
	Condition *string `json:"condition,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The type of exclusion.
	ExclusionType *ExclusionType `json:"exclusion_type,omitempty"`
	// Whether to generate a log upon matching.
	Logging *bool `default:"true" json:"logging"`
	// Name of the exclusion.
	Name *string `json:"name,omitempty"`
	// A numeric ID identifying a WAF exclusion.
	Number *int64 `json:"number,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The variable to exclude. An optional selector can be specified after the variable separated by a colon (`:`) to restrict the variable to a particular parameter. Required for `exclusion_type=variable`.
	Variable *Variable `json:"variable,omitempty"`
}

func (w WafExclusionResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafExclusionResponseDataAttributes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafExclusionResponseDataAttributes) GetCondition() *string {
	if o == nil {
		return nil
	}
	return o.Condition
}

func (o *WafExclusionResponseDataAttributes) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *WafExclusionResponseDataAttributes) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *WafExclusionResponseDataAttributes) GetExclusionType() *ExclusionType {
	if o == nil {
		return nil
	}
	return o.ExclusionType
}

func (o *WafExclusionResponseDataAttributes) GetLogging() *bool {
	if o == nil {
		return nil
	}
	return o.Logging
}

func (o *WafExclusionResponseDataAttributes) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *WafExclusionResponseDataAttributes) GetNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *WafExclusionResponseDataAttributes) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *WafExclusionResponseDataAttributes) GetVariable() *Variable {
	if o == nil {
		return nil
	}
	return o.Variable
}
