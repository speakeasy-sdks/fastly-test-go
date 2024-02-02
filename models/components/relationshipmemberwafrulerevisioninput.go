// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type RelationshipMemberWafRuleRevisionInput struct {
	// Resource type.
	Type *TypeWafRuleRevision `default:"waf_rule_revision" json:"type"`
}

func (r RelationshipMemberWafRuleRevisionInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationshipMemberWafRuleRevisionInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RelationshipMemberWafRuleRevisionInput) GetType() *TypeWafRuleRevision {
	if o == nil {
		return nil
	}
	return o.Type
}
