// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"net/http"
)

// FilterExclusionType - Filters the results based on this exclusion type.
type FilterExclusionType string

const (
	FilterExclusionTypeRule     FilterExclusionType = "rule"
	FilterExclusionTypeVariable FilterExclusionType = "variable"
	FilterExclusionTypeWaf      FilterExclusionType = "waf"
)

func (e FilterExclusionType) ToPointer() *FilterExclusionType {
	return &e
}

func (e *FilterExclusionType) UnmarshalJSON(data []byte) error {
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
		*e = FilterExclusionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FilterExclusionType: %v", v)
	}
}

type ListWafRuleExclusionsRequest struct {
	// Filters the results based on this exclusion type.
	FilterExclusionType *FilterExclusionType `queryParam:"style=form,explode=true,name=filter[exclusion_type]"`
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
	PageSize *int64 `default:"20" queryParam:"style=form,explode=true,name=page[size]"`
}

func (l ListWafRuleExclusionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListWafRuleExclusionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListWafRuleExclusionsRequest) GetFilterExclusionType() *FilterExclusionType {
	if o == nil {
		return nil
	}
	return o.FilterExclusionType
}

func (o *ListWafRuleExclusionsRequest) GetFilterName() *string {
	if o == nil {
		return nil
	}
	return o.FilterName
}

func (o *ListWafRuleExclusionsRequest) GetFilterWafRulesModsecRuleID() *int64 {
	if o == nil {
		return nil
	}
	return o.FilterWafRulesModsecRuleID
}

func (o *ListWafRuleExclusionsRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *ListWafRuleExclusionsRequest) GetFirewallVersionNumber() int64 {
	if o == nil {
		return 0
	}
	return o.FirewallVersionNumber
}

func (o *ListWafRuleExclusionsRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *ListWafRuleExclusionsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListWafRuleExclusionsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type ListWafRuleExclusionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WafExclusionsResponse *components.WafExclusionsResponse
}

func (o *ListWafRuleExclusionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListWafRuleExclusionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListWafRuleExclusionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListWafRuleExclusionsResponse) GetWafExclusionsResponse() *components.WafExclusionsResponse {
	if o == nil {
		return nil
	}
	return o.WafExclusionsResponse
}
