// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetCustomVclBoilerplateSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *GetCustomVclBoilerplateSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type GetCustomVclBoilerplateRequest struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetCustomVclBoilerplateRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetCustomVclBoilerplateRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetCustomVclBoilerplateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetCustomVclBoilerplate200TextPlainString *string
}

func (o *GetCustomVclBoilerplateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCustomVclBoilerplateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCustomVclBoilerplateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCustomVclBoilerplateResponse) GetGetCustomVclBoilerplate200TextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.GetCustomVclBoilerplate200TextPlainString
}
