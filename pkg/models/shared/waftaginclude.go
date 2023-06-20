// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// WafTagInclude - Include relationships. Optional.
type WafTagInclude string

const (
	WafTagIncludeWafRules WafTagInclude = "waf_rules"
)

func (e WafTagInclude) ToPointer() *WafTagInclude {
	return &e
}

func (e *WafTagInclude) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "waf_rules":
		*e = WafTagInclude(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WafTagInclude: %v", v)
	}
}
