// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

// WafExclusionDataExclusionType - The type of exclusion.
type WafExclusionDataExclusionType string

const (
	WafExclusionDataExclusionTypeRule     WafExclusionDataExclusionType = "rule"
	WafExclusionDataExclusionTypeVariable WafExclusionDataExclusionType = "variable"
	WafExclusionDataExclusionTypeWaf      WafExclusionDataExclusionType = "waf"
)

func (e WafExclusionDataExclusionType) ToPointer() *WafExclusionDataExclusionType {
	return &e
}

func (e *WafExclusionDataExclusionType) UnmarshalJSON(data []byte) error {
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
		*e = WafExclusionDataExclusionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WafExclusionDataExclusionType: %v", v)
	}
}

// WafExclusionDataVariable - The variable to exclude. An optional selector can be specified after the variable separated by a colon (`:`) to restrict the variable to a particular parameter. Required for `exclusion_type=variable`.
type WafExclusionDataVariable string

const (
	WafExclusionDataVariableReqCookies      WafExclusionDataVariable = "req.cookies"
	WafExclusionDataVariableReqHeaders      WafExclusionDataVariable = "req.headers"
	WafExclusionDataVariableReqPost         WafExclusionDataVariable = "req.post"
	WafExclusionDataVariableReqPostFilename WafExclusionDataVariable = "req.post_filename"
	WafExclusionDataVariableReqQs           WafExclusionDataVariable = "req.qs"
)

func (e WafExclusionDataVariable) ToPointer() *WafExclusionDataVariable {
	return &e
}

func (e *WafExclusionDataVariable) UnmarshalJSON(data []byte) error {
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
		*e = WafExclusionDataVariable(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WafExclusionDataVariable: %v", v)
	}
}

type WafExclusionDataAttributes struct {
	// A conditional expression in VCL used to determine if the condition is met.
	Condition *string `json:"condition,omitempty"`
	// The type of exclusion.
	ExclusionType *WafExclusionDataExclusionType `json:"exclusion_type,omitempty"`
	// Whether to generate a log upon matching.
	Logging *bool `default:"true" json:"logging"`
	// Name of the exclusion.
	Name *string `json:"name,omitempty"`
	// A numeric ID identifying a WAF exclusion.
	Number *int64 `json:"number,omitempty"`
	// The variable to exclude. An optional selector can be specified after the variable separated by a colon (`:`) to restrict the variable to a particular parameter. Required for `exclusion_type=variable`.
	Variable *WafExclusionDataVariable `json:"variable,omitempty"`
}

func (w WafExclusionDataAttributes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafExclusionDataAttributes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafExclusionDataAttributes) GetCondition() *string {
	if o == nil {
		return nil
	}
	return o.Condition
}

func (o *WafExclusionDataAttributes) GetExclusionType() *WafExclusionDataExclusionType {
	if o == nil {
		return nil
	}
	return o.ExclusionType
}

func (o *WafExclusionDataAttributes) GetLogging() *bool {
	if o == nil {
		return nil
	}
	return o.Logging
}

func (o *WafExclusionDataAttributes) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *WafExclusionDataAttributes) GetNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *WafExclusionDataAttributes) GetVariable() *WafExclusionDataVariable {
	if o == nil {
		return nil
	}
	return o.Variable
}

type WafExclusionData struct {
	Attributes    *WafExclusionDataAttributes   `json:"attributes,omitempty"`
	Relationships *RelationshipsForWafExclusion `json:"relationships,omitempty"`
	// Resource type.
	Type *TypeWafExclusion `default:"waf_exclusion" json:"type"`
}

func (w WafExclusionData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafExclusionData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafExclusionData) GetAttributes() *WafExclusionDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *WafExclusionData) GetRelationships() *RelationshipsForWafExclusion {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *WafExclusionData) GetType() *TypeWafExclusion {
	if o == nil {
		return nil
	}
	return o.Type
}
