// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"errors"
)

type RelationshipsForInvitationType string

const (
	RelationshipsForInvitationTypeRelationshipCustomer           RelationshipsForInvitationType = "relationship_customer"
	RelationshipsForInvitationTypeRelationshipServiceInvitations RelationshipsForInvitationType = "relationship_service_invitations"
)

type RelationshipsForInvitation struct {
	RelationshipCustomer           *RelationshipCustomer
	RelationshipServiceInvitations *RelationshipServiceInvitations

	Type RelationshipsForInvitationType
}

func CreateRelationshipsForInvitationRelationshipCustomer(relationshipCustomer RelationshipCustomer) RelationshipsForInvitation {
	typ := RelationshipsForInvitationTypeRelationshipCustomer

	return RelationshipsForInvitation{
		RelationshipCustomer: &relationshipCustomer,
		Type:                 typ,
	}
}

func CreateRelationshipsForInvitationRelationshipServiceInvitations(relationshipServiceInvitations RelationshipServiceInvitations) RelationshipsForInvitation {
	typ := RelationshipsForInvitationTypeRelationshipServiceInvitations

	return RelationshipsForInvitation{
		RelationshipServiceInvitations: &relationshipServiceInvitations,
		Type:                           typ,
	}
}

func (u *RelationshipsForInvitation) UnmarshalJSON(data []byte) error {

	relationshipCustomer := new(RelationshipCustomer)
	if err := utils.UnmarshalJSON(data, &relationshipCustomer, "", true, true); err == nil {
		u.RelationshipCustomer = relationshipCustomer
		u.Type = RelationshipsForInvitationTypeRelationshipCustomer
		return nil
	}

	relationshipServiceInvitations := new(RelationshipServiceInvitations)
	if err := utils.UnmarshalJSON(data, &relationshipServiceInvitations, "", true, true); err == nil {
		u.RelationshipServiceInvitations = relationshipServiceInvitations
		u.Type = RelationshipsForInvitationTypeRelationshipServiceInvitations
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RelationshipsForInvitation) MarshalJSON() ([]byte, error) {
	if u.RelationshipCustomer != nil {
		return utils.MarshalJSON(u.RelationshipCustomer, "", true)
	}

	if u.RelationshipServiceInvitations != nil {
		return utils.MarshalJSON(u.RelationshipServiceInvitations, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}