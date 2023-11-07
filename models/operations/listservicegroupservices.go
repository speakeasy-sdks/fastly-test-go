// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/internal/utils"
	"net/http"
)

type ListServiceGroupServicesRequest struct {
	// Current page.
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page.
	PerPage *int64 `default:"20" queryParam:"style=form,explode=true,name=per_page"`
	// Alphanumeric string identifying the service group.
	ServiceGroupID string `pathParam:"style=simple,explode=false,name=service_group_id"`
}

func (l ListServiceGroupServicesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListServiceGroupServicesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListServiceGroupServicesRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListServiceGroupServicesRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListServiceGroupServicesRequest) GetServiceGroupID() string {
	if o == nil {
		return ""
	}
	return o.ServiceGroupID
}

// ListServiceGroupServicesResponseBody - OK
type ListServiceGroupServicesResponseBody struct {
}

type ListServiceGroupServicesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *ListServiceGroupServicesResponseBody
}

func (o *ListServiceGroupServicesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListServiceGroupServicesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListServiceGroupServicesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListServiceGroupServicesResponse) GetObject() *ListServiceGroupServicesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
