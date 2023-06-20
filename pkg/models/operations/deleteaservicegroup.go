// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteAServiceGroupSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type DeleteAServiceGroupRequest struct {
	// Alphanumeric string identifying the service group.
	ServiceGroupID string `pathParam:"style=simple,explode=false,name=service_group_id"`
}

type DeleteAServiceGroupResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}