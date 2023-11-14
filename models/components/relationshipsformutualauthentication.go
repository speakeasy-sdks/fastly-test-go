// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"errors"
)

type RelationshipsForMutualAuthenticationType string

const (
	RelationshipsForMutualAuthenticationTypeRelationshipTLSActivations RelationshipsForMutualAuthenticationType = "relationship_tls_activations"
)

type RelationshipsForMutualAuthentication struct {
	RelationshipTLSActivations *RelationshipTLSActivations

	Type RelationshipsForMutualAuthenticationType
}

func CreateRelationshipsForMutualAuthenticationRelationshipTLSActivations(relationshipTLSActivations RelationshipTLSActivations) RelationshipsForMutualAuthentication {
	typ := RelationshipsForMutualAuthenticationTypeRelationshipTLSActivations

	return RelationshipsForMutualAuthentication{
		RelationshipTLSActivations: &relationshipTLSActivations,
		Type:                       typ,
	}
}

func (u *RelationshipsForMutualAuthentication) UnmarshalJSON(data []byte) error {

	relationshipTLSActivations := RelationshipTLSActivations{}
	if err := utils.UnmarshalJSON(data, &relationshipTLSActivations, "", true, true); err == nil {
		u.RelationshipTLSActivations = &relationshipTLSActivations
		u.Type = RelationshipsForMutualAuthenticationTypeRelationshipTLSActivations
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RelationshipsForMutualAuthentication) MarshalJSON() ([]byte, error) {
	if u.RelationshipTLSActivations != nil {
		return utils.MarshalJSON(u.RelationshipTLSActivations, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}