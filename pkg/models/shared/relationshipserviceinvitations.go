// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RelationshipServiceInvitationsServiceInvitations struct {
	Data []RelationshipMemberServiceInvitation `json:"data,omitempty"`
}

func (o *RelationshipServiceInvitationsServiceInvitations) GetData() []RelationshipMemberServiceInvitation {
	if o == nil {
		return nil
	}
	return o.Data
}

type RelationshipServiceInvitations struct {
	ServiceInvitations *RelationshipServiceInvitationsServiceInvitations `json:"service_invitations,omitempty"`
}

func (o *RelationshipServiceInvitations) GetServiceInvitations() *RelationshipServiceInvitationsServiceInvitations {
	if o == nil {
		return nil
	}
	return o.ServiceInvitations
}
