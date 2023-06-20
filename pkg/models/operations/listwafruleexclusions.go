// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

type ListWafRuleExclusionsSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

// ListWafRuleExclusionsFilterExclusionType - Filters the results based on this exclusion type.
type ListWafRuleExclusionsFilterExclusionType string

const (
	ListWafRuleExclusionsFilterExclusionTypeRule     ListWafRuleExclusionsFilterExclusionType = "rule"
	ListWafRuleExclusionsFilterExclusionTypeVariable ListWafRuleExclusionsFilterExclusionType = "variable"
	ListWafRuleExclusionsFilterExclusionTypeWaf      ListWafRuleExclusionsFilterExclusionType = "waf"
)

func (e ListWafRuleExclusionsFilterExclusionType) ToPointer() *ListWafRuleExclusionsFilterExclusionType {
	return &e
}

func (e *ListWafRuleExclusionsFilterExclusionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "rule":
		fallthrough
	case "variable":
		fallthrough
	case "waf":
		*e = ListWafRuleExclusionsFilterExclusionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListWafRuleExclusionsFilterExclusionType: %v", v)
	}
}

type ListWafRuleExclusionsRequest struct {
	// Filters the results based on this exclusion type.
	FilterExclusionType *ListWafRuleExclusionsFilterExclusionType `queryParam:"style=form,explode=true,name=filter[exclusion_type]"`
	// Filters the results based on name.
	FilterName *string `queryParam:"style=form,explode=true,name=filter[name]"`
	// Filters the results based on this ModSecurity rule ID.
	FilterWafRulesModsecRuleID *int64 `queryParam:"style=form,explode=true,name=filter[waf_rules.modsec_rule_id]"`
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// Integer identifying a WAF firewall version.
	FirewallVersionNumber int64 `pathParam:"style=simple,explode=false,name=firewall_version_number"`
	// Include relationships. Optional, comma-separated values. Permitted values: `waf_rules` and `waf_rule_revisions`.
	//
	Include *string `queryParam:"style=form,explode=true,name=include"`
	// Current page.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Number of records per page.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
}

type ListWafRuleExclusionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	WafExclusionsResponse *shared.WafExclusionsResponse
}