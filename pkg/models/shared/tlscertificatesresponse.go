// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TLSCertificatesResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

func (o *TLSCertificatesResponseLinks) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *TLSCertificatesResponseLinks) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *TLSCertificatesResponseLinks) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *TLSCertificatesResponseLinks) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type TLSCertificatesResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `json:"per_page,omitempty"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

func (o *TLSCertificatesResponseMeta) GetCurrentPage() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentPage
}

func (o *TLSCertificatesResponseMeta) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *TLSCertificatesResponseMeta) GetRecordCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordCount
}

func (o *TLSCertificatesResponseMeta) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

type TLSCertificatesResponse struct {
	Data  []TLSCertificateResponseData  `json:"data,omitempty"`
	Links *TLSCertificatesResponseLinks `json:"links,omitempty"`
	Meta  *TLSCertificatesResponseMeta  `json:"meta,omitempty"`
}

func (o *TLSCertificatesResponse) GetData() []TLSCertificateResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *TLSCertificatesResponse) GetLinks() *TLSCertificatesResponseLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *TLSCertificatesResponse) GetMeta() *TLSCertificatesResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}
