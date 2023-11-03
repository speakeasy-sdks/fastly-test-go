// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateApexRedirectRequest struct {
	ApexRedirectInput *shared.ApexRedirectInput `request:"mediaType=application/x-www-form-urlencoded"`
	ApexRedirectID    string                    `pathParam:"style=simple,explode=false,name=apex_redirect_id"`
}

func (o *UpdateApexRedirectRequest) GetApexRedirectInput() *shared.ApexRedirectInput {
	if o == nil {
		return nil
	}
	return o.ApexRedirectInput
}

func (o *UpdateApexRedirectRequest) GetApexRedirectID() string {
	if o == nil {
		return ""
	}
	return o.ApexRedirectID
}

type UpdateApexRedirectResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ApexRedirect *shared.ApexRedirect
}

func (o *UpdateApexRedirectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateApexRedirectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateApexRedirectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateApexRedirectResponse) GetApexRedirect() *shared.ApexRedirect {
	if o == nil {
		return nil
	}
	return o.ApexRedirect
}
