// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type GetWafRuleRequest struct {
	// Include relationships. Optional, comma-separated values. Permitted values: `waf_tags` and `waf_rule_revisions`.
	//
	Include *string `queryParam:"style=form,explode=true,name=include"`
	// Alphanumeric string identifying a WAF rule.
	WafRuleID string `pathParam:"style=simple,explode=false,name=waf_rule_id"`
}

func (o *GetWafRuleRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *GetWafRuleRequest) GetWafRuleID() string {
	if o == nil {
		return ""
	}
	return o.WafRuleID
}

type GetWafRuleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WafRuleResponse *components.WafRuleResponse
}

func (o *GetWafRuleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWafRuleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWafRuleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetWafRuleResponse) GetWafRuleResponse() *components.WafRuleResponse {
	if o == nil {
		return nil
	}
	return o.WafRuleResponse
}
