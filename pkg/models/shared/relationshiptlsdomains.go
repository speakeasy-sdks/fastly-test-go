// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RelationshipTLSDomainsTLSDomainsInput struct {
	Data []RelationshipMemberTLSDomainInput `json:"data,omitempty"`
}

type RelationshipTLSDomainsInput struct {
	TLSDomains *RelationshipTLSDomainsTLSDomainsInput `json:"tls_domains,omitempty"`
}

type RelationshipTLSDomainsTLSDomains struct {
	Data []RelationshipMemberTLSDomain `json:"data,omitempty"`
}

type RelationshipTLSDomains struct {
	TLSDomains *RelationshipTLSDomainsTLSDomains `json:"tls_domains,omitempty"`
}