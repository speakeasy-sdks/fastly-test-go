// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
)

type RelationshipsForWafActiveRuleType string

const (
	RelationshipsForWafActiveRuleTypeRelationshipWafFirewallVersion RelationshipsForWafActiveRuleType = "relationship_waf_firewall_version"
	RelationshipsForWafActiveRuleTypeRelationshipWafRuleRevision    RelationshipsForWafActiveRuleType = "relationship_waf_rule_revision"
)

type RelationshipsForWafActiveRule struct {
	RelationshipWafFirewallVersion *RelationshipWafFirewallVersion
	RelationshipWafRuleRevision    *RelationshipWafRuleRevision

	Type RelationshipsForWafActiveRuleType
}

func CreateRelationshipsForWafActiveRuleRelationshipWafFirewallVersion(relationshipWafFirewallVersion RelationshipWafFirewallVersion) RelationshipsForWafActiveRule {
	typ := RelationshipsForWafActiveRuleTypeRelationshipWafFirewallVersion

	return RelationshipsForWafActiveRule{
		RelationshipWafFirewallVersion: &relationshipWafFirewallVersion,
		Type:                           typ,
	}
}

func CreateRelationshipsForWafActiveRuleRelationshipWafRuleRevision(relationshipWafRuleRevision RelationshipWafRuleRevision) RelationshipsForWafActiveRule {
	typ := RelationshipsForWafActiveRuleTypeRelationshipWafRuleRevision

	return RelationshipsForWafActiveRule{
		RelationshipWafRuleRevision: &relationshipWafRuleRevision,
		Type:                        typ,
	}
}

func (u *RelationshipsForWafActiveRule) UnmarshalJSON(data []byte) error {

	relationshipWafFirewallVersion := RelationshipWafFirewallVersion{}
	if err := utils.UnmarshalJSON(data, &relationshipWafFirewallVersion, "", true, true); err == nil {
		u.RelationshipWafFirewallVersion = &relationshipWafFirewallVersion
		u.Type = RelationshipsForWafActiveRuleTypeRelationshipWafFirewallVersion
		return nil
	}

	relationshipWafRuleRevision := RelationshipWafRuleRevision{}
	if err := utils.UnmarshalJSON(data, &relationshipWafRuleRevision, "", true, true); err == nil {
		u.RelationshipWafRuleRevision = &relationshipWafRuleRevision
		u.Type = RelationshipsForWafActiveRuleTypeRelationshipWafRuleRevision
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RelationshipsForWafActiveRule) MarshalJSON() ([]byte, error) {
	if u.RelationshipWafFirewallVersion != nil {
		return utils.MarshalJSON(u.RelationshipWafFirewallVersion, "", true)
	}

	if u.RelationshipWafRuleRevision != nil {
		return utils.MarshalJSON(u.RelationshipWafRuleRevision, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
