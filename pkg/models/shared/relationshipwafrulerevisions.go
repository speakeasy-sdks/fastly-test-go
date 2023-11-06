// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RelationshipWafRuleRevisionsWafRuleRevisions struct {
	Data []RelationshipMemberWafRuleRevision `json:"data,omitempty"`
}

func (o *RelationshipWafRuleRevisionsWafRuleRevisions) GetData() []RelationshipMemberWafRuleRevision {
	if o == nil {
		return nil
	}
	return o.Data
}

type RelationshipWafRuleRevisions struct {
	WafRuleRevisions *RelationshipWafRuleRevisionsWafRuleRevisions `json:"waf_rule_revisions,omitempty"`
}

func (o *RelationshipWafRuleRevisions) GetWafRuleRevisions() *RelationshipWafRuleRevisionsWafRuleRevisions {
	if o == nil {
		return nil
	}
	return o.WafRuleRevisions
}

type RelationshipWafRuleRevisionsWafRuleRevisionsInput struct {
	Data []RelationshipMemberWafRuleRevisionInput `json:"data,omitempty"`
}

func (o *RelationshipWafRuleRevisionsWafRuleRevisionsInput) GetData() []RelationshipMemberWafRuleRevisionInput {
	if o == nil {
		return nil
	}
	return o.Data
}

type RelationshipWafRuleRevisionsInput struct {
	WafRuleRevisions *RelationshipWafRuleRevisionsWafRuleRevisionsInput `json:"waf_rule_revisions,omitempty"`
}

func (o *RelationshipWafRuleRevisionsInput) GetWafRuleRevisions() *RelationshipWafRuleRevisionsWafRuleRevisionsInput {
	if o == nil {
		return nil
	}
	return o.WafRuleRevisions
}
