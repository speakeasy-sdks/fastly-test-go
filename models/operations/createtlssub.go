// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type CreateTLSSubRequest struct {
	// A flag that allows you to edit and delete a subscription with active domains. Valid to use on PATCH and DELETE actions. As a warning, removing an active domain from a subscription or forcing the deletion of a subscription may result in breaking TLS termination to that domain.
	//
	Force           *bool                       `queryParam:"style=form,explode=false,name=force"`
	TLSSubscription *components.TLSSubscription `request:"mediaType=application/vnd.api+json"`
}

func (o *CreateTLSSubRequest) GetForce() *bool {
	if o == nil {
		return nil
	}
	return o.Force
}

func (o *CreateTLSSubRequest) GetTLSSubscription() *components.TLSSubscription {
	if o == nil {
		return nil
	}
	return o.TLSSubscription
}

type CreateTLSSubResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	TLSSubscriptionResponse *components.TLSSubscriptionResponse
}

func (o *CreateTLSSubResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTLSSubResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTLSSubResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTLSSubResponse) GetTLSSubscriptionResponse() *components.TLSSubscriptionResponse {
	if o == nil {
		return nil
	}
	return o.TLSSubscriptionResponse
}
