// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// WafExclusionDataAttributesExclusionType - The type of exclusion.
type WafExclusionDataAttributesExclusionType string

const (
	WafExclusionDataAttributesExclusionTypeRule     WafExclusionDataAttributesExclusionType = "rule"
	WafExclusionDataAttributesExclusionTypeVariable WafExclusionDataAttributesExclusionType = "variable"
	WafExclusionDataAttributesExclusionTypeWaf      WafExclusionDataAttributesExclusionType = "waf"
)

func (e WafExclusionDataAttributesExclusionType) ToPointer() *WafExclusionDataAttributesExclusionType {
	return &e
}

func (e *WafExclusionDataAttributesExclusionType) UnmarshalJSON(data []byte) error {
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
		*e = WafExclusionDataAttributesExclusionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WafExclusionDataAttributesExclusionType: %v", v)
	}
}

// WafExclusionDataAttributesVariable - The variable to exclude. An optional selector can be specified after the variable separated by a colon (`:`) to restrict the variable to a particular parameter. Required for `exclusion_type=variable`.
type WafExclusionDataAttributesVariable string

const (
	WafExclusionDataAttributesVariableReqCookies             WafExclusionDataAttributesVariable = "req.cookies"
	WafExclusionDataAttributesVariableReqHeaders             WafExclusionDataAttributesVariable = "req.headers"
	WafExclusionDataAttributesVariableReqPost                WafExclusionDataAttributesVariable = "req.post"
	WafExclusionDataAttributesVariableReqPostFilename        WafExclusionDataAttributesVariable = "req.post_filename"
	WafExclusionDataAttributesVariableReqQs                  WafExclusionDataAttributesVariable = "req.qs"
	WafExclusionDataAttributesVariableLessThanNilGreaterThan WafExclusionDataAttributesVariable = "<nil>"
)

func (e WafExclusionDataAttributesVariable) ToPointer() *WafExclusionDataAttributesVariable {
	return &e
}

func (e *WafExclusionDataAttributesVariable) UnmarshalJSON(data []byte) error {
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
		fallthrough
	case "<nil>":
		*e = WafExclusionDataAttributesVariable(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WafExclusionDataAttributesVariable: %v", v)
	}
}

type WafExclusionDataAttributes struct {
	// A conditional expression in VCL used to determine if the condition is met.
	Condition *string `json:"condition,omitempty"`
	// The type of exclusion.
	ExclusionType *WafExclusionDataAttributesExclusionType `json:"exclusion_type,omitempty"`
	// Whether to generate a log upon matching.
	Logging *bool `json:"logging,omitempty"`
	// Name of the exclusion.
	Name *string `json:"name,omitempty"`
	// A numeric ID identifying a WAF exclusion.
	Number *int64 `json:"number,omitempty"`
	// The variable to exclude. An optional selector can be specified after the variable separated by a colon (`:`) to restrict the variable to a particular parameter. Required for `exclusion_type=variable`.
	Variable *WafExclusionDataAttributesVariable `json:"variable,omitempty"`
}

type WafExclusionDataInput struct {
	Attributes    *WafExclusionDataAttributes `json:"attributes,omitempty"`
	Relationships interface{}                 `json:"relationships,omitempty"`
	// Resource type.
	Type *TypeWafExclusion `json:"type,omitempty"`
}
