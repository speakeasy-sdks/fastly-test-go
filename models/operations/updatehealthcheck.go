// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateHealthcheckRequest struct {
	Healthcheck *components.Healthcheck `request:"mediaType=application/x-www-form-urlencoded"`
	// The name of the health check.
	HealthcheckName string `pathParam:"style=simple,explode=false,name=healthcheck_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateHealthcheckRequest) GetHealthcheck() *components.Healthcheck {
	if o == nil {
		return nil
	}
	return o.Healthcheck
}

func (o *UpdateHealthcheckRequest) GetHealthcheckName() string {
	if o == nil {
		return ""
	}
	return o.HealthcheckName
}

func (o *UpdateHealthcheckRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateHealthcheckRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateHealthcheckResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	HealthcheckResponse *components.HealthcheckResponse
}

func (o *UpdateHealthcheckResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateHealthcheckResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateHealthcheckResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateHealthcheckResponse) GetHealthcheckResponse() *components.HealthcheckResponse {
	if o == nil {
		return nil
	}
	return o.HealthcheckResponse
}
