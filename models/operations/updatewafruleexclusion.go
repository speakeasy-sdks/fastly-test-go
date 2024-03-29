// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateWafRuleExclusionRequest struct {
	// A numeric ID identifying a WAF exclusion.
	ExclusionNumber int64 `pathParam:"style=simple,explode=false,name=exclusion_number"`
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// Integer identifying a WAF firewall version.
	FirewallVersionNumber int64                    `pathParam:"style=simple,explode=false,name=firewall_version_number"`
	WafExclusion          *components.WafExclusion `request:"mediaType=application/vnd.api+json"`
}

func (o *UpdateWafRuleExclusionRequest) GetExclusionNumber() int64 {
	if o == nil {
		return 0
	}
	return o.ExclusionNumber
}

func (o *UpdateWafRuleExclusionRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *UpdateWafRuleExclusionRequest) GetFirewallVersionNumber() int64 {
	if o == nil {
		return 0
	}
	return o.FirewallVersionNumber
}

func (o *UpdateWafRuleExclusionRequest) GetWafExclusion() *components.WafExclusion {
	if o == nil {
		return nil
	}
	return o.WafExclusion
}

type UpdateWafRuleExclusionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	WafExclusionResponse *components.WafExclusionResponse
}

func (o *UpdateWafRuleExclusionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateWafRuleExclusionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateWafRuleExclusionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateWafRuleExclusionResponse) GetWafExclusionResponse() *components.WafExclusionResponse {
	if o == nil {
		return nil
	}
	return o.WafExclusionResponse
}
