// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"errors"
)

type RelationshipsForTLSSubscriptionInputType string

const (
	RelationshipsForTLSSubscriptionInputTypeRelationshipTLSConfigurationInput RelationshipsForTLSSubscriptionInputType = "relationship_tls_configurationInput"
	RelationshipsForTLSSubscriptionInputTypeRelationshipCommonNameInput       RelationshipsForTLSSubscriptionInputType = "relationship_common_nameInput"
	RelationshipsForTLSSubscriptionInputTypeRelationshipTLSDomainsInput       RelationshipsForTLSSubscriptionInputType = "relationship_tls_domainsInput"
	RelationshipsForTLSSubscriptionInputTypeRelationshipTLSCertificatesInput  RelationshipsForTLSSubscriptionInputType = "relationship_tls_certificatesInput"
)

type RelationshipsForTLSSubscriptionInput struct {
	RelationshipTLSConfigurationInput *RelationshipTLSConfigurationInput
	RelationshipCommonNameInput       *RelationshipCommonNameInput
	RelationshipTLSDomainsInput       *RelationshipTLSDomainsInput
	RelationshipTLSCertificatesInput  *RelationshipTLSCertificatesInput

	Type RelationshipsForTLSSubscriptionInputType
}

func CreateRelationshipsForTLSSubscriptionInputRelationshipTLSConfigurationInput(relationshipTLSConfigurationInput RelationshipTLSConfigurationInput) RelationshipsForTLSSubscriptionInput {
	typ := RelationshipsForTLSSubscriptionInputTypeRelationshipTLSConfigurationInput

	return RelationshipsForTLSSubscriptionInput{
		RelationshipTLSConfigurationInput: &relationshipTLSConfigurationInput,
		Type:                              typ,
	}
}

func CreateRelationshipsForTLSSubscriptionInputRelationshipCommonNameInput(relationshipCommonNameInput RelationshipCommonNameInput) RelationshipsForTLSSubscriptionInput {
	typ := RelationshipsForTLSSubscriptionInputTypeRelationshipCommonNameInput

	return RelationshipsForTLSSubscriptionInput{
		RelationshipCommonNameInput: &relationshipCommonNameInput,
		Type:                        typ,
	}
}

func CreateRelationshipsForTLSSubscriptionInputRelationshipTLSDomainsInput(relationshipTLSDomainsInput RelationshipTLSDomainsInput) RelationshipsForTLSSubscriptionInput {
	typ := RelationshipsForTLSSubscriptionInputTypeRelationshipTLSDomainsInput

	return RelationshipsForTLSSubscriptionInput{
		RelationshipTLSDomainsInput: &relationshipTLSDomainsInput,
		Type:                        typ,
	}
}

func CreateRelationshipsForTLSSubscriptionInputRelationshipTLSCertificatesInput(relationshipTLSCertificatesInput RelationshipTLSCertificatesInput) RelationshipsForTLSSubscriptionInput {
	typ := RelationshipsForTLSSubscriptionInputTypeRelationshipTLSCertificatesInput

	return RelationshipsForTLSSubscriptionInput{
		RelationshipTLSCertificatesInput: &relationshipTLSCertificatesInput,
		Type:                             typ,
	}
}

func (u *RelationshipsForTLSSubscriptionInput) UnmarshalJSON(data []byte) error {

	relationshipTLSConfigurationInput := RelationshipTLSConfigurationInput{}
	if err := utils.UnmarshalJSON(data, &relationshipTLSConfigurationInput, "", true, true); err == nil {
		u.RelationshipTLSConfigurationInput = &relationshipTLSConfigurationInput
		u.Type = RelationshipsForTLSSubscriptionInputTypeRelationshipTLSConfigurationInput
		return nil
	}

	relationshipCommonNameInput := RelationshipCommonNameInput{}
	if err := utils.UnmarshalJSON(data, &relationshipCommonNameInput, "", true, true); err == nil {
		u.RelationshipCommonNameInput = &relationshipCommonNameInput
		u.Type = RelationshipsForTLSSubscriptionInputTypeRelationshipCommonNameInput
		return nil
	}

	relationshipTLSDomainsInput := RelationshipTLSDomainsInput{}
	if err := utils.UnmarshalJSON(data, &relationshipTLSDomainsInput, "", true, true); err == nil {
		u.RelationshipTLSDomainsInput = &relationshipTLSDomainsInput
		u.Type = RelationshipsForTLSSubscriptionInputTypeRelationshipTLSDomainsInput
		return nil
	}

	relationshipTLSCertificatesInput := RelationshipTLSCertificatesInput{}
	if err := utils.UnmarshalJSON(data, &relationshipTLSCertificatesInput, "", true, true); err == nil {
		u.RelationshipTLSCertificatesInput = &relationshipTLSCertificatesInput
		u.Type = RelationshipsForTLSSubscriptionInputTypeRelationshipTLSCertificatesInput
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RelationshipsForTLSSubscriptionInput) MarshalJSON() ([]byte, error) {
	if u.RelationshipTLSConfigurationInput != nil {
		return utils.MarshalJSON(u.RelationshipTLSConfigurationInput, "", true)
	}

	if u.RelationshipCommonNameInput != nil {
		return utils.MarshalJSON(u.RelationshipCommonNameInput, "", true)
	}

	if u.RelationshipTLSDomainsInput != nil {
		return utils.MarshalJSON(u.RelationshipTLSDomainsInput, "", true)
	}

	if u.RelationshipTLSCertificatesInput != nil {
		return utils.MarshalJSON(u.RelationshipTLSCertificatesInput, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
