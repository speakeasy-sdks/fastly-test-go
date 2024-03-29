// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type RelationshipMemberWafFirewallVersionInput struct {
	// Resource type.
	Type *TypeWafFirewallVersion `default:"waf_firewall_version" json:"type"`
}

func (r RelationshipMemberWafFirewallVersionInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationshipMemberWafFirewallVersionInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RelationshipMemberWafFirewallVersionInput) GetType() *TypeWafFirewallVersion {
	if o == nil {
		return nil
	}
	return o.Type
}
