// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type CreateWafFirewallVersionRequest struct {
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID         string                         `pathParam:"style=simple,explode=false,name=firewall_id"`
	WafFirewallVersion *components.WafFirewallVersion `request:"mediaType=application/vnd.api+json"`
}

func (o *CreateWafFirewallVersionRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *CreateWafFirewallVersionRequest) GetWafFirewallVersion() *components.WafFirewallVersion {
	if o == nil {
		return nil
	}
	return o.WafFirewallVersion
}

type CreateWafFirewallVersionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	WafFirewallVersionResponse *components.WafFirewallVersionResponse
}

func (o *CreateWafFirewallVersionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateWafFirewallVersionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateWafFirewallVersionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateWafFirewallVersionResponse) GetWafFirewallVersionResponse() *components.WafFirewallVersionResponse {
	if o == nil {
		return nil
	}
	return o.WafFirewallVersionResponse
}
