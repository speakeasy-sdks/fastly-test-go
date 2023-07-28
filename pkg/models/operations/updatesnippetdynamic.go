// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateSnippetDynamicSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *UpdateSnippetDynamicSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type UpdateSnippetDynamicRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string           `pathParam:"style=simple,explode=false,name=service_id"`
	Snippet2  *shared.Snippet2 `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying a VCL Snippet.
	SnippetID string `pathParam:"style=simple,explode=false,name=snippet_id"`
}

func (o *UpdateSnippetDynamicRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateSnippetDynamicRequest) GetSnippet2() *shared.Snippet2 {
	if o == nil {
		return nil
	}
	return o.Snippet2
}

func (o *UpdateSnippetDynamicRequest) GetSnippetID() string {
	if o == nil {
		return ""
	}
	return o.SnippetID
}

type UpdateSnippetDynamicResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	SnippetResponse *shared.SnippetResponse
}

func (o *UpdateSnippetDynamicResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSnippetDynamicResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSnippetDynamicResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSnippetDynamicResponse) GetSnippetResponse() *shared.SnippetResponse {
	if o == nil {
		return nil
	}
	return o.SnippetResponse
}
