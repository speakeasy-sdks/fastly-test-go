// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLogSftpSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteLogSftpRequest struct {
	// The name for the real-time logging configuration.
	LoggingSftpName string `pathParam:"style=simple,explode=false,name=logging_sftp_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

// DeleteLogSftp200ApplicationJSON - OK
type DeleteLogSftp200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteLogSftpResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteLogSftp200ApplicationJSONObject *DeleteLogSftp200ApplicationJSON
}
