// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type RelationshipWafRuleWafRule struct {
	Data []RelationshipMemberWafRule `json:"data,omitempty"`
}

func (o *RelationshipWafRuleWafRule) GetData() []RelationshipMemberWafRule {
	if o == nil {
		return nil
	}
	return o.Data
}

type RelationshipWafRule struct {
	WafRule *RelationshipWafRuleWafRule `json:"waf_rule,omitempty"`
}

func (o *RelationshipWafRule) GetWafRule() *RelationshipWafRuleWafRule {
	if o == nil {
		return nil
	}
	return o.WafRule
}
