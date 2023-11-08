// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
)

type TLSBulkCertificatesResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

func (o *TLSBulkCertificatesResponseLinks) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *TLSBulkCertificatesResponseLinks) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *TLSBulkCertificatesResponseLinks) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *TLSBulkCertificatesResponseLinks) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type TLSBulkCertificatesResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `default:"20" json:"per_page"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

func (t TLSBulkCertificatesResponseMeta) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSBulkCertificatesResponseMeta) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSBulkCertificatesResponseMeta) GetCurrentPage() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentPage
}

func (o *TLSBulkCertificatesResponseMeta) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *TLSBulkCertificatesResponseMeta) GetRecordCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordCount
}

func (o *TLSBulkCertificatesResponseMeta) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

type TLSBulkCertificatesResponse struct {
	Data  []TLSBulkCertificateResponseData  `json:"data,omitempty"`
	Links *TLSBulkCertificatesResponseLinks `json:"links,omitempty"`
	Meta  *TLSBulkCertificatesResponseMeta  `json:"meta,omitempty"`
}

func (o *TLSBulkCertificatesResponse) GetData() []TLSBulkCertificateResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *TLSBulkCertificatesResponse) GetLinks() *TLSBulkCertificatesResponseLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *TLSBulkCertificatesResponse) GetMeta() *TLSBulkCertificatesResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}
