// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type WafFirewallVersionsResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

func (o *WafFirewallVersionsResponseLinks) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *WafFirewallVersionsResponseLinks) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *WafFirewallVersionsResponseLinks) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *WafFirewallVersionsResponseLinks) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type WafFirewallVersionsResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `default:"20" json:"per_page"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

func (w WafFirewallVersionsResponseMeta) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafFirewallVersionsResponseMeta) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafFirewallVersionsResponseMeta) GetCurrentPage() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentPage
}

func (o *WafFirewallVersionsResponseMeta) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *WafFirewallVersionsResponseMeta) GetRecordCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordCount
}

func (o *WafFirewallVersionsResponseMeta) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

type WafFirewallVersionsResponse struct {
	Data     []WafFirewallVersionResponseData     `json:"data,omitempty"`
	Included []IncludedWithWafFirewallVersionItem `json:"included,omitempty"`
	Links    *WafFirewallVersionsResponseLinks    `json:"links,omitempty"`
	Meta     *WafFirewallVersionsResponseMeta     `json:"meta,omitempty"`
}

func (o *WafFirewallVersionsResponse) GetData() []WafFirewallVersionResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *WafFirewallVersionsResponse) GetIncluded() []IncludedWithWafFirewallVersionItem {
	if o == nil {
		return nil
	}
	return o.Included
}

func (o *WafFirewallVersionsResponse) GetLinks() *WafFirewallVersionsResponseLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *WafFirewallVersionsResponse) GetMeta() *WafFirewallVersionsResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}
