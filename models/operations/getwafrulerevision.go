// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type GetWafRuleRevisionRequest struct {
	// Include relationships. Optional, comma-separated values. Permitted values: `waf_rule`, `vcl`, and `source`. The `vcl` and `source` relationships show the WAF VCL and corresponding ModSecurity source. These fields are blank unless the relationship is included.
	//
	Include *string `queryParam:"style=form,explode=true,name=include"`
	// Alphanumeric string identifying a WAF rule.
	WafRuleID string `pathParam:"style=simple,explode=false,name=waf_rule_id"`
	// Revision number.
	WafRuleRevisionNumber int64 `pathParam:"style=simple,explode=false,name=waf_rule_revision_number"`
}

func (o *GetWafRuleRevisionRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *GetWafRuleRevisionRequest) GetWafRuleID() string {
	if o == nil {
		return ""
	}
	return o.WafRuleID
}

func (o *GetWafRuleRevisionRequest) GetWafRuleRevisionNumber() int64 {
	if o == nil {
		return 0
	}
	return o.WafRuleRevisionNumber
}

type GetWafRuleRevisionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WafRuleRevisionResponse *components.WafRuleRevisionResponse
}

func (o *GetWafRuleRevisionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWafRuleRevisionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWafRuleRevisionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetWafRuleRevisionResponse) GetWafRuleRevisionResponse() *components.WafRuleRevisionResponse {
	if o == nil {
		return nil
	}
	return o.WafRuleRevisionResponse
}
