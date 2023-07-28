// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateUserSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *CreateUserSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type CreateUserResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	UserResponse *shared.UserResponse
}

func (o *CreateUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateUserResponse) GetUserResponse() *shared.UserResponse {
	if o == nil {
		return nil
	}
	return o.UserResponse
}
