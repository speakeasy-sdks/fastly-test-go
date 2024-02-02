// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type TLSDomainData struct {
	// The domain name.
	ID            *string                    `json:"id,omitempty"`
	Relationships *RelationshipsForTLSDomain `json:"relationships,omitempty"`
	// Resource type
	Type *TypeTLSDomain `default:"tls_domain" json:"type"`
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

func (o *TLSDomainData) GetRelationships() *RelationshipsForTLSDomain {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *TLSDomainData) GetType() *TypeTLSDomain {
	if o == nil {
		return nil
	}
	return o.Type
}
