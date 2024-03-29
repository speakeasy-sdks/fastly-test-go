// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type TLSCertificateResponseData struct {
	Attributes    *TLSCertificateResponseAttributes `json:"attributes,omitempty"`
	ID            *string                           `json:"id,omitempty"`
	Relationships *RelationshipTLSDomains           `json:"relationships,omitempty"`
	// Resource type
	Type *TypeTLSCertificate `default:"tls_certificate" json:"type"`
}

func (t TLSCertificateResponseData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSCertificateResponseData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSCertificateResponseData) GetAttributes() *TLSCertificateResponseAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *TLSCertificateResponseData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TLSCertificateResponseData) GetRelationships() *RelationshipTLSDomains {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *TLSCertificateResponseData) GetType() *TypeTLSCertificate {
	if o == nil {
		return nil
	}
	return o.Type
}
