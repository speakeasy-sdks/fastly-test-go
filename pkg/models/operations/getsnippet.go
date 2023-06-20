// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetSnippetSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type GetSnippetRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// The name for the snippet.
	SnippetName string `pathParam:"style=simple,explode=false,name=snippet_name"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

type GetSnippetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	SnippetResponse *shared.SnippetResponse
}