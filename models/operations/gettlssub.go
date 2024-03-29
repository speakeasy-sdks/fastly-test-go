// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type GetTLSSubRequest struct {
	// Include related objects. Optional, comma-separated values. Permitted values: `tls_authorizations` and `tls_authorizations.globalsign_email_challenge`.
	//
	Include *string `queryParam:"style=form,explode=true,name=include"`
	// Alphanumeric string identifying a TLS subscription.
	TLSSubscriptionID string `pathParam:"style=simple,explode=false,name=tls_subscription_id"`
}

func (o *GetTLSSubRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *GetTLSSubRequest) GetTLSSubscriptionID() string {
	if o == nil {
		return ""
	}
	return o.TLSSubscriptionID
}

type GetTLSSubResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TLSSubscriptionResponse *components.TLSSubscriptionResponse
}

func (o *GetTLSSubResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTLSSubResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTLSSubResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTLSSubResponse) GetTLSSubscriptionResponse() *components.TLSSubscriptionResponse {
	if o == nil {
		return nil
	}
	return o.TLSSubscriptionResponse
}
