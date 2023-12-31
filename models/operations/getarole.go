// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetARoleRequest struct {
	// Alphanumeric string identifying the role.
	RoleID string `pathParam:"style=simple,explode=false,name=role_id"`
}

func (o *GetARoleRequest) GetRoleID() string {
	if o == nil {
		return ""
	}
	return o.RoleID
}

// GetARoleResponseBody - OK
type GetARoleResponseBody struct {
}

type GetARoleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *GetARoleResponseBody
}

func (o *GetARoleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetARoleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetARoleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetARoleResponse) GetObject() *GetARoleResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
