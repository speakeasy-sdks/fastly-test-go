// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type ListTLSSubsRequest struct {
	// Limit the returned subscriptions to those that have currently active orders. Permitted values: `true`.
	//
	FilterHasActiveOrder *bool `queryParam:"style=form,explode=true,name=filter[has_active_order]"`
	// Limit the returned subscriptions by state. Valid values are `pending`, `processing`, `issued`, `renewing`, and `failed`. Accepts parameters: `not` (e.g., `filter[state][not]=renewing`).
	//
	FilterState *string `queryParam:"style=form,explode=true,name=filter[state]"`
	// Limit the returned subscriptions to those that include the specific domain.
	FilterTLSDomainsID *string `queryParam:"style=form,explode=true,name=filter[tls_domains.id]"`
	// Include related objects. Optional, comma-separated values. Permitted values: `tls_authorizations` and `tls_authorizations.globalsign_email_challenge`.
	//
	Include *string `queryParam:"style=form,explode=true,name=include"`
	// Current page.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Number of records per page.
	PageSize *int64 `default:"20" queryParam:"style=form,explode=true,name=page[size]"`
	// The order in which to list the results by creation date.
	Sort *components.Sort `default:"created_at" queryParam:"style=form,explode=true,name=sort"`
}

func (l ListTLSSubsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTLSSubsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListTLSSubsRequest) GetFilterHasActiveOrder() *bool {
	if o == nil {
		return nil
	}
	return o.FilterHasActiveOrder
}

func (o *ListTLSSubsRequest) GetFilterState() *string {
	if o == nil {
		return nil
	}
	return o.FilterState
}

func (o *ListTLSSubsRequest) GetFilterTLSDomainsID() *string {
	if o == nil {
		return nil
	}
	return o.FilterTLSDomainsID
}

func (o *ListTLSSubsRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *ListTLSSubsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListTLSSubsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListTLSSubsRequest) GetSort() *components.Sort {
	if o == nil {
		return nil
	}
	return o.Sort
}

type ListTLSSubsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TLSSubscriptionsResponse *components.TLSSubscriptionsResponse
}

func (o *ListTLSSubsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTLSSubsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTLSSubsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTLSSubsResponse) GetTLSSubscriptionsResponse() *components.TLSSubscriptionsResponse {
	if o == nil {
		return nil
	}
	return o.TLSSubscriptionsResponse
}
