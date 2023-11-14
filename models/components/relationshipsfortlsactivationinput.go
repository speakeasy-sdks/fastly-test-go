// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type RelationshipsForTLSActivationTLSCertificateInput struct {
	Data *RelationshipMemberTLSCertificateInput `json:"data,omitempty"`
}

func (o *RelationshipsForTLSActivationTLSCertificateInput) GetData() *RelationshipMemberTLSCertificateInput {
	if o == nil {
		return nil
	}
	return o.Data
}

type RelationshipsForTLSActivationTLSDomain struct {
	Data *RelationshipMemberTLSDomainInput `json:"data,omitempty"`
}

func (o *RelationshipsForTLSActivationTLSDomain) GetData() *RelationshipMemberTLSDomainInput {
	if o == nil {
		return nil
	}
	return o.Data
}

// RelationshipsForTLSActivationInput - The [TLS domain](/reference/api/tls/custom-certs/domains/) being enabled for TLS traffic. Required.
type RelationshipsForTLSActivationInput struct {
	TLSCertificate *RelationshipsForTLSActivationTLSCertificateInput `json:"tls_certificate,omitempty"`
	TLSDomain      *RelationshipsForTLSActivationTLSDomain           `json:"tls_domain,omitempty"`
}

func (o *RelationshipsForTLSActivationInput) GetTLSCertificate() *RelationshipsForTLSActivationTLSCertificateInput {
	if o == nil {
		return nil
	}
	return o.TLSCertificate
}

func (o *RelationshipsForTLSActivationInput) GetTLSDomain() *RelationshipsForTLSActivationTLSDomain {
	if o == nil {
		return nil
	}
	return o.TLSDomain
}
