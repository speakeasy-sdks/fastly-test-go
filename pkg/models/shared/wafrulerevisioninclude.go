// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// WafRuleRevisionInclude - Include relationships. Optional.
type WafRuleRevisionInclude string

const (
	WafRuleRevisionIncludeWafRule WafRuleRevisionInclude = "waf_rule"
)

func (e WafRuleRevisionInclude) ToPointer() *WafRuleRevisionInclude {
	return &e
}

func (e *WafRuleRevisionInclude) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "waf_rule":
		*e = WafRuleRevisionInclude(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WafRuleRevisionInclude: %v", v)
	}
}
