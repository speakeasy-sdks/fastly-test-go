// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type ListWafFirewallVersionsSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type ListWafFirewallVersionsRequest struct {
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// Include relationships. Optional.
	Include *string `queryParam:"style=form,explode=true,name=include"`
	// Current page.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Number of records per page.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
}

type ListWafFirewallVersionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	WafFirewallVersionsResponse *shared.WafFirewallVersionsResponse
}