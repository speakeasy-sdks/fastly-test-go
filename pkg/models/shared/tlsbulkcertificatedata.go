// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TLSBulkCertificateDataAttributes struct {
	// Allow certificates that chain to untrusted roots.
	AllowUntrustedRoot *bool `json:"allow_untrusted_root,omitempty"`
	// The PEM-formatted certificate blob. Required.
	CertBlob *string `json:"cert_blob,omitempty"`
	// The PEM-formatted chain of intermediate blobs. Required.
	IntermediatesBlob *string `json:"intermediates_blob,omitempty"`
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

type TLSBulkCertificateDataRelationships2TLSDomainsInput struct {
	Data []RelationshipMemberTLSDomainInput `json:"data,omitempty"`
}

func (o *TLSBulkCertificateDataRelationships2TLSDomainsInput) GetData() []RelationshipMemberTLSDomainInput {
	if o == nil {
		return nil
	}
	return o.Data
}

// TLSBulkCertificateDataRelationships2Input - All the domains (including wildcard domains) that are listed in any certificate's Subject Alternative Names (SAN) list.
type TLSBulkCertificateDataRelationships2Input struct {
	TLSDomains *TLSBulkCertificateDataRelationships2TLSDomainsInput `json:"tls_domains,omitempty"`
}

func (o *TLSBulkCertificateDataRelationships2Input) GetTLSDomains() *TLSBulkCertificateDataRelationships2TLSDomainsInput {
	if o == nil {
		return nil
	}
	return o.TLSDomains
}

type TLSBulkCertificateDataInput struct {
	Attributes    *TLSBulkCertificateDataAttributes `json:"attributes,omitempty"`
	Relationships interface{}                       `json:"relationships,omitempty"`
	// Resource type
	Type *TypeTLSBulkCertificate `json:"type,omitempty"`
}

func (o *TLSBulkCertificateDataInput) GetAttributes() *TLSBulkCertificateDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *TLSBulkCertificateDataInput) GetRelationships() interface{} {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *TLSBulkCertificateDataInput) GetType() *TypeTLSBulkCertificate {
	if o == nil {
		return nil
	}
	return o.Type
}
