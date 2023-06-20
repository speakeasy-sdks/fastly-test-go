// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ListServiceGroupServicesSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type ListServiceGroupServicesRequest struct {
	// Current page.
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page.
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
	// Alphanumeric string identifying the service group.
	ServiceGroupID string `pathParam:"style=simple,explode=false,name=service_group_id"`
}

// ListServiceGroupServices200ApplicationJSON - OK
type ListServiceGroupServices200ApplicationJSON struct {
}

type ListServiceGroupServicesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ListServiceGroupServices200ApplicationJSONObject *ListServiceGroupServices200ApplicationJSON
}
