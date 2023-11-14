// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
)

type TLSSubscriptionsResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

func (o *TLSSubscriptionsResponseLinks) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *TLSSubscriptionsResponseLinks) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *TLSSubscriptionsResponseLinks) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *TLSSubscriptionsResponseLinks) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type TLSSubscriptionsResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `default:"20" json:"per_page"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

func (t TLSSubscriptionsResponseMeta) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSSubscriptionsResponseMeta) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSSubscriptionsResponseMeta) GetCurrentPage() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentPage
}

func (o *TLSSubscriptionsResponseMeta) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *TLSSubscriptionsResponseMeta) GetRecordCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordCount
}

func (o *TLSSubscriptionsResponseMeta) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

type TLSSubscriptionsResponse struct {
	Data  []TLSSubscriptionResponse      `json:"data,omitempty"`
	Links *TLSSubscriptionsResponseLinks `json:"links,omitempty"`
	Meta  *TLSSubscriptionsResponseMeta  `json:"meta,omitempty"`
}

func (o *TLSSubscriptionsResponse) GetData() []TLSSubscriptionResponse {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *TLSSubscriptionsResponse) GetLinks() *TLSSubscriptionsResponseLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *TLSSubscriptionsResponse) GetMeta() *TLSSubscriptionsResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}