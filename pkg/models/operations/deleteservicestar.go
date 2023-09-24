// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteServiceStarRequest struct {
	// Alphanumeric string identifying a star.
	StarID string `pathParam:"style=simple,explode=false,name=star_id"`
}

func (o *DeleteServiceStarRequest) GetStarID() string {
	if o == nil {
		return ""
	}
	return o.StarID
}

type DeleteServiceStarResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteServiceStarResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteServiceStarResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteServiceStarResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
