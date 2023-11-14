// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteResponseObjectRequest struct {
	// Name for the request settings.
	ResponseObjectName string `pathParam:"style=simple,explode=false,name=response_object_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteResponseObjectRequest) GetResponseObjectName() string {
	if o == nil {
		return ""
	}
	return o.ResponseObjectName
}

func (o *DeleteResponseObjectRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteResponseObjectRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteResponseObjectResponseBody - OK
type DeleteResponseObjectResponseBody struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteResponseObjectResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteResponseObjectResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *DeleteResponseObjectResponseBody
}

func (o *DeleteResponseObjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteResponseObjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteResponseObjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteResponseObjectResponse) GetObject() *DeleteResponseObjectResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}