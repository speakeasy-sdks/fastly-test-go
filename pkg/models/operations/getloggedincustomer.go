// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetLoggedInCustomerSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type GetLoggedInCustomerResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	CustomerResponse *shared.CustomerResponse
}
