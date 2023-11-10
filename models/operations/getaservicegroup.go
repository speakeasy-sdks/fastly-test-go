// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAServiceGroupRequest struct {
	// Alphanumeric string identifying the service group.
	ServiceGroupID string `pathParam:"style=simple,explode=false,name=service_group_id"`
}

func (o *GetAServiceGroupRequest) GetServiceGroupID() string {
	if o == nil {
		return ""
	}
	return o.ServiceGroupID
}

// GetAServiceGroupResponseBody - OK
type GetAServiceGroupResponseBody struct {
}

type GetAServiceGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *GetAServiceGroupResponseBody
}

func (o *GetAServiceGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAServiceGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAServiceGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAServiceGroupResponse) GetObject() *GetAServiceGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
