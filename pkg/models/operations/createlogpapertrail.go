// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateLogPapertrailRequest struct {
	LoggingPapertrail2 *shared.LoggingPapertrail2 `request:"mediaType=application/x-www-form-urlencoded"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateLogPapertrailRequest) GetLoggingPapertrail2() *shared.LoggingPapertrail2 {
	if o == nil {
		return nil
	}
	return o.LoggingPapertrail2
}

func (o *CreateLogPapertrailRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateLogPapertrailRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type CreateLogPapertrailResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LoggingPapertrailResponse *shared.LoggingPapertrailResponse
}

func (o *CreateLogPapertrailResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLogPapertrailResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLogPapertrailResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLogPapertrailResponse) GetLoggingPapertrailResponse() *shared.LoggingPapertrailResponse {
	if o == nil {
		return nil
	}
	return o.LoggingPapertrailResponse
}
