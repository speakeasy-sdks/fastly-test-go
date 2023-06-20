// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateSnippetDynamicSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type UpdateSnippetDynamicRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string           `pathParam:"style=simple,explode=false,name=service_id"`
	Snippet2  *shared.Snippet2 `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying a VCL Snippet.
	SnippetID string `pathParam:"style=simple,explode=false,name=snippet_id"`
}

type UpdateSnippetDynamicResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	SnippetResponse *shared.SnippetResponse
}