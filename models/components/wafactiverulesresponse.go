// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
)

type WafActiveRulesResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

func (o *WafActiveRulesResponseLinks) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *WafActiveRulesResponseLinks) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *WafActiveRulesResponseLinks) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *WafActiveRulesResponseLinks) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type WafActiveRulesResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `default:"20" json:"per_page"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

func (w WafActiveRulesResponseMeta) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafActiveRulesResponseMeta) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafActiveRulesResponseMeta) GetCurrentPage() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentPage
}

func (o *WafActiveRulesResponseMeta) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *WafActiveRulesResponseMeta) GetRecordCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordCount
}

func (o *WafActiveRulesResponseMeta) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

type WafActiveRulesResponse struct {
	Data     []WafActiveRuleResponseData     `json:"data,omitempty"`
	Included []IncludedWithWafActiveRuleItem `json:"included,omitempty"`
	Links    *WafActiveRulesResponseLinks    `json:"links,omitempty"`
	Meta     *WafActiveRulesResponseMeta     `json:"meta,omitempty"`
}

func (o *WafActiveRulesResponse) GetData() []WafActiveRuleResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *WafActiveRulesResponse) GetIncluded() []IncludedWithWafActiveRuleItem {
	if o == nil {
		return nil
	}
	return o.Included
}

func (o *WafActiveRulesResponse) GetLinks() *WafActiveRulesResponseLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *WafActiveRulesResponse) GetMeta() *WafActiveRulesResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}
