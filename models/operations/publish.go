// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type PublishRequest struct {
	PublishRequest *components.PublishRequest `request:"mediaType=application/json"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *PublishRequest) GetPublishRequest() *components.PublishRequest {
	if o == nil {
		return nil
	}
	return o.PublishRequest
}

func (o *PublishRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type PublishResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Res *string
}

func (o *PublishResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PublishResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PublishResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PublishResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
