// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type ListHeaderObjectsSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *ListHeaderObjectsSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type ListHeaderObjectsRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *ListHeaderObjectsRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *ListHeaderObjectsRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type ListHeaderObjectsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	HeaderResponses []shared.HeaderResponse
}

func (o *ListHeaderObjectsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListHeaderObjectsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListHeaderObjectsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListHeaderObjectsResponse) GetHeaderResponses() []shared.HeaderResponse {
	if o == nil {
		return nil
	}
	return o.HeaderResponses
}
