// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateUserPasswordSecurity struct {
	Password string `security:"scheme,type=http,subtype=basic,name=password"`
	Username string `security:"scheme,type=http,subtype=basic,name=username"`
}

func (o *UpdateUserPasswordSecurity) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *UpdateUserPasswordSecurity) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type UpdateUserPasswordResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UserResponse *shared.UserResponse
}

func (o *UpdateUserPasswordResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateUserPasswordResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateUserPasswordResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateUserPasswordResponse) GetUserResponse() *shared.UserResponse {
	if o == nil {
		return nil
	}
	return o.UserResponse
}
