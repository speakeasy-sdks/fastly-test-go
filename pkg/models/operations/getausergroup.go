// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAUserGroupSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type GetAUserGroupRequest struct {
	// Alphanumeric string identifying the user group.
	UserGroupID string `pathParam:"style=simple,explode=false,name=user_group_id"`
}

// GetAUserGroup200ApplicationJSON - OK
type GetAUserGroup200ApplicationJSON struct {
}

type GetAUserGroupResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetAUserGroup200ApplicationJSONObject *GetAUserGroup200ApplicationJSON
}