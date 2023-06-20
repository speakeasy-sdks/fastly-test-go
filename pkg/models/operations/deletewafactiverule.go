// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteWafActiveRuleSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteWafActiveRuleRequest struct {
	// Alphanumeric string identifying a WAF Firewall.
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
	// Alphanumeric string identifying a WAF rule.
	WafRuleID string `pathParam:"style=simple,explode=false,name=waf_rule_id"`
}

type DeleteWafActiveRuleResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}