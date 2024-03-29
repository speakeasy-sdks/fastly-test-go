// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type ListWafFirewallsRequest struct {
	// Limit the results returned to a specific service.
	FilterServiceID *string `queryParam:"style=form,explode=true,name=filter[service_id]"`
	// Limit the results returned to a specific service version.
	FilterServiceVersionNumber *string `queryParam:"style=form,explode=true,name=filter[service_version_number]"`
	// Include related objects. Optional.
	Include *components.FirewallInclude `default:"waf_firewall_versions" queryParam:"style=form,explode=true,name=include"`
	// Current page.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Number of records per page.
	PageSize *int64 `default:"20" queryParam:"style=form,explode=true,name=page[size]"`
}

func (l ListWafFirewallsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListWafFirewallsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListWafFirewallsRequest) GetFilterServiceID() *string {
	if o == nil {
		return nil
	}
	return o.FilterServiceID
}

func (o *ListWafFirewallsRequest) GetFilterServiceVersionNumber() *string {
	if o == nil {
		return nil
	}
	return o.FilterServiceVersionNumber
}

func (o *ListWafFirewallsRequest) GetInclude() *components.FirewallInclude {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *ListWafFirewallsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListWafFirewallsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type ListWafFirewallsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WafFirewallsResponse *components.WafFirewallsResponse
}

func (o *ListWafFirewallsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListWafFirewallsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListWafFirewallsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListWafFirewallsResponse) GetWafFirewallsResponse() *components.WafFirewallsResponse {
	if o == nil {
		return nil
	}
	return o.WafFirewallsResponse
}
