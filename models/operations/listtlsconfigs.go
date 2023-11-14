// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/internal/utils"
	"Fastly/models/components"
	"net/http"
)

type ListTLSConfigsRequest struct {
	// Optionally filters by the bulk attribute.
	FilterBulk *string `queryParam:"style=form,explode=true,name=filter[bulk]"`
	// Include related objects. Optional, comma-separated values. Permitted values: `dns_records`.
	//
	Include *string `queryParam:"style=form,explode=true,name=include"`
	// Current page.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Number of records per page.
	PageSize *int64 `default:"20" queryParam:"style=form,explode=true,name=page[size]"`
}

func (l ListTLSConfigsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTLSConfigsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListTLSConfigsRequest) GetFilterBulk() *string {
	if o == nil {
		return nil
	}
	return o.FilterBulk
}

func (o *ListTLSConfigsRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *ListTLSConfigsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListTLSConfigsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type ListTLSConfigsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TLSConfigurationsResponse *components.TLSConfigurationsResponse
}

func (o *ListTLSConfigsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTLSConfigsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTLSConfigsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTLSConfigsResponse) GetTLSConfigurationsResponse() *components.TLSConfigurationsResponse {
	if o == nil {
		return nil
	}
	return o.TLSConfigurationsResponse
}
