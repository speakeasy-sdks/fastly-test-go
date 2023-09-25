// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type PatchTLSSubRequest struct {
	// A flag that allows you to edit and delete a subscription with active domains. Valid to use on PATCH and DELETE actions. As a warning, removing an active domain from a subscription or forcing the deletion of a subscription may result in breaking TLS termination to that domain.
	//
	Force           *bool                   `queryParam:"style=form,explode=false,name=force"`
	TLSSubscription *shared.TLSSubscription `request:"mediaType=application/vnd.api+json"`
	// Alphanumeric string identifying a TLS subscription.
	TLSSubscriptionID string `pathParam:"style=simple,explode=false,name=tls_subscription_id"`
}

func (o *PatchTLSSubRequest) GetForce() *bool {
	if o == nil {
		return nil
	}
	return o.Force
}

func (o *PatchTLSSubRequest) GetTLSSubscription() *shared.TLSSubscription {
	if o == nil {
		return nil
	}
	return o.TLSSubscription
}

func (o *PatchTLSSubRequest) GetTLSSubscriptionID() string {
	if o == nil {
		return ""
	}
	return o.TLSSubscriptionID
}

type PatchTLSSubResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	TLSSubscriptionResponse *shared.TLSSubscriptionResponse
}

func (o *PatchTLSSubResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchTLSSubResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchTLSSubResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchTLSSubResponse) GetTLSSubscriptionResponse() *shared.TLSSubscriptionResponse {
	if o == nil {
		return nil
	}
	return o.TLSSubscriptionResponse
}
