// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TLSCertificateResponseData struct {
	Attributes    *TLSCertificateResponseAttributes `json:"attributes,omitempty"`
	ID            *string                           `json:"id,omitempty"`
	Relationships *RelationshipTLSDomains           `json:"relationships,omitempty"`
	// Resource type
	Type *TypeTLSCertificate `json:"type,omitempty"`
}