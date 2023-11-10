// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/models/components"
	"net/http"
)

type CreateDirectorBackendRequest struct {
	// The name of the backend.
	BackendName string `pathParam:"style=simple,explode=false,name=backend_name"`
	// Name for the Director.
	DirectorName string `pathParam:"style=simple,explode=false,name=director_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *CreateDirectorBackendRequest) GetBackendName() string {
	if o == nil {
		return ""
	}
	return o.BackendName
}

func (o *CreateDirectorBackendRequest) GetDirectorName() string {
	if o == nil {
		return ""
	}
	return o.DirectorName
}

func (o *CreateDirectorBackendRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateDirectorBackendRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type CreateDirectorBackendResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DirectorBackend *components.DirectorBackend
}

func (o *CreateDirectorBackendResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDirectorBackendResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDirectorBackendResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateDirectorBackendResponse) GetDirectorBackend() *components.DirectorBackend {
	if o == nil {
		return nil
	}
	return o.DirectorBackend
}
