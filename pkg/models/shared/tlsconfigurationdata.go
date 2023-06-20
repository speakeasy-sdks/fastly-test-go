// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TLSConfigurationDataAttributes struct {
	// A custom name for your TLS configuration.
	Name *string `json:"name,omitempty"`
}

type TLSConfigurationDataRelationships2DNSRecordsInput struct {
	Data []RelationshipMemberTLSDNSRecordInput `json:"data,omitempty"`
}

// TLSConfigurationDataRelationships2Input - The [DNS records](/reference/api/tls/custom-certs/dns-records/) to use for this configuration.
type TLSConfigurationDataRelationships2Input struct {
	DNSRecords *TLSConfigurationDataRelationships2DNSRecordsInput `json:"dns_records,omitempty"`
}

// TLSConfigurationDataRelationships1Input - The [Fastly Service](/reference/api/services/service/) that is automatically selected when this TLS Configuration is used.
type TLSConfigurationDataRelationships1Input struct {
	Service *RelationshipMemberServiceInput `json:"service,omitempty"`
}

type TLSConfigurationDataInput struct {
	Attributes    *TLSConfigurationDataAttributes `json:"attributes,omitempty"`
	Relationships interface{}                     `json:"relationships,omitempty"`
	// Resource type
	Type *TypeTLSConfiguration `json:"type,omitempty"`
}
