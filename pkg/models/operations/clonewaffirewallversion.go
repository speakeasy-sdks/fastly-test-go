// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CloneWafFirewallVersionRequest struct {
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// Integer identifying a WAF firewall version.
	FirewallVersionNumber int64 `pathParam:"style=simple,explode=false,name=firewall_version_number"`
}

func (o *CloneWafFirewallVersionRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *CloneWafFirewallVersionRequest) GetFirewallVersionNumber() int64 {
	if o == nil {
		return 0
	}
	return o.FirewallVersionNumber
}

type CloneWafFirewallVersionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	WafFirewallVersionResponse *shared.WafFirewallVersionResponse
}

func (o *CloneWafFirewallVersionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CloneWafFirewallVersionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CloneWafFirewallVersionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CloneWafFirewallVersionResponse) GetWafFirewallVersionResponse() *shared.WafFirewallVersionResponse {
	if o == nil {
		return nil
	}
	return o.WafFirewallVersionResponse
}
