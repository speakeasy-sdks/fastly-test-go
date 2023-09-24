// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/utils"
	"net/http"
)

type ListUserGroupServiceGroupsRequest struct {
	// Current page.
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page.
	PerPage *int64 `default:"20" queryParam:"style=form,explode=true,name=per_page"`
	// Alphanumeric string identifying the user group.
	UserGroupID string `pathParam:"style=simple,explode=false,name=user_group_id"`
}

func (l ListUserGroupServiceGroupsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUserGroupServiceGroupsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUserGroupServiceGroupsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListUserGroupServiceGroupsRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListUserGroupServiceGroupsRequest) GetUserGroupID() string {
	if o == nil {
		return ""
	}
	return o.UserGroupID
}

// ListUserGroupServiceGroups200ApplicationJSON - OK
type ListUserGroupServiceGroups200ApplicationJSON struct {
}

type ListUserGroupServiceGroupsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ListUserGroupServiceGroups200ApplicationJSONObject *ListUserGroupServiceGroups200ApplicationJSON
}

func (o *ListUserGroupServiceGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUserGroupServiceGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUserGroupServiceGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListUserGroupServiceGroupsResponse) GetListUserGroupServiceGroups200ApplicationJSONObject() *ListUserGroupServiceGroups200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListUserGroupServiceGroups200ApplicationJSONObject
}
