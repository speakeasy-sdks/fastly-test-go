// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteApexRedirectSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *DeleteApexRedirectSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type DeleteApexRedirectRequest struct {
	ApexRedirectID string `pathParam:"style=simple,explode=false,name=apex_redirect_id"`
}

func (o *DeleteApexRedirectRequest) GetApexRedirectID() string {
	if o == nil {
		return ""
	}
	return o.ApexRedirectID
}

// DeleteApexRedirect200ApplicationJSON - OK
type DeleteApexRedirect200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteApexRedirect200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteApexRedirectResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteApexRedirect200ApplicationJSONObject *DeleteApexRedirect200ApplicationJSON
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

func (o *DeleteApexRedirectResponse) GetDeleteApexRedirect200ApplicationJSONObject() *DeleteApexRedirect200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteApexRedirect200ApplicationJSONObject
}
