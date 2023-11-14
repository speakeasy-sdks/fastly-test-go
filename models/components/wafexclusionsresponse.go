// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
)

type WafExclusionsResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

func (o *WafExclusionsResponseLinks) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *WafExclusionsResponseLinks) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *WafExclusionsResponseLinks) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *WafExclusionsResponseLinks) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type WafExclusionsResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `default:"20" json:"per_page"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

func (w WafExclusionsResponseMeta) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafExclusionsResponseMeta) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafExclusionsResponseMeta) GetCurrentPage() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentPage
}

func (o *WafExclusionsResponseMeta) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *WafExclusionsResponseMeta) GetRecordCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordCount
}

func (o *WafExclusionsResponseMeta) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

type WafExclusionsResponse struct {
	Data     []WafExclusionResponseData     `json:"data,omitempty"`
	Included []IncludedWithWafExclusionItem `json:"included,omitempty"`
	Links    *WafExclusionsResponseLinks    `json:"links,omitempty"`
	Meta     *WafExclusionsResponseMeta     `json:"meta,omitempty"`
}

func (o *WafExclusionsResponse) GetData() []WafExclusionResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *WafExclusionsResponse) GetIncluded() []IncludedWithWafExclusionItem {
	if o == nil {
		return nil
	}
	return o.Included
}

func (o *WafExclusionsResponse) GetLinks() *WafExclusionsResponseLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *WafExclusionsResponse) GetMeta() *WafExclusionsResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}
