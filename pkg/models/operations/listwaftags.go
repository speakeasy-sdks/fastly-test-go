// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"Fastly/pkg/utils"
	"net/http"
)

type ListWafTagsRequest struct {
	// Limit the returned tags to a specific name.
	FilterName *string `queryParam:"style=form,explode=true,name=filter[name]"`
	// Include relationships. Optional.
	Include *shared.WafTagInclude `default:"waf_rules" queryParam:"style=form,explode=true,name=include"`
	// Current page.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Number of records per page.
	PageSize *int64 `default:"20" queryParam:"style=form,explode=true,name=page[size]"`
}

func (l ListWafTagsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListWafTagsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListWafTagsRequest) GetFilterName() *string {
	if o == nil {
		return nil
	}
	return o.FilterName
}

func (o *ListWafTagsRequest) GetInclude() *shared.WafTagInclude {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *ListWafTagsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListWafTagsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type ListWafTagsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WafTagsResponse *shared.WafTagsResponse
}

func (o *ListWafTagsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListWafTagsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListWafTagsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListWafTagsResponse) GetWafTagsResponse() *shared.WafTagsResponse {
	if o == nil {
		return nil
	}
	return o.WafTagsResponse
}
