// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type WafTagsResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

func (o *WafTagsResponseLinks) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *WafTagsResponseLinks) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *WafTagsResponseLinks) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *WafTagsResponseLinks) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type WafTagsResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `default:"20" json:"per_page"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

func (w WafTagsResponseMeta) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafTagsResponseMeta) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafTagsResponseMeta) GetCurrentPage() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentPage
}

func (o *WafTagsResponseMeta) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *WafTagsResponseMeta) GetRecordCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordCount
}

func (o *WafTagsResponseMeta) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

type WafTagsResponse struct {
	Data     []WafTagsResponseDataItem `json:"data,omitempty"`
	Included []WafRule                 `json:"included,omitempty"`
	Links    *WafTagsResponseLinks     `json:"links,omitempty"`
	Meta     *WafTagsResponseMeta      `json:"meta,omitempty"`
}

func (o *WafTagsResponse) GetData() []WafTagsResponseDataItem {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *WafTagsResponse) GetIncluded() []WafRule {
	if o == nil {
		return nil
	}
	return o.Included
}

func (o *WafTagsResponse) GetLinks() *WafTagsResponseLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *WafTagsResponse) GetMeta() *WafTagsResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}
