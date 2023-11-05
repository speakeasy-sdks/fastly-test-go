// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"encoding/json"
	"fmt"
)

// ConditionType - Type of the condition. Required.
type ConditionType string

const (
	ConditionTypeRequest  ConditionType = "REQUEST"
	ConditionTypeCache    ConditionType = "CACHE"
	ConditionTypeResponse ConditionType = "RESPONSE"
	ConditionTypePrefetch ConditionType = "PREFETCH"
)

func (e ConditionType) ToPointer() *ConditionType {
	return &e
}

func (e *ConditionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "REQUEST":
		fallthrough
	case "CACHE":
		fallthrough
	case "RESPONSE":
		fallthrough
	case "PREFETCH":
		*e = ConditionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConditionType: %v", v)
	}
}

type ConditionInput struct {
	// A freeform descriptive note.
	Comment *string `form:"name=comment"`
	// Name of the condition. Required.
	Name *string `form:"name=name"`
	// A numeric string. Priority determines execution order. Lower numbers execute first.
	Priority *string `default:"100" form:"name=priority"`
	// A conditional expression in VCL used to determine if the condition is met.
	Statement *string `form:"name=statement"`
	// Type of the condition. Required.
	Type *ConditionType `form:"name=type"`
	// A numeric string that represents the service version.
	Version *string `form:"name=version"`
}

func (c ConditionInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConditionInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConditionInput) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *ConditionInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConditionInput) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *ConditionInput) GetStatement() *string {
	if o == nil {
		return nil
	}
	return o.Statement
}

func (o *ConditionInput) GetType() *ConditionType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConditionInput) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
