// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"errors"
)

type RelationshipsForWafExclusionType string

const (
	RelationshipsForWafExclusionTypeRelationshipWafRulesInput         RelationshipsForWafExclusionType = "relationship_waf_rules_input"
	RelationshipsForWafExclusionTypeRelationshipWafRuleRevisionsInput RelationshipsForWafExclusionType = "relationship_waf_rule_revisions_input"
)

type RelationshipsForWafExclusion struct {
	RelationshipWafRulesInput         *RelationshipWafRulesInput
	RelationshipWafRuleRevisionsInput *RelationshipWafRuleRevisionsInput

	Type RelationshipsForWafExclusionType
}

func CreateRelationshipsForWafExclusionRelationshipWafRulesInput(relationshipWafRulesInput RelationshipWafRulesInput) RelationshipsForWafExclusion {
	typ := RelationshipsForWafExclusionTypeRelationshipWafRulesInput

	return RelationshipsForWafExclusion{
		RelationshipWafRulesInput: &relationshipWafRulesInput,
		Type:                      typ,
	}
}

func CreateRelationshipsForWafExclusionRelationshipWafRuleRevisionsInput(relationshipWafRuleRevisionsInput RelationshipWafRuleRevisionsInput) RelationshipsForWafExclusion {
	typ := RelationshipsForWafExclusionTypeRelationshipWafRuleRevisionsInput

	return RelationshipsForWafExclusion{
		RelationshipWafRuleRevisionsInput: &relationshipWafRuleRevisionsInput,
		Type:                              typ,
	}
}

func (u *RelationshipsForWafExclusion) UnmarshalJSON(data []byte) error {

	relationshipWafRulesInput := RelationshipWafRulesInput{}
	if err := utils.UnmarshalJSON(data, &relationshipWafRulesInput, "", true, true); err == nil {
		u.RelationshipWafRulesInput = &relationshipWafRulesInput
		u.Type = RelationshipsForWafExclusionTypeRelationshipWafRulesInput
		return nil
	}

	relationshipWafRuleRevisionsInput := RelationshipWafRuleRevisionsInput{}
	if err := utils.UnmarshalJSON(data, &relationshipWafRuleRevisionsInput, "", true, true); err == nil {
		u.RelationshipWafRuleRevisionsInput = &relationshipWafRuleRevisionsInput
		u.Type = RelationshipsForWafExclusionTypeRelationshipWafRuleRevisionsInput
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RelationshipsForWafExclusion) MarshalJSON() ([]byte, error) {
	if u.RelationshipWafRulesInput != nil {
		return utils.MarshalJSON(u.RelationshipWafRulesInput, "", true)
	}

	if u.RelationshipWafRuleRevisionsInput != nil {
		return utils.MarshalJSON(u.RelationshipWafRuleRevisionsInput, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
