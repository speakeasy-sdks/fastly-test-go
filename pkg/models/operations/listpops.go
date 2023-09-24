// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type ListPopsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Pops []shared.Pop
}

func (o *ListPopsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPopsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPopsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListPopsResponse) GetPops() []shared.Pop {
	if o == nil {
		return nil
	}
	return o.Pops
}
