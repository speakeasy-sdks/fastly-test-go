// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type TLSPrivateKeysResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

func (o *TLSPrivateKeysResponseLinks) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *TLSPrivateKeysResponseLinks) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *TLSPrivateKeysResponseLinks) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *TLSPrivateKeysResponseLinks) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type TLSPrivateKeysResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `default:"20" json:"per_page"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

func (t TLSPrivateKeysResponseMeta) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSPrivateKeysResponseMeta) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSPrivateKeysResponseMeta) GetCurrentPage() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentPage
}

func (o *TLSPrivateKeysResponseMeta) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *TLSPrivateKeysResponseMeta) GetRecordCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordCount
}

func (o *TLSPrivateKeysResponseMeta) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

type TLSPrivateKeysResponse struct {
	Data  []TLSPrivateKeyResponseData  `json:"data,omitempty"`
	Links *TLSPrivateKeysResponseLinks `json:"links,omitempty"`
	Meta  *TLSPrivateKeysResponseMeta  `json:"meta,omitempty"`
}

func (o *TLSPrivateKeysResponse) GetData() []TLSPrivateKeyResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *TLSPrivateKeysResponse) GetLinks() *TLSPrivateKeysResponseLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *TLSPrivateKeysResponse) GetMeta() *TLSPrivateKeysResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}
