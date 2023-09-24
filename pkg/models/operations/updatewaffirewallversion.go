// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateWafFirewallVersionRequest struct {
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// Integer identifying a WAF firewall version.
	FirewallVersionNumber   int64                           `pathParam:"style=simple,explode=false,name=firewall_version_number"`
	WafFirewallVersionInput *shared.WafFirewallVersionInput `request:"mediaType=application/vnd.api+json"`
}

func (o *UpdateWafFirewallVersionRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *UpdateWafFirewallVersionRequest) GetFirewallVersionNumber() int64 {
	if o == nil {
		return 0
	}
	return o.FirewallVersionNumber
}

func (o *UpdateWafFirewallVersionRequest) GetWafFirewallVersionInput() *shared.WafFirewallVersionInput {
	if o == nil {
		return nil
	}
	return o.WafFirewallVersionInput
}

type UpdateWafFirewallVersionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	WafFirewallVersionResponse *shared.WafFirewallVersionResponse
}

func (o *UpdateWafFirewallVersionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateWafFirewallVersionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateWafFirewallVersionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateWafFirewallVersionResponse) GetWafFirewallVersionResponse() *shared.WafFirewallVersionResponse {
	if o == nil {
		return nil
	}
	return o.WafFirewallVersionResponse
}
