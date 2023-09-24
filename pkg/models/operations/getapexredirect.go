// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetApexRedirectRequest struct {
	ApexRedirectID string `pathParam:"style=simple,explode=false,name=apex_redirect_id"`
}

func (o *GetApexRedirectRequest) GetApexRedirectID() string {
	if o == nil {
		return ""
	}
	return o.ApexRedirectID
}

type GetApexRedirectResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ApexRedirect *shared.ApexRedirect
}

func (o *GetApexRedirectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetApexRedirectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetApexRedirectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetApexRedirectResponse) GetApexRedirect() *shared.ApexRedirect {
	if o == nil {
		return nil
	}
	return o.ApexRedirect
}
