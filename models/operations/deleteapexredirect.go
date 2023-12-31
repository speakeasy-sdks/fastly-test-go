// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteApexRedirectRequest struct {
	ApexRedirectID string `pathParam:"style=simple,explode=false,name=apex_redirect_id"`
}

func (o *DeleteApexRedirectRequest) GetApexRedirectID() string {
	if o == nil {
		return ""
	}
	return o.ApexRedirectID
}

// DeleteApexRedirectResponseBody - OK
type DeleteApexRedirectResponseBody struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteApexRedirectResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteApexRedirectResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *DeleteApexRedirectResponseBody
}

func (o *DeleteApexRedirectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteApexRedirectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteApexRedirectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteApexRedirectResponse) GetObject() *DeleteApexRedirectResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
