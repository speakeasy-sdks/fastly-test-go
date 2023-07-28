// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetSnippetDynamicSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *GetSnippetDynamicSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type GetSnippetDynamicRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Alphanumeric string identifying a VCL Snippet.
	SnippetID string `pathParam:"style=simple,explode=false,name=snippet_id"`
}

func (o *GetSnippetDynamicRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetSnippetDynamicRequest) GetSnippetID() string {
	if o == nil {
		return ""
	}
	return o.SnippetID
}

type GetSnippetDynamicResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	SnippetResponse *shared.SnippetResponse
}

func (o *GetSnippetDynamicResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSnippetDynamicResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSnippetDynamicResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSnippetDynamicResponse) GetSnippetResponse() *shared.SnippetResponse {
	if o == nil {
		return nil
	}
	return o.SnippetResponse
}
