// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeployActivateWafFirewallVersionRequest struct {
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// Integer identifying a WAF firewall version.
	FirewallVersionNumber int64 `pathParam:"style=simple,explode=false,name=firewall_version_number"`
}

func (o *DeployActivateWafFirewallVersionRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *DeployActivateWafFirewallVersionRequest) GetFirewallVersionNumber() int64 {
	if o == nil {
		return 0
	}
	return o.FirewallVersionNumber
}

// DeployActivateWafFirewallVersion202ApplicationVndAPIPlusJSON - Accepted
type DeployActivateWafFirewallVersion202ApplicationVndAPIPlusJSON struct {
}

type DeployActivateWafFirewallVersionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Accepted
	DeployActivateWafFirewallVersion202ApplicationVndAPIPlusJSONObject *DeployActivateWafFirewallVersion202ApplicationVndAPIPlusJSON
}

func (o *DeployActivateWafFirewallVersionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeployActivateWafFirewallVersionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeployActivateWafFirewallVersionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeployActivateWafFirewallVersionResponse) GetDeployActivateWafFirewallVersion202ApplicationVndAPIPlusJSONObject() *DeployActivateWafFirewallVersion202ApplicationVndAPIPlusJSON {
	if o == nil {
		return nil
	}
	return o.DeployActivateWafFirewallVersion202ApplicationVndAPIPlusJSONObject
}
