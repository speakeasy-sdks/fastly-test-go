// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type GetSnippetRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// The name for the snippet.
	SnippetName string `pathParam:"style=simple,explode=false,name=snippet_name"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetSnippetRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetSnippetRequest) GetSnippetName() string {
	if o == nil {
		return ""
	}
	return o.SnippetName
}

func (o *GetSnippetRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetSnippetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	SnippetResponse *components.SnippetResponse
}

func (o *GetSnippetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSnippetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSnippetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSnippetResponse) GetSnippetResponse() *components.SnippetResponse {
	if o == nil {
		return nil
	}
	return o.SnippetResponse
}
