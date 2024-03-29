// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type TLSConfigurationResponseData struct {
	Attributes    *TLSConfigurationResponseAttributes `json:"attributes,omitempty"`
	ID            *string                             `json:"id,omitempty"`
	Relationships *RelationshipsForTLSConfiguration   `json:"relationships,omitempty"`
	// Resource type
	Type *TypeTLSConfiguration `default:"tls_configuration" json:"type"`
}

func (t TLSConfigurationResponseData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSConfigurationResponseData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSConfigurationResponseData) GetAttributes() *TLSConfigurationResponseAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *TLSConfigurationResponseData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TLSConfigurationResponseData) GetRelationships() *RelationshipsForTLSConfiguration {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *TLSConfigurationResponseData) GetType() *TypeTLSConfiguration {
	if o == nil {
		return nil
	}
	return o.Type
}
