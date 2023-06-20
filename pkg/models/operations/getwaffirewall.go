// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetWafFirewallSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type GetWafFirewallRequest struct {
	// Limit the results returned to a specific service version.
	FilterServiceVersionNumber *string `queryParam:"style=form,explode=true,name=filter[service_version_number]"`
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// Include related objects. Optional.
	Include *shared.FirewallInclude `queryParam:"style=form,explode=true,name=include"`
}

type GetWafFirewallResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	WafFirewallResponse *shared.WafFirewallResponse
}
