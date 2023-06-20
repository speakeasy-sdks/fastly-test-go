// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TypeWafFirewall - Resource type.
type TypeWafFirewall string

const (
	TypeWafFirewallWafFirewall TypeWafFirewall = "waf_firewall"
)

func (e TypeWafFirewall) ToPointer() *TypeWafFirewall {
	return &e
}

func (e *TypeWafFirewall) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "waf_firewall":
		*e = TypeWafFirewall(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TypeWafFirewall: %v", v)
	}
}
