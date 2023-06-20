// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteARoleSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteARoleRequest struct {
	// Alphanumeric string identifying the role.
	RoleID string `pathParam:"style=simple,explode=false,name=role_id"`
}

type DeleteARoleResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}