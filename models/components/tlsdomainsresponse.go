// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
)

type TLSDomainsResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

func (o *TLSDomainsResponseLinks) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *TLSDomainsResponseLinks) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *TLSDomainsResponseLinks) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *TLSDomainsResponseLinks) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type TLSDomainsResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `default:"20" json:"per_page"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

func (t TLSDomainsResponseMeta) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSDomainsResponseMeta) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSDomainsResponseMeta) GetCurrentPage() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentPage
}

func (o *TLSDomainsResponseMeta) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *TLSDomainsResponseMeta) GetRecordCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordCount
}

func (o *TLSDomainsResponseMeta) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

type TLSDomainsResponse struct {
	Data  []TLSDomainData          `json:"data,omitempty"`
	Links *TLSDomainsResponseLinks `json:"links,omitempty"`
	Meta  *TLSDomainsResponseMeta  `json:"meta,omitempty"`
}

func (o *TLSDomainsResponse) GetData() []TLSDomainData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *TLSDomainsResponse) GetLinks() *TLSDomainsResponseLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *TLSDomainsResponse) GetMeta() *TLSDomainsResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}
