// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WafFirewallsResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

func (o *WafFirewallsResponseLinks) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *WafFirewallsResponseLinks) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *WafFirewallsResponseLinks) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *WafFirewallsResponseLinks) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type WafFirewallsResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `json:"per_page,omitempty"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

func (o *WafFirewallsResponseMeta) GetCurrentPage() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentPage
}

func (o *WafFirewallsResponseMeta) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *WafFirewallsResponseMeta) GetRecordCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordCount
}

func (o *WafFirewallsResponseMeta) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

type WafFirewallsResponse struct {
	Data     []WafFirewallResponseData   `json:"data,omitempty"`
	Included []SchemasWafFirewallVersion `json:"included,omitempty"`
	Links    *WafFirewallsResponseLinks  `json:"links,omitempty"`
	Meta     *WafFirewallsResponseMeta   `json:"meta,omitempty"`
}

func (o *WafFirewallsResponse) GetData() []WafFirewallResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *WafFirewallsResponse) GetIncluded() []SchemasWafFirewallVersion {
	if o == nil {
		return nil
	}
	return o.Included
}

func (o *WafFirewallsResponse) GetLinks() *WafFirewallsResponseLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *WafFirewallsResponse) GetMeta() *WafFirewallsResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}
