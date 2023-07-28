// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type ListContactsSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *ListContactsSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type ListContactsRequest struct {
	// Alphanumeric string identifying the customer.
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
}

func (o *ListContactsRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

type ListContactsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	SchemasContactResponses []shared.SchemasContactResponse
}

func (o *ListContactsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListContactsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListContactsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListContactsResponse) GetSchemasContactResponses() []shared.SchemasContactResponse {
	if o == nil {
		return nil
	}
	return o.SchemasContactResponses
}
