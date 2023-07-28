// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateWafFirewallVersionSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *CreateWafFirewallVersionSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type CreateWafFirewallVersionRequest struct {
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID              string                          `pathParam:"style=simple,explode=false,name=firewall_id"`
	WafFirewallVersionInput *shared.WafFirewallVersionInput `request:"mediaType=application/vnd.api+json"`
}

func (o *CreateWafFirewallVersionRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *CreateWafFirewallVersionRequest) GetWafFirewallVersionInput() *shared.WafFirewallVersionInput {
	if o == nil {
		return nil
	}
	return o.WafFirewallVersionInput
}

type CreateWafFirewallVersionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Created
	WafFirewallVersionResponse *shared.WafFirewallVersionResponse
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

func (o *CreateWafFirewallVersionResponse) GetWafFirewallVersionResponse() *shared.WafFirewallVersionResponse {
	if o == nil {
		return nil
	}
	return o.WafFirewallVersionResponse
}
