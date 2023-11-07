// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
)

type TLSPrivateKeyResponseData struct {
	Attributes *TLSPrivateKeyResponseAttributes `json:"attributes,omitempty"`
	ID         *string                          `json:"id,omitempty"`
	// Resource type
	Type *TypeTLSPrivateKey `default:"tls_private_key" json:"type"`
}

func (t TLSPrivateKeyResponseData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSPrivateKeyResponseData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSPrivateKeyResponseData) GetAttributes() *TLSPrivateKeyResponseAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *TLSPrivateKeyResponseData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TLSPrivateKeyResponseData) GetType() *TypeTLSPrivateKey {
	if o == nil {
		return nil
	}
	return o.Type
}
