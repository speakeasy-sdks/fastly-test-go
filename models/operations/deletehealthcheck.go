// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteHealthcheckRequest struct {
	// The name of the health check.
	HealthcheckName string `pathParam:"style=simple,explode=false,name=healthcheck_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteHealthcheckRequest) GetHealthcheckName() string {
	if o == nil {
		return ""
	}
	return o.HealthcheckName
}

func (o *DeleteHealthcheckRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteHealthcheckRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteHealthcheckResponseBody - OK
type DeleteHealthcheckResponseBody struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteHealthcheckResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteHealthcheckResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *DeleteHealthcheckResponseBody
}

func (o *DeleteHealthcheckResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteHealthcheckResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteHealthcheckResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteHealthcheckResponse) GetObject() *DeleteHealthcheckResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
