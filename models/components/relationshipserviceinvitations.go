// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type ServiceInvitations struct {
	Data []RelationshipMemberServiceInvitation `json:"data,omitempty"`
}

func (o *ServiceInvitations) GetData() []RelationshipMemberServiceInvitation {
	if o == nil {
		return nil
	}
	return o.Data
}

type RelationshipServiceInvitations struct {
	ServiceInvitations *ServiceInvitations `json:"service_invitations,omitempty"`
}

func (o *RelationshipServiceInvitations) GetServiceInvitations() *ServiceInvitations {
	if o == nil {
		return nil
	}
	return o.ServiceInvitations
}