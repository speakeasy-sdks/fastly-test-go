// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type UpdateDictionarySecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *UpdateDictionarySecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type UpdateDictionaryRequest struct {
	Dictionary *shared.Dictionary `request:"mediaType=application/x-www-form-urlencoded"`
	// Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace).
	DictionaryName string `pathParam:"style=simple,explode=false,name=dictionary_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateDictionaryRequest) GetDictionary() *shared.Dictionary {
	if o == nil {
		return nil
	}
	return o.Dictionary
}

func (o *UpdateDictionaryRequest) GetDictionaryName() string {
	if o == nil {
		return ""
	}
	return o.DictionaryName
}

func (o *UpdateDictionaryRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateDictionaryRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateDictionaryResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DictionaryResponse *shared.DictionaryResponse
}

func (o *UpdateDictionaryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDictionaryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDictionaryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateDictionaryResponse) GetDictionaryResponse() *shared.DictionaryResponse {
	if o == nil {
		return nil
	}
	return o.DictionaryResponse
}
