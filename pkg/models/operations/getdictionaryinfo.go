// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetDictionaryInfoRequest struct {
	// Alphanumeric string identifying a Dictionary.
	DictionaryID string `pathParam:"style=simple,explode=false,name=dictionary_id"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetDictionaryInfoRequest) GetDictionaryID() string {
	if o == nil {
		return ""
	}
	return o.DictionaryID
}

func (o *GetDictionaryInfoRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetDictionaryInfoRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetDictionaryInfoResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DictionaryInfoResponse *shared.DictionaryInfoResponse
}

func (o *GetDictionaryInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDictionaryInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDictionaryInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDictionaryInfoResponse) GetDictionaryInfoResponse() *shared.DictionaryInfoResponse {
	if o == nil {
		return nil
	}
	return o.DictionaryInfoResponse
}
