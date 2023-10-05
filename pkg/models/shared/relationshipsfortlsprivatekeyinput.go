// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"errors"
)

type RelationshipsForTLSPrivateKey2TLSDomainsInput struct {
	Data []RelationshipMemberTLSDomainInput `json:"data,omitempty"`
}

func (o *RelationshipsForTLSPrivateKey2TLSDomainsInput) GetData() []RelationshipMemberTLSDomainInput {
	if o == nil {
		return nil
	}
	return o.Data
}

// RelationshipsForTLSPrivateKey2Input - All the domains (including wildcard domains) that are listed in any certificate's Subject Alternative Names (SAN) list.
type RelationshipsForTLSPrivateKey2Input struct {
	TLSDomains *RelationshipsForTLSPrivateKey2TLSDomainsInput `json:"tls_domains,omitempty"`
}

func (o *RelationshipsForTLSPrivateKey2Input) GetTLSDomains() *RelationshipsForTLSPrivateKey2TLSDomainsInput {
	if o == nil {
		return nil
	}
	return o.TLSDomains
}

type RelationshipsForTLSPrivateKeyInputType string

const (
	RelationshipsForTLSPrivateKeyInputTypeRelationshipTLSActivationsInput     RelationshipsForTLSPrivateKeyInputType = "relationship_tls_activationsInput"
	RelationshipsForTLSPrivateKeyInputTypeRelationshipsForTLSPrivateKey2Input RelationshipsForTLSPrivateKeyInputType = "relationships_for_tls_private_key_2Input"
)

type RelationshipsForTLSPrivateKeyInput struct {
	RelationshipTLSActivationsInput     *RelationshipTLSActivationsInput
	RelationshipsForTLSPrivateKey2Input *RelationshipsForTLSPrivateKey2Input

	Type RelationshipsForTLSPrivateKeyInputType
}

func CreateRelationshipsForTLSPrivateKeyInputRelationshipTLSActivationsInput(relationshipTLSActivationsInput RelationshipTLSActivationsInput) RelationshipsForTLSPrivateKeyInput {
	typ := RelationshipsForTLSPrivateKeyInputTypeRelationshipTLSActivationsInput

	return RelationshipsForTLSPrivateKeyInput{
		RelationshipTLSActivationsInput: &relationshipTLSActivationsInput,
		Type:                            typ,
	}
}

func CreateRelationshipsForTLSPrivateKeyInputRelationshipsForTLSPrivateKey2Input(relationshipsForTLSPrivateKey2Input RelationshipsForTLSPrivateKey2Input) RelationshipsForTLSPrivateKeyInput {
	typ := RelationshipsForTLSPrivateKeyInputTypeRelationshipsForTLSPrivateKey2Input

	return RelationshipsForTLSPrivateKeyInput{
		RelationshipsForTLSPrivateKey2Input: &relationshipsForTLSPrivateKey2Input,
		Type:                                typ,
	}
}

func (u *RelationshipsForTLSPrivateKeyInput) UnmarshalJSON(data []byte) error {

	relationshipTLSActivationsInput := new(RelationshipTLSActivationsInput)
	if err := utils.UnmarshalJSON(data, &relationshipTLSActivationsInput, "", true, true); err == nil {
		u.RelationshipTLSActivationsInput = relationshipTLSActivationsInput
		u.Type = RelationshipsForTLSPrivateKeyInputTypeRelationshipTLSActivationsInput
		return nil
	}

	relationshipsForTLSPrivateKey2Input := new(RelationshipsForTLSPrivateKey2Input)
	if err := utils.UnmarshalJSON(data, &relationshipsForTLSPrivateKey2Input, "", true, true); err == nil {
		u.RelationshipsForTLSPrivateKey2Input = relationshipsForTLSPrivateKey2Input
		u.Type = RelationshipsForTLSPrivateKeyInputTypeRelationshipsForTLSPrivateKey2Input
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RelationshipsForTLSPrivateKeyInput) MarshalJSON() ([]byte, error) {
	if u.RelationshipTLSActivationsInput != nil {
		return utils.MarshalJSON(u.RelationshipTLSActivationsInput, "", true)
	}

	if u.RelationshipsForTLSPrivateKey2Input != nil {
		return utils.MarshalJSON(u.RelationshipsForTLSPrivateKey2Input, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
