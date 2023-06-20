// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type PurgeTagSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type PurgeTagRequest struct {
	// If present, this header triggers the purge to be 'soft', which marks the affected object as stale rather than making it inaccessible.  Typically set to "1" when used, but the value is not important.
	FastlySoftPurge *int64 `header:"style=simple,explode=false,name=fastly-soft-purge"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Surrogate keys are used to efficiently purge content from cache. Instead of purging your entire site or individual URLs, you can tag related assets (like all images and descriptions associated with a single product) with surrogate keys, and these grouped URLs can be purged in a single request.
	SurrogateKey string `pathParam:"style=simple,explode=false,name=surrogate_key"`
}

type PurgeTagResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PurgeResponse *shared.PurgeResponse
}