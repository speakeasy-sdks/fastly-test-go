// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/utils"
	"net/http"
)

type ListServiceGroupsRequest struct {
	// Current page.
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page.
	PerPage *int64 `default:"20" queryParam:"style=form,explode=true,name=per_page"`
}

func (l ListServiceGroupsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListServiceGroupsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListServiceGroupsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListServiceGroupsRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

// ListServiceGroups200ApplicationJSON - OK
type ListServiceGroups200ApplicationJSON struct {
}

type ListServiceGroupsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ListServiceGroups200ApplicationJSONObject *ListServiceGroups200ApplicationJSON
}

func (o *ListServiceGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListServiceGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListServiceGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListServiceGroupsResponse) GetListServiceGroups200ApplicationJSONObject() *ListServiceGroups200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListServiceGroups200ApplicationJSONObject
}
