// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateLogPapertrailRequest struct {
	LoggingPapertrail2 *shared.LoggingPapertrail2 `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingPapertrailName string `pathParam:"style=simple,explode=false,name=logging_papertrail_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogPapertrailRequest) GetLoggingPapertrail2() *shared.LoggingPapertrail2 {
	if o == nil {
		return nil
	}
	return o.LoggingPapertrail2
}

func (o *UpdateLogPapertrailRequest) GetLoggingPapertrailName() string {
	if o == nil {
		return ""
	}
	return o.LoggingPapertrailName
}

func (o *UpdateLogPapertrailRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogPapertrailRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogPapertrailResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingPapertrailResponse *shared.LoggingPapertrailResponse
}

func (o *UpdateLogPapertrailResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogPapertrailResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogPapertrailResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogPapertrailResponse) GetLoggingPapertrailResponse() *shared.LoggingPapertrailResponse {
	if o == nil {
		return nil
	}
	return o.LoggingPapertrailResponse
}
