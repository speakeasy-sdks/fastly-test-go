// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateLogPapertrailSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type UpdateLogPapertrailRequest struct {
	LoggingPapertrail2 *shared.LoggingPapertrail2 `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingPapertrailName string `pathParam:"style=simple,explode=false,name=logging_papertrail_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

type UpdateLogPapertrailResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingPapertrailResponse *shared.LoggingPapertrailResponse
}