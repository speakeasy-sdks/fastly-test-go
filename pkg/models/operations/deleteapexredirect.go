// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteApexRedirectSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteApexRedirectRequest struct {
	ApexRedirectID string `pathParam:"style=simple,explode=false,name=apex_redirect_id"`
}

// DeleteApexRedirect200ApplicationJSON - OK
type DeleteApexRedirect200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteApexRedirectResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteApexRedirect200ApplicationJSONObject *DeleteApexRedirect200ApplicationJSON
}