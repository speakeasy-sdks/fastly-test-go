// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetUserRequest struct {
	// Alphanumeric string identifying the user.
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

func (o *GetUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type GetUserResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	UserResponse *shared.UserResponse
}

func (o *GetUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUserResponse) GetUserResponse() *shared.UserResponse {
	if o == nil {
		return nil
	}
	return o.UserResponse
}
