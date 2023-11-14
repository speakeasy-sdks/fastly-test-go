// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
)

type ServiceAuthorizationsResponseLinks struct {
	// The first page of data.
	First *string `json:"first,omitempty"`
	// The last page of data.
	Last *string `json:"last,omitempty"`
	// The next page of data.
	Next *string `json:"next,omitempty"`
	// The previous page of data.
	Prev *string `json:"prev,omitempty"`
}

func (o *ServiceAuthorizationsResponseLinks) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *ServiceAuthorizationsResponseLinks) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *ServiceAuthorizationsResponseLinks) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ServiceAuthorizationsResponseLinks) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type ServiceAuthorizationsResponseMeta struct {
	// Current page.
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int64 `default:"20" json:"per_page"`
	// Total records in result set.
	RecordCount *int64 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages *int64 `json:"total_pages,omitempty"`
}

func (s ServiceAuthorizationsResponseMeta) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServiceAuthorizationsResponseMeta) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ServiceAuthorizationsResponseMeta) GetCurrentPage() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentPage
}

func (o *ServiceAuthorizationsResponseMeta) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ServiceAuthorizationsResponseMeta) GetRecordCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordCount
}

func (o *ServiceAuthorizationsResponseMeta) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

type ServiceAuthorizationsResponse struct {
	Data  []ServiceAuthorizationResponseData  `json:"data,omitempty"`
	Links *ServiceAuthorizationsResponseLinks `json:"links,omitempty"`
	Meta  *ServiceAuthorizationsResponseMeta  `json:"meta,omitempty"`
}

func (o *ServiceAuthorizationsResponse) GetData() []ServiceAuthorizationResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ServiceAuthorizationsResponse) GetLinks() *ServiceAuthorizationsResponseLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *ServiceAuthorizationsResponse) GetMeta() *ServiceAuthorizationsResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}
