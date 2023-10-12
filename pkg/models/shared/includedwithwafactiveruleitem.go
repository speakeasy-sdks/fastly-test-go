// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"errors"
)

type IncludedWithWafActiveRuleItemType string

const (
	IncludedWithWafActiveRuleItemTypeSchemasWafFirewallVersion IncludedWithWafActiveRuleItemType = "schemas_waf_firewall_version"
	IncludedWithWafActiveRuleItemTypeWafRuleRevision           IncludedWithWafActiveRuleItemType = "waf_rule_revision"
)

type IncludedWithWafActiveRuleItem struct {
	SchemasWafFirewallVersion *SchemasWafFirewallVersion
	WafRuleRevision           *WafRuleRevision

	Type IncludedWithWafActiveRuleItemType
}

func CreateIncludedWithWafActiveRuleItemSchemasWafFirewallVersion(schemasWafFirewallVersion SchemasWafFirewallVersion) IncludedWithWafActiveRuleItem {
	typ := IncludedWithWafActiveRuleItemTypeSchemasWafFirewallVersion

	return IncludedWithWafActiveRuleItem{
		SchemasWafFirewallVersion: &schemasWafFirewallVersion,
		Type:                      typ,
	}
}

func CreateIncludedWithWafActiveRuleItemWafRuleRevision(wafRuleRevision WafRuleRevision) IncludedWithWafActiveRuleItem {
	typ := IncludedWithWafActiveRuleItemTypeWafRuleRevision

	return IncludedWithWafActiveRuleItem{
		WafRuleRevision: &wafRuleRevision,
		Type:            typ,
	}
}

func (u *IncludedWithWafActiveRuleItem) UnmarshalJSON(data []byte) error {

	schemasWafFirewallVersion := new(SchemasWafFirewallVersion)
	if err := utils.UnmarshalJSON(data, &schemasWafFirewallVersion, "", true, true); err == nil {
		u.SchemasWafFirewallVersion = schemasWafFirewallVersion
		u.Type = IncludedWithWafActiveRuleItemTypeSchemasWafFirewallVersion
		return nil
	}

	wafRuleRevision := new(WafRuleRevision)
	if err := utils.UnmarshalJSON(data, &wafRuleRevision, "", true, true); err == nil {
		u.WafRuleRevision = wafRuleRevision
		u.Type = IncludedWithWafActiveRuleItemTypeWafRuleRevision
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u IncludedWithWafActiveRuleItem) MarshalJSON() ([]byte, error) {
	if u.SchemasWafFirewallVersion != nil {
		return utils.MarshalJSON(u.SchemasWafFirewallVersion, "", true)
	}

	if u.WafRuleRevision != nil {
		return utils.MarshalJSON(u.WafRuleRevision, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
