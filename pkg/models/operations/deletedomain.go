// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteDomainRequest struct {
	// The name of the domain or domains associated with this service.
	DomainName string `pathParam:"style=simple,explode=false,name=domain_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *DeleteDomainRequest) GetDomainName() string {
	if o == nil {
		return ""
	}
	return o.DomainName
}

func (o *DeleteDomainRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteDomainRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

// DeleteDomain200ApplicationJSON - OK
type DeleteDomain200ApplicationJSON struct {
	// ok
	Status *string `json:"status,omitempty"`
}

func (o *DeleteDomain200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteDomain200ApplicationJSONObject *DeleteDomain200ApplicationJSON
}

func (o *DeleteDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteDomainResponse) GetDeleteDomain200ApplicationJSONObject() *DeleteDomain200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteDomain200ApplicationJSONObject
}
