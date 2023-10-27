// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
)

type TLSBulkCertificateDataAttributes struct {
	// Allow certificates that chain to untrusted roots.
	AllowUntrustedRoot *bool `default:"false" json:"allow_untrusted_root"`
	// The PEM-formatted certificate blob. Required.
	CertBlob *string `json:"cert_blob,omitempty"`
	// The PEM-formatted chain of intermediate blobs. Required.
	IntermediatesBlob *string `json:"intermediates_blob,omitempty"`
}

func (t TLSBulkCertificateDataAttributes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSBulkCertificateDataAttributes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSBulkCertificateDataAttributes) GetAllowUntrustedRoot() *bool {
	if o == nil {
		return nil
	}
	return o.AllowUntrustedRoot
}

func (o *TLSBulkCertificateDataAttributes) GetCertBlob() *string {
	if o == nil {
		return nil
	}
	return o.CertBlob
}

func (o *TLSBulkCertificateDataAttributes) GetIntermediatesBlob() *string {
	if o == nil {
		return nil
	}
	return o.IntermediatesBlob
}

type TLSBulkCertificateData struct {
	Attributes    *TLSBulkCertificateDataAttributes        `json:"attributes,omitempty"`
	Relationships *RelationshipsForTLSBulkCertificateInput `json:"relationships,omitempty"`
	// Resource type
	Type *TypeTLSBulkCertificate `default:"tls_bulk_certificate" json:"type"`
}

func (t TLSBulkCertificateData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSBulkCertificateData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSBulkCertificateData) GetAttributes() *TLSBulkCertificateDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *TLSBulkCertificateData) GetRelationships() *RelationshipsForTLSBulkCertificateInput {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *TLSBulkCertificateData) GetType() *TypeTLSBulkCertificate {
	if o == nil {
		return nil
	}
	return o.Type
}
