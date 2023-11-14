// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
	"net/http"
)

type ListUserGroupRolesRequest struct {
	// Current page.
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page.
	PerPage *int64 `default:"20" queryParam:"style=form,explode=true,name=per_page"`
	// Alphanumeric string identifying the user group.
	UserGroupID string `pathParam:"style=simple,explode=false,name=user_group_id"`
}

func (l ListUserGroupRolesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUserGroupRolesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUserGroupRolesRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListUserGroupRolesRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListUserGroupRolesRequest) GetUserGroupID() string {
	if o == nil {
		return ""
	}
	return o.UserGroupID
}

// ListUserGroupRolesResponseBody - OK
type ListUserGroupRolesResponseBody struct {
}

type ListUserGroupRolesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *ListUserGroupRolesResponseBody
}

func (o *ListUserGroupRolesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUserGroupRolesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUserGroupRolesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListUserGroupRolesResponse) GetObject() *ListUserGroupRolesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
