// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"encoding/json"
	"fmt"
)

// Type of the condition. Required.
type Type string

const (
	TypeRequest  Type = "REQUEST"
	TypeCache    Type = "CACHE"
	TypeResponse Type = "RESPONSE"
	TypePrefetch Type = "PREFETCH"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
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
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type Condition struct {
	// A freeform descriptive note.
	Comment *string `form:"name=comment"`
	// Name of the condition. Required.
	Name *string `form:"name=name"`
	// A numeric string. Priority determines execution order. Lower numbers execute first.
	Priority *string `default:"100" form:"name=priority"`
	// A conditional expression in VCL used to determine if the condition is met.
	Statement *string `form:"name=statement"`
	// Type of the condition. Required.
	Type *Type `form:"name=type"`
	// A numeric string that represents the service version.
	Version *string `form:"name=version"`
}

func (c Condition) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Condition) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Condition) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *Condition) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Condition) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *Condition) GetStatement() *string {
	if o == nil {
		return nil
	}
	return o.Statement
}

func (o *Condition) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Condition) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}