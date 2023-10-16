// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
)

type RelationshipMemberServiceInvitation struct {
	// Alphanumeric string identifying a service invitation.
	ID *string `json:"id,omitempty"`
	// Resource type
	Type *TypeServiceInvitation `default:"service_invitation" json:"type"`
}

func (r RelationshipMemberServiceInvitation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationshipMemberServiceInvitation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RelationshipMemberServiceInvitation) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RelationshipMemberServiceInvitation) GetType() *TypeServiceInvitation {
	if o == nil {
		return nil
	}
	return o.Type
}
