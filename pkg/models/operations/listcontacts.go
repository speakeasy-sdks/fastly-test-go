// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type ListContactsSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type ListContactsRequest struct {
	// Alphanumeric string identifying the customer.
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
}

type ListContactsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	SchemasContactResponses []shared.SchemasContactResponse
}