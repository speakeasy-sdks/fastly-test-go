// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"errors"
)

type RelationshipsForWafExclusionInputType string

const (
	RelationshipsForWafExclusionInputTypeRelationshipWafRulesInput         RelationshipsForWafExclusionInputType = "relationship_waf_rulesInput"
	RelationshipsForWafExclusionInputTypeRelationshipWafRuleRevisionsInput RelationshipsForWafExclusionInputType = "relationship_waf_rule_revisionsInput"
)

type RelationshipsForWafExclusionInput struct {
	RelationshipWafRulesInput         *RelationshipWafRulesInput
	RelationshipWafRuleRevisionsInput *RelationshipWafRuleRevisionsInput

	Type RelationshipsForWafExclusionInputType
}

func CreateRelationshipsForWafExclusionInputRelationshipWafRulesInput(relationshipWafRulesInput RelationshipWafRulesInput) RelationshipsForWafExclusionInput {
	typ := RelationshipsForWafExclusionInputTypeRelationshipWafRulesInput

	return RelationshipsForWafExclusionInput{
		RelationshipWafRulesInput: &relationshipWafRulesInput,
		Type:                      typ,
	}
}

func CreateRelationshipsForWafExclusionInputRelationshipWafRuleRevisionsInput(relationshipWafRuleRevisionsInput RelationshipWafRuleRevisionsInput) RelationshipsForWafExclusionInput {
	typ := RelationshipsForWafExclusionInputTypeRelationshipWafRuleRevisionsInput

	return RelationshipsForWafExclusionInput{
		RelationshipWafRuleRevisionsInput: &relationshipWafRuleRevisionsInput,
		Type:                              typ,
	}
}

func (u *RelationshipsForWafExclusionInput) UnmarshalJSON(data []byte) error {

	relationshipWafRulesInput := new(RelationshipWafRulesInput)
	if err := utils.UnmarshalJSON(data, &relationshipWafRulesInput, "", true, true); err == nil {
		u.RelationshipWafRulesInput = relationshipWafRulesInput
		u.Type = RelationshipsForWafExclusionInputTypeRelationshipWafRulesInput
		return nil
	}

	relationshipWafRuleRevisionsInput := new(RelationshipWafRuleRevisionsInput)
	if err := utils.UnmarshalJSON(data, &relationshipWafRuleRevisionsInput, "", true, true); err == nil {
		u.RelationshipWafRuleRevisionsInput = relationshipWafRuleRevisionsInput
		u.Type = RelationshipsForWafExclusionInputTypeRelationshipWafRuleRevisionsInput
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RelationshipsForWafExclusionInput) MarshalJSON() ([]byte, error) {
	if u.RelationshipWafRulesInput != nil {
		return utils.MarshalJSON(u.RelationshipWafRulesInput, "", true)
	}

	if u.RelationshipWafRuleRevisionsInput != nil {
		return utils.MarshalJSON(u.RelationshipWafRuleRevisionsInput, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}