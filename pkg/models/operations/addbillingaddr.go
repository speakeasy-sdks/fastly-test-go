// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type AddBillingAddrSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type AddBillingAddrRequest struct {
	// Billing address
	BillingAddressRequestInput *shared.BillingAddressRequestInput `request:"mediaType=application/vnd.api+json"`
	// Alphanumeric string identifying the customer.
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
}

type AddBillingAddrResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Created
	BillingAddressResponse *shared.BillingAddressResponse
	// Could not validate address
	BillingAddressVerificationErrorResponse *shared.BillingAddressVerificationErrorResponse
}