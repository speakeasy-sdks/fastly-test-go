// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type GetWafFirewallRequest struct {
	// Limit the results returned to a specific service version.
	FilterServiceVersionNumber *string `queryParam:"style=form,explode=true,name=filter[service_version_number]"`
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// Include related objects. Optional.
	Include *components.FirewallInclude `default:"waf_firewall_versions" queryParam:"style=form,explode=true,name=include"`
}

func (g GetWafFirewallRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetWafFirewallRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetWafFirewallRequest) GetFilterServiceVersionNumber() *string {
	if o == nil {
		return nil
	}
	return o.FilterServiceVersionNumber
}

func (o *GetWafFirewallRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *GetWafFirewallRequest) GetInclude() *components.FirewallInclude {
	if o == nil {
		return nil
	}
	return o.Include
}

type GetWafFirewallResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WafFirewallResponse *components.WafFirewallResponse
}

func (o *GetWafFirewallResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWafFirewallResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWafFirewallResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetWafFirewallResponse) GetWafFirewallResponse() *components.WafFirewallResponse {
	if o == nil {
		return nil
	}
	return o.WafFirewallResponse
}
