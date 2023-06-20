// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type ListServicesSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type ListServicesRequest struct {
	// Direction in which to sort results.
	Direction *shared.Direction `queryParam:"style=form,explode=true,name=direction"`
	// Current page.
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page.
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
	// Field on which to sort.
	Sort *string `queryParam:"style=form,explode=true,name=sort"`
}

type ListServicesResponse struct {
	ContentType string
	Headers     map[string][]string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ServiceListResponses []shared.ServiceListResponse
}