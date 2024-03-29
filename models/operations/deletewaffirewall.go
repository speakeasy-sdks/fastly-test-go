// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type DeleteWafFirewallRequest struct {
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID  string                  `pathParam:"style=simple,explode=false,name=firewall_id"`
	WafFirewall *components.WafFirewall `request:"mediaType=application/json"`
}

func (o *DeleteWafFirewallRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *DeleteWafFirewallRequest) GetWafFirewall() *components.WafFirewall {
	if o == nil {
		return nil
	}
	return o.WafFirewall
}

type DeleteWafFirewallResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteWafFirewallResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteWafFirewallResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteWafFirewallResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
