// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type BulkUpdateWafActiveRulesRequest struct {
	BulkWafActiveRule *components.BulkWafActiveRule `request:"mediaType=application/vnd.api+json"`
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *BulkUpdateWafActiveRulesRequest) GetBulkWafActiveRule() *components.BulkWafActiveRule {
	if o == nil {
		return nil
	}
	return o.BulkWafActiveRule
}

func (o *BulkUpdateWafActiveRulesRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *BulkUpdateWafActiveRulesRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type BulkUpdateWafActiveRulesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *BulkUpdateWafActiveRulesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *BulkUpdateWafActiveRulesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *BulkUpdateWafActiveRulesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
