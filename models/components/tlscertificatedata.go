// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
)

type TLSCertificateDataAttributes struct {
	// The PEM-formatted certificate blob. Required.
	CertBlob *string `json:"cert_blob,omitempty"`
	// A customizable name for your certificate. Defaults to the certificate's Common Name or first Subject Alternative Names (SAN) entry. Optional.
	Name *string `json:"name,omitempty"`
}

func (o *TLSCertificateDataAttributes) GetCertBlob() *string {
	if o == nil {
		return nil
	}
	return o.CertBlob
}

func (o *TLSCertificateDataAttributes) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type TLSCertificateData struct {
	Attributes    *TLSCertificateDataAttributes `json:"attributes,omitempty"`
	Relationships *RelationshipTLSDomainsInput  `json:"relationships,omitempty"`
	// Resource type
	Type *TypeTLSCertificate `default:"tls_certificate" json:"type"`
}

func (t TLSCertificateData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSCertificateData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSCertificateData) GetAttributes() *TLSCertificateDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *TLSCertificateData) GetRelationships() *RelationshipTLSDomainsInput {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *TLSCertificateData) GetType() *TypeTLSCertificate {
	if o == nil {
		return nil
	}
	return o.Type
}
