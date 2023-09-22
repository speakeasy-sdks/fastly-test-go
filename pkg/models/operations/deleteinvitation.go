// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteInvitationRequest struct {
	// Alphanumeric string identifying an invitation.
	InvitationID string `pathParam:"style=simple,explode=false,name=invitation_id"`
}

func (o *DeleteInvitationRequest) GetInvitationID() string {
	if o == nil {
		return ""
	}
	return o.InvitationID
}

type DeleteInvitationResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteInvitationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteInvitationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteInvitationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
