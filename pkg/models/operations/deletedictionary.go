// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteDictionarySecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteDictionaryRequest struct {
	// Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace).
	DictionaryName string `pathParam:"style=simple,explode=false,name=dictionary_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

// DeleteDictionary200ApplicationJSON - OK
type DeleteDictionary200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

type DeleteDictionaryResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteDictionary200ApplicationJSONObject *DeleteDictionary200ApplicationJSON
}