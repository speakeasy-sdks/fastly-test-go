// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateServiceStarResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	StarResponse *shared.StarResponse
}

func (o *CreateServiceStarResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateServiceStarResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateServiceStarResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateServiceStarResponse) GetStarResponse() *shared.StarResponse {
	if o == nil {
		return nil
	}
	return o.StarResponse
}
