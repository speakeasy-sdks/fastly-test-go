// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type UpdateConditionRequest struct {
	Condition *components.Condition `request:"mediaType=application/x-www-form-urlencoded"`
	// Name of the condition. Required.
	ConditionName string `pathParam:"style=simple,explode=false,name=condition_name"`
	// Alphanumeric string identifying the service.
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Integer identifying a service version.
	VersionID int64 `pathParam:"style=simple,explode=false,name=version_id"`
}

func (o *UpdateConditionRequest) GetCondition() *components.Condition {
	if o == nil {
		return nil
	}
	return o.Condition
}

func (o *UpdateConditionRequest) GetConditionName() string {
	if o == nil {
		return ""
	}
	return o.ConditionName
}

func (o *UpdateConditionRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateConditionRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

type UpdateConditionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ConditionResponse *components.ConditionResponse
}

func (o *UpdateConditionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConditionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConditionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateConditionResponse) GetConditionResponse() *components.ConditionResponse {
	if o == nil {
		return nil
	}
	return o.ConditionResponse
}
