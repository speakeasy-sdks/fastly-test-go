// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type InvitationsResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

func (o *InvitationsResponseLinks) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *InvitationsResponseLinks) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *InvitationsResponseLinks) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *InvitationsResponseLinks) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type InvitationsResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `json:"per_page,omitempty"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

func (o *InvitationsResponseMeta) GetCurrentPage() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentPage
}

func (o *InvitationsResponseMeta) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *InvitationsResponseMeta) GetRecordCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordCount
}

func (o *InvitationsResponseMeta) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

type InvitationsResponse struct {
	Data  []InvitationResponseData  `json:"data,omitempty"`
	Links *InvitationsResponseLinks `json:"links,omitempty"`
	Meta  *InvitationsResponseMeta  `json:"meta,omitempty"`
}

func (o *InvitationsResponse) GetData() []InvitationResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *InvitationsResponse) GetLinks() *InvitationsResponseLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *InvitationsResponse) GetMeta() *InvitationsResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}
