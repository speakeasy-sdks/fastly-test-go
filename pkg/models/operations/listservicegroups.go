// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ListServiceGroupsSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *ListServiceGroupsSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type ListServiceGroupsRequest struct {
	// Current page.
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page.
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
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
	ContentType string
	StatusCode  int
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
