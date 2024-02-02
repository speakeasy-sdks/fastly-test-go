// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type RelationshipsForWafFirewallVersionType string

const (
	RelationshipsForWafFirewallVersionTypeRelationshipWafFirewallVersions RelationshipsForWafFirewallVersionType = "relationship_waf_firewall_versions"
	RelationshipsForWafFirewallVersionTypeRelationshipWafActiveRules      RelationshipsForWafFirewallVersionType = "relationship_waf_active_rules"
)

type RelationshipsForWafFirewallVersion struct {
	RelationshipWafFirewallVersions *RelationshipWafFirewallVersions
	RelationshipWafActiveRules      *RelationshipWafActiveRules

	Type RelationshipsForWafFirewallVersionType
}

func CreateRelationshipsForWafFirewallVersionRelationshipWafFirewallVersions(relationshipWafFirewallVersions RelationshipWafFirewallVersions) RelationshipsForWafFirewallVersion {
	typ := RelationshipsForWafFirewallVersionTypeRelationshipWafFirewallVersions

	return RelationshipsForWafFirewallVersion{
		RelationshipWafFirewallVersions: &relationshipWafFirewallVersions,
		Type:                            typ,
	}
}

func CreateRelationshipsForWafFirewallVersionRelationshipWafActiveRules(relationshipWafActiveRules RelationshipWafActiveRules) RelationshipsForWafFirewallVersion {
	typ := RelationshipsForWafFirewallVersionTypeRelationshipWafActiveRules

	return RelationshipsForWafFirewallVersion{
		RelationshipWafActiveRules: &relationshipWafActiveRules,
		Type:                       typ,
	}
}

func (u *RelationshipsForWafFirewallVersion) UnmarshalJSON(data []byte) error {

	relationshipWafFirewallVersions := RelationshipWafFirewallVersions{}
	if err := utils.UnmarshalJSON(data, &relationshipWafFirewallVersions, "", true, true); err == nil {
		u.RelationshipWafFirewallVersions = &relationshipWafFirewallVersions
		u.Type = RelationshipsForWafFirewallVersionTypeRelationshipWafFirewallVersions
		return nil
	}

	relationshipWafActiveRules := RelationshipWafActiveRules{}
	if err := utils.UnmarshalJSON(data, &relationshipWafActiveRules, "", true, true); err == nil {
		u.RelationshipWafActiveRules = &relationshipWafActiveRules
		u.Type = RelationshipsForWafFirewallVersionTypeRelationshipWafActiveRules
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RelationshipsForWafFirewallVersion) MarshalJSON() ([]byte, error) {
	if u.RelationshipWafFirewallVersions != nil {
		return utils.MarshalJSON(u.RelationshipWafFirewallVersions, "", true)
	}

	if u.RelationshipWafActiveRules != nil {
		return utils.MarshalJSON(u.RelationshipWafActiveRules, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
