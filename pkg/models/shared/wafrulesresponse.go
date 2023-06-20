// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WafRulesResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

type WafRulesResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `json:"per_page,omitempty"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

// WafRulesResponse - OK
type WafRulesResponse struct {
	Data     []WafRuleResponseData  `json:"data,omitempty"`
	Included []interface{}          `json:"included,omitempty"`
	Links    *WafRulesResponseLinks `json:"links,omitempty"`
	Meta     *WafRulesResponseMeta  `json:"meta,omitempty"`
}