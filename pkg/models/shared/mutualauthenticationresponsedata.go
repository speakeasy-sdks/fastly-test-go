// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/types"
	"Fastly/pkg/utils"
)

type MutualAuthenticationResponseData struct {
	Attributes    *MutualAuthenticationResponseAttributes `json:"attributes,omitempty"`
	ID            *string                                 `json:"id,omitempty"`
	Relationships *RelationshipsForMutualAuthentication   `json:"relationships,omitempty"`
	// Resource type
	type_ *string `const:"mutual_authentication" json:"type"`
}

func (m MutualAuthenticationResponseData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MutualAuthenticationResponseData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MutualAuthenticationResponseData) GetAttributes() *MutualAuthenticationResponseAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *MutualAuthenticationResponseData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MutualAuthenticationResponseData) GetRelationships() *RelationshipsForMutualAuthentication {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *MutualAuthenticationResponseData) GetType() *string {
	return types.String("mutual_authentication")
}
