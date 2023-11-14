// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
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
	Attributes    *TLSConfigurationDataAttributes        `json:"attributes,omitempty"`
	Relationships *RelationshipsForTLSConfigurationInput `json:"relationships,omitempty"`
	// Resource type
	Type *TypeTLSConfiguration `default:"tls_configuration" json:"type"`
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

func (o *TLSConfigurationData) GetRelationships() *RelationshipsForTLSConfigurationInput {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *TLSConfigurationData) GetType() *TypeTLSConfiguration {
	if o == nil {
		return nil
	}
	return o.Type
}