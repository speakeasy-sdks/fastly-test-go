// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetHttp3Security struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *GetHttp3Security) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type GetHttp3Request struct {
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *GetHttp3Request) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetHttp3Request) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetHttp3Response struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Http3 *shared.Http3
}

func (o *GetHttp3Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHttp3Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHttp3Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetHttp3Response) GetHttp3() *shared.Http3 {
	if o == nil {
		return nil
	}
	return o.Http3
}
