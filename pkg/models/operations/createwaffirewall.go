// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type CreateWafFirewallResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Created
	WafFirewallResponse *shared.WafFirewallResponse
}

func (o *CreateWafFirewallResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateWafFirewallResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateWafFirewallResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateWafFirewallResponse) GetWafFirewallResponse() *shared.WafFirewallResponse {
	if o == nil {
		return nil
	}
	return o.WafFirewallResponse
}
