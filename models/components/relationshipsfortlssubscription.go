// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type RelationshipsForTLSSubscriptionType string

const (
	RelationshipsForTLSSubscriptionTypeRelationshipTLSConfiguration     RelationshipsForTLSSubscriptionType = "relationship_tls_configuration"
	RelationshipsForTLSSubscriptionTypeRelationshipCommonName           RelationshipsForTLSSubscriptionType = "relationship_common_name"
	RelationshipsForTLSSubscriptionTypeRelationshipTLSDomainsInput      RelationshipsForTLSSubscriptionType = "relationship_tls_domains_input"
	RelationshipsForTLSSubscriptionTypeRelationshipTLSCertificatesInput RelationshipsForTLSSubscriptionType = "relationship_tls_certificates_input"
)

type RelationshipsForTLSSubscription struct {
	RelationshipTLSConfiguration     *RelationshipTLSConfiguration
	RelationshipCommonName           *RelationshipCommonName
	RelationshipTLSDomainsInput      *RelationshipTLSDomainsInput
	RelationshipTLSCertificatesInput *RelationshipTLSCertificatesInput

	Type RelationshipsForTLSSubscriptionType
}

func CreateRelationshipsForTLSSubscriptionRelationshipTLSConfiguration(relationshipTLSConfiguration RelationshipTLSConfiguration) RelationshipsForTLSSubscription {
	typ := RelationshipsForTLSSubscriptionTypeRelationshipTLSConfiguration

	return RelationshipsForTLSSubscription{
		RelationshipTLSConfiguration: &relationshipTLSConfiguration,
		Type:                         typ,
	}
}

func CreateRelationshipsForTLSSubscriptionRelationshipCommonName(relationshipCommonName RelationshipCommonName) RelationshipsForTLSSubscription {
	typ := RelationshipsForTLSSubscriptionTypeRelationshipCommonName

	return RelationshipsForTLSSubscription{
		RelationshipCommonName: &relationshipCommonName,
		Type:                   typ,
	}
}

func CreateRelationshipsForTLSSubscriptionRelationshipTLSDomainsInput(relationshipTLSDomainsInput RelationshipTLSDomainsInput) RelationshipsForTLSSubscription {
	typ := RelationshipsForTLSSubscriptionTypeRelationshipTLSDomainsInput

	return RelationshipsForTLSSubscription{
		RelationshipTLSDomainsInput: &relationshipTLSDomainsInput,
		Type:                        typ,
	}
}

func CreateRelationshipsForTLSSubscriptionRelationshipTLSCertificatesInput(relationshipTLSCertificatesInput RelationshipTLSCertificatesInput) RelationshipsForTLSSubscription {
	typ := RelationshipsForTLSSubscriptionTypeRelationshipTLSCertificatesInput

	return RelationshipsForTLSSubscription{
		RelationshipTLSCertificatesInput: &relationshipTLSCertificatesInput,
		Type:                             typ,
	}
}

func (u *RelationshipsForTLSSubscription) UnmarshalJSON(data []byte) error {

	relationshipTLSConfiguration := RelationshipTLSConfiguration{}
	if err := utils.UnmarshalJSON(data, &relationshipTLSConfiguration, "", true, true); err == nil {
		u.RelationshipTLSConfiguration = &relationshipTLSConfiguration
		u.Type = RelationshipsForTLSSubscriptionTypeRelationshipTLSConfiguration
		return nil
	}

	relationshipCommonName := RelationshipCommonName{}
	if err := utils.UnmarshalJSON(data, &relationshipCommonName, "", true, true); err == nil {
		u.RelationshipCommonName = &relationshipCommonName
		u.Type = RelationshipsForTLSSubscriptionTypeRelationshipCommonName
		return nil
	}

	relationshipTLSDomainsInput := RelationshipTLSDomainsInput{}
	if err := utils.UnmarshalJSON(data, &relationshipTLSDomainsInput, "", true, true); err == nil {
		u.RelationshipTLSDomainsInput = &relationshipTLSDomainsInput
		u.Type = RelationshipsForTLSSubscriptionTypeRelationshipTLSDomainsInput
		return nil
	}

	relationshipTLSCertificatesInput := RelationshipTLSCertificatesInput{}
	if err := utils.UnmarshalJSON(data, &relationshipTLSCertificatesInput, "", true, true); err == nil {
		u.RelationshipTLSCertificatesInput = &relationshipTLSCertificatesInput
		u.Type = RelationshipsForTLSSubscriptionTypeRelationshipTLSCertificatesInput
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RelationshipsForTLSSubscription) MarshalJSON() ([]byte, error) {
	if u.RelationshipTLSConfiguration != nil {
		return utils.MarshalJSON(u.RelationshipTLSConfiguration, "", true)
	}

	if u.RelationshipCommonName != nil {
		return utils.MarshalJSON(u.RelationshipCommonName, "", true)
	}

	if u.RelationshipTLSDomainsInput != nil {
		return utils.MarshalJSON(u.RelationshipTLSDomainsInput, "", true)
	}

	if u.RelationshipTLSCertificatesInput != nil {
		return utils.MarshalJSON(u.RelationshipTLSCertificatesInput, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
