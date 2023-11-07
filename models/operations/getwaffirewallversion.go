// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type GetWafFirewallVersionRequest struct {
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// Integer identifying a WAF firewall version.
	FirewallVersionNumber int64 `pathParam:"style=simple,explode=false,name=firewall_version_number"`
	// Include relationships. Optional, comma-separated values. Permitted values: `waf_firewall` and `waf_active_rules`.
	//
	Include *string `queryParam:"style=form,explode=true,name=include"`
}

func (o *GetWafFirewallVersionRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *GetWafFirewallVersionRequest) GetFirewallVersionNumber() int64 {
	if o == nil {
		return 0
	}
	return o.FirewallVersionNumber
}

func (o *GetWafFirewallVersionRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

type GetWafFirewallVersionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WafFirewallVersionResponse *components.WafFirewallVersionResponse
}

func (o *GetWafFirewallVersionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWafFirewallVersionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWafFirewallVersionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetWafFirewallVersionResponse) GetWafFirewallVersionResponse() *components.WafFirewallVersionResponse {
	if o == nil {
		return nil
	}
	return o.WafFirewallVersionResponse
}
