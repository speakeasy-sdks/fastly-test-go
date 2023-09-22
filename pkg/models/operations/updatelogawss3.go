// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateLogAwsS3Request struct {
	LoggingS3Input *shared.LoggingS3Input `request:"mediaType=application/x-www-form-urlencoded"`
	// The name for the real-time logging configuration.
	LoggingS3Name string `pathParam:"style=simple,explode=false,name=logging_s3_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateLogAwsS3Request) GetLoggingS3Input() *shared.LoggingS3Input {
	if o == nil {
		return nil
	}
	return o.LoggingS3Input
}

func (o *UpdateLogAwsS3Request) GetLoggingS3Name() string {
	if o == nil {
		return ""
	}
	return o.LoggingS3Name
}

func (o *UpdateLogAwsS3Request) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateLogAwsS3Request) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateLogAwsS3Response struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingS3Response *shared.LoggingS3Response
}

func (o *UpdateLogAwsS3Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLogAwsS3Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLogAwsS3Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLogAwsS3Response) GetLoggingS3Response() *shared.LoggingS3Response {
	if o == nil {
		return nil
	}
	return o.LoggingS3Response
}
