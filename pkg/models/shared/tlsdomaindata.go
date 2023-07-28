// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TLSDomainDataRelationships2TLSCertificates struct {
	Data []RelationshipMemberTLSCertificate `json:"data,omitempty"`
}

func (o *TLSDomainDataRelationships2TLSCertificates) GetData() []RelationshipMemberTLSCertificate {
	if o == nil {
		return nil
	}
	return o.Data
}

// TLSDomainDataRelationships2 - The list of all the [TLS certificates](#tls_certificates) that include this domain in their SAN list.
type TLSDomainDataRelationships2 struct {
	TLSCertificates *TLSDomainDataRelationships2TLSCertificates `json:"tls_certificates,omitempty"`
}

func (o *TLSDomainDataRelationships2) GetTLSCertificates() *TLSDomainDataRelationships2TLSCertificates {
	if o == nil {
		return nil
	}
	return o.TLSCertificates
}

type TLSDomainDataRelationships1TLSActivations struct {
	Data []RelationshipMemberTLSActivation `json:"data,omitempty"`
}

func (o *TLSDomainDataRelationships1TLSActivations) GetData() []RelationshipMemberTLSActivation {
	if o == nil {
		return nil
	}
	return o.Data
}

// TLSDomainDataRelationships1 - The list of [TLS activations](#tls_activations) that exist for the domain. If empty, then this domain is not enabled to serve TLS traffic.
type TLSDomainDataRelationships1 struct {
	TLSActivations *TLSDomainDataRelationships1TLSActivations `json:"tls_activations,omitempty"`
}

func (o *TLSDomainDataRelationships1) GetTLSActivations() *TLSDomainDataRelationships1TLSActivations {
	if o == nil {
		return nil
	}
	return o.TLSActivations
}

type TLSDomainData struct {
	// The domain name.
	ID            *string     `json:"id,omitempty"`
	Relationships interface{} `json:"relationships,omitempty"`
	// Resource type
	Type *TypeTLSDomain `json:"type,omitempty"`
}

func (o *TLSDomainData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TLSDomainData) GetRelationships() interface{} {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *TLSDomainData) GetType() *TypeTLSDomain {
	if o == nil {
		return nil
	}
	return o.Type
}
