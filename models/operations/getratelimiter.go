// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type GetRateLimiterRequest struct {
	// Alphanumeric string identifying the rate limiter.
	RateLimiterID string `pathParam:"style=simple,explode=false,name=rate_limiter_id"`
}

func (o *GetRateLimiterRequest) GetRateLimiterID() string {
	if o == nil {
		return ""
	}
	return o.RateLimiterID
}

type GetRateLimiterResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	RateLimiterResponse *components.RateLimiterResponse
}

func (o *GetRateLimiterResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRateLimiterResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRateLimiterResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetRateLimiterResponse) GetRateLimiterResponse() *components.RateLimiterResponse {
	if o == nil {
		return nil
	}
	return o.RateLimiterResponse
}
