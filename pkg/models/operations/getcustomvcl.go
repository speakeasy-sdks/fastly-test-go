// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"Fastly/pkg/utils"
	"net/http"
)

type GetCustomVclRequest struct {
	// Omit VCL content.
	NoContent *string `default:"0" queryParam:"style=form,explode=true,name=no_content"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// The name of this VCL.
	VclName string `pathParam:"style=simple,explode=false,name=vcl_name"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (g GetCustomVclRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetCustomVclRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetCustomVclRequest) GetNoContent() *string {
	if o == nil {
		return nil
	}
	return o.NoContent
}

func (o *GetCustomVclRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetCustomVclRequest) GetVclName() string {
	if o == nil {
		return ""
	}
	return o.VclName
}

func (o *GetCustomVclRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type GetCustomVclResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	VclResponse *shared.VclResponse
}

func (o *GetCustomVclResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCustomVclResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCustomVclResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCustomVclResponse) GetVclResponse() *shared.VclResponse {
	if o == nil {
		return nil
	}
	return o.VclResponse
}
