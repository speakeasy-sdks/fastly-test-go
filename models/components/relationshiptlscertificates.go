// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type TLSCertificates struct {
	Data []RelationshipMemberTLSCertificate `json:"data,omitempty"`
}

func (o *TLSCertificates) GetData() []RelationshipMemberTLSCertificate {
	if o == nil {
		return nil
	}
	return o.Data
}

type RelationshipTLSCertificates struct {
	TLSCertificates *TLSCertificates `json:"tls_certificates,omitempty"`
}

func (o *RelationshipTLSCertificates) GetTLSCertificates() *TLSCertificates {
	if o == nil {
		return nil
	}
	return o.TLSCertificates
}
