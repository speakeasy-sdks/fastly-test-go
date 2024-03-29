// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type PatchTLSSubRequest struct {
	// A flag that allows you to edit and delete a subscription with active domains. Valid to use on PATCH and DELETE actions. As a warning, removing an active domain from a subscription or forcing the deletion of a subscription may result in breaking TLS termination to that domain.
	//
	Force           *bool                       `queryParam:"style=form,explode=false,name=force"`
	TLSSubscription *components.TLSSubscription `request:"mediaType=application/vnd.api+json"`
	// Alphanumeric string identifying a TLS subscription.
	TLSSubscriptionID string `pathParam:"style=simple,explode=false,name=tls_subscription_id"`
}

func (o *PatchTLSSubRequest) GetForce() *bool {
	if o == nil {
		return nil
	}
	return o.Force
}

func (o *PatchTLSSubRequest) GetTLSSubscription() *components.TLSSubscription {
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
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TLSSubscriptionResponse *components.TLSSubscriptionResponse
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

func (o *PatchTLSSubResponse) GetTLSSubscriptionResponse() *components.TLSSubscriptionResponse {
	if o == nil {
		return nil
	}
	return o.TLSSubscriptionResponse
}
