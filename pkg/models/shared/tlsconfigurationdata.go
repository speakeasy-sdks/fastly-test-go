// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/types"
	"Fastly/pkg/utils"
)

type TLSConfigurationDataAttributes struct {
	// A custom name for your TLS configuration.
	Name *string `json:"name,omitempty"`
}

func (o *TLSConfigurationDataAttributes) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type TLSConfigurationData struct {
	Attributes    *TLSConfigurationDataAttributes   `json:"attributes,omitempty"`
	Relationships *RelationshipsForTLSConfiguration `json:"relationships,omitempty"`
	// Resource type
	type_ *string `const:"tls_configuration" json:"type"`
}

func (t TLSConfigurationData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSConfigurationData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSConfigurationData) GetAttributes() *TLSConfigurationDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *TLSConfigurationData) GetRelationships() *RelationshipsForTLSConfiguration {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *TLSConfigurationData) GetType() *string {
	return types.String("tls_configuration")
}
