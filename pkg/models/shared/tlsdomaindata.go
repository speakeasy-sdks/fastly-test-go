// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/types"
	"Fastly/pkg/utils"
)

type TLSDomainData struct {
	// The domain name.
	ID            *string                       `json:"id,omitempty"`
	Relationships *RelationshipTLSSubscriptions `json:"relationships,omitempty"`
	// Resource type
	type_ *string `const:"tls_domain" json:"type"`
}

func (t TLSDomainData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSDomainData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSDomainData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TLSDomainData) GetRelationships() *RelationshipTLSSubscriptions {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *TLSDomainData) GetType() *string {
	return types.String("tls_domain")
}
