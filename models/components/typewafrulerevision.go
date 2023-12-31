// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// TypeWafRuleRevision - Resource type.
type TypeWafRuleRevision string

const (
	TypeWafRuleRevisionWafRuleRevision TypeWafRuleRevision = "waf_rule_revision"
)

func (e TypeWafRuleRevision) ToPointer() *TypeWafRuleRevision {
	return &e
}

func (e *TypeWafRuleRevision) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "waf_rule_revision":
		*e = TypeWafRuleRevision(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TypeWafRuleRevision: %v", v)
	}
}
