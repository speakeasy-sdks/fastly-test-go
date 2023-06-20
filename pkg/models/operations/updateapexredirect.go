// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateApexRedirectSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type UpdateApexRedirectRequest struct {
	ApexRedirectInput *shared.ApexRedirectInput `request:"mediaType=application/x-www-form-urlencoded"`
	ApexRedirectID    string                    `pathParam:"style=simple,explode=false,name=apex_redirect_id"`
}

type UpdateApexRedirectResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ApexRedirect *shared.ApexRedirect
}