// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type CreateWafRuleExclusionRequest struct {
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// Integer identifying a WAF firewall version.
	FirewallVersionNumber int64                    `pathParam:"style=simple,explode=false,name=firewall_version_number"`
	WafExclusion          *components.WafExclusion `request:"mediaType=application/vnd.api+json"`
}

func (o *CreateWafRuleExclusionRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *CreateWafRuleExclusionRequest) GetFirewallVersionNumber() int64 {
	if o == nil {
		return 0
	}
	return o.FirewallVersionNumber
}

func (o *CreateWafRuleExclusionRequest) GetWafExclusion() *components.WafExclusion {
	if o == nil {
		return nil
	}
	return o.WafExclusion
}

type CreateWafRuleExclusionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	WafExclusionResponse *components.WafExclusionResponse
}

func (o *CreateWafRuleExclusionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateWafRuleExclusionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateWafRuleExclusionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateWafRuleExclusionResponse) GetWafExclusionResponse() *components.WafExclusionResponse {
	if o == nil {
		return nil
	}
	return o.WafExclusionResponse
}
