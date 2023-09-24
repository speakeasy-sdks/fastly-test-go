// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetLogKinesisRequest struct {
	// The name for the real-time logging configuration.
	LoggingKinesisName string `pathParam:"style=simple,explode=false,name=logging_kinesis_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetLogKinesisRequest) GetLoggingKinesisName() string {
	if o == nil {
		return ""
	}
	return o.LoggingKinesisName
}

func (o *GetLogKinesisRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetLogKinesisRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetLogKinesisResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	LoggingKinesisResponse *shared.LoggingKinesisResponse
}

func (o *GetLogKinesisResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLogKinesisResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLogKinesisResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLogKinesisResponse) GetLoggingKinesisResponse() *shared.LoggingKinesisResponse {
	if o == nil {
		return nil
	}
	return o.LoggingKinesisResponse
}
