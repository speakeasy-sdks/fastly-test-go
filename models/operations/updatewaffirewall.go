// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateWafFirewallRequest struct {
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID  string                  `pathParam:"style=simple,explode=false,name=firewall_id"`
	WafFirewall *components.WafFirewall `request:"mediaType=application/vnd.api+json"`
}

func (o *UpdateWafFirewallRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *UpdateWafFirewallRequest) GetWafFirewall() *components.WafFirewall {
	if o == nil {
		return nil
	}
	return o.WafFirewall
}

type UpdateWafFirewallResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WafFirewallResponse *components.WafFirewallResponse
}

func (o *UpdateWafFirewallResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateWafFirewallResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateWafFirewallResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateWafFirewallResponse) GetWafFirewallResponse() *components.WafFirewallResponse {
	if o == nil {
		return nil
	}
	return o.WafFirewallResponse
}
