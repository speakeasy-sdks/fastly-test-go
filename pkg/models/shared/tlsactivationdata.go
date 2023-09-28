// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/types"
	"Fastly/pkg/utils"
)

type TLSActivationDataInput struct {
	Relationships *RelationshipsForTLSActivationInput `json:"relationships,omitempty"`
	// Resource type.
	type_ *string `const:"tls_activation" json:"type"`
}

func (t TLSActivationDataInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSActivationDataInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSActivationDataInput) GetRelationships() *RelationshipsForTLSActivationInput {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *TLSActivationDataInput) GetType() *string {
	return types.String("tls_activation")
}
