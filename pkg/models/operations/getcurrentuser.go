// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetCurrentUserSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type GetCurrentUserResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	UserResponse *shared.UserResponse
}
