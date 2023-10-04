// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"errors"
)

type RelationshipsForTLSBulkCertificate2TLSDomainsInput struct {
	Data []RelationshipMemberTLSDomainInput `json:"data,omitempty"`
}

func (o *RelationshipsForTLSBulkCertificate2TLSDomainsInput) GetData() []RelationshipMemberTLSDomainInput {
	if o == nil {
		return nil
	}
	return o.Data
}

// RelationshipsForTLSBulkCertificate2Input - All the domains (including wildcard domains) that are listed in any certificate's Subject Alternative Names (SAN) list.
type RelationshipsForTLSBulkCertificate2Input struct {
	TLSDomains *RelationshipsForTLSBulkCertificate2TLSDomainsInput `json:"tls_domains,omitempty"`
}

func (o *RelationshipsForTLSBulkCertificate2Input) GetTLSDomains() *RelationshipsForTLSBulkCertificate2TLSDomainsInput {
	if o == nil {
		return nil
	}
	return o.TLSDomains
}

type RelationshipsForTLSBulkCertificateInputType string

const (
	RelationshipsForTLSBulkCertificateInputTypeRelationshipTLSConfigurationsInput       RelationshipsForTLSBulkCertificateInputType = "relationship_tls_configurationsInput"
	RelationshipsForTLSBulkCertificateInputTypeRelationshipsForTLSBulkCertificate2Input RelationshipsForTLSBulkCertificateInputType = "relationships_for_tls_bulk_certificate_2Input"
)

type RelationshipsForTLSBulkCertificateInput struct {
	RelationshipTLSConfigurationsInput       *RelationshipTLSConfigurationsInput
	RelationshipsForTLSBulkCertificate2Input *RelationshipsForTLSBulkCertificate2Input

	Type RelationshipsForTLSBulkCertificateInputType
}

func CreateRelationshipsForTLSBulkCertificateInputRelationshipTLSConfigurationsInput(relationshipTLSConfigurationsInput RelationshipTLSConfigurationsInput) RelationshipsForTLSBulkCertificateInput {
	typ := RelationshipsForTLSBulkCertificateInputTypeRelationshipTLSConfigurationsInput

	return RelationshipsForTLSBulkCertificateInput{
		RelationshipTLSConfigurationsInput: &relationshipTLSConfigurationsInput,
		Type:                               typ,
	}
}

func CreateRelationshipsForTLSBulkCertificateInputRelationshipsForTLSBulkCertificate2Input(relationshipsForTLSBulkCertificate2Input RelationshipsForTLSBulkCertificate2Input) RelationshipsForTLSBulkCertificateInput {
	typ := RelationshipsForTLSBulkCertificateInputTypeRelationshipsForTLSBulkCertificate2Input

	return RelationshipsForTLSBulkCertificateInput{
		RelationshipsForTLSBulkCertificate2Input: &relationshipsForTLSBulkCertificate2Input,
		Type:                                     typ,
	}
}

func (u *RelationshipsForTLSBulkCertificateInput) UnmarshalJSON(data []byte) error {

	relationshipTLSConfigurationsInput := new(RelationshipTLSConfigurationsInput)
	if err := utils.UnmarshalJSON(data, &relationshipTLSConfigurationsInput, "", true, true); err == nil {
		u.RelationshipTLSConfigurationsInput = relationshipTLSConfigurationsInput
		u.Type = RelationshipsForTLSBulkCertificateInputTypeRelationshipTLSConfigurationsInput
		return nil
	}

	relationshipsForTLSBulkCertificate2Input := new(RelationshipsForTLSBulkCertificate2Input)
	if err := utils.UnmarshalJSON(data, &relationshipsForTLSBulkCertificate2Input, "", true, true); err == nil {
		u.RelationshipsForTLSBulkCertificate2Input = relationshipsForTLSBulkCertificate2Input
		u.Type = RelationshipsForTLSBulkCertificateInputTypeRelationshipsForTLSBulkCertificate2Input
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RelationshipsForTLSBulkCertificateInput) MarshalJSON() ([]byte, error) {
	if u.RelationshipTLSConfigurationsInput != nil {
		return utils.MarshalJSON(u.RelationshipTLSConfigurationsInput, "", true)
	}

	if u.RelationshipsForTLSBulkCertificate2Input != nil {
		return utils.MarshalJSON(u.RelationshipsForTLSBulkCertificate2Input, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
