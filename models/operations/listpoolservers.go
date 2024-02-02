// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type ListPoolServersRequest struct {
	// Alphanumeric string identifying a Pool.
	PoolID string `pathParam:"style=simple,explode=false,name=pool_id"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *ListPoolServersRequest) GetPoolID() string {
	if o == nil {
		return ""
	}
	return o.PoolID
}

func (o *ListPoolServersRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type ListPoolServersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Classes []components.ServerResponse
}

func (o *ListPoolServersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPoolServersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPoolServersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListPoolServersResponse) GetClasses() []components.ServerResponse {
	if o == nil {
		return nil
	}
	return o.Classes
}
