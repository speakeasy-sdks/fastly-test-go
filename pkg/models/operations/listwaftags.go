// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type ListWafTagsSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type ListWafTagsRequest struct {
	// Limit the returned tags to a specific name.
	FilterName *string `queryParam:"style=form,explode=true,name=filter[name]"`
	// Include relationships. Optional.
	Include *shared.WafTagInclude `queryParam:"style=form,explode=true,name=include"`
	// Current page.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Number of records per page.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
}

type ListWafTagsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	WafTagsResponse *shared.WafTagsResponse
}