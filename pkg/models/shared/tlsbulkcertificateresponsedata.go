// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TLSBulkCertificateResponseDataRelationships2TLSDomains struct {
	Data []RelationshipMemberTLSDomain `json:"data,omitempty"`
}

// TLSBulkCertificateResponseDataRelationships2 - All the domains (including wildcard domains) that are listed in any certificate's Subject Alternative Names (SAN) list.
type TLSBulkCertificateResponseDataRelationships2 struct {
	TLSDomains *TLSBulkCertificateResponseDataRelationships2TLSDomains `json:"tls_domains,omitempty"`
}

type TLSBulkCertificateResponseData struct {
	Attributes    *TLSBulkCertificateResponseAttributes `json:"attributes,omitempty"`
	ID            *string                               `json:"id,omitempty"`
	Relationships interface{}                           `json:"relationships,omitempty"`
	// Resource type
	Type *TypeTLSBulkCertificate `json:"type,omitempty"`
}
