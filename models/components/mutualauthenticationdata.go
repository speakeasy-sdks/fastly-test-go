// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
)

type MutualAuthenticationDataAttributes struct {
	// One or more certificates. Enter each individual certificate blob on a new line. Must be PEM-formatted. Required on create. You may optionally rotate the cert_bundle on update.
	CertBundle *string `json:"cert_bundle,omitempty"`
	// Determines whether Mutual TLS will fail closed (enforced) or fail open. A true value will require a successful Mutual TLS handshake for the connection to continue and will fail closed if unsuccessful. A false value will fail open and allow the connection to proceed. Optional. Defaults to true.
	Enforced *bool `json:"enforced,omitempty"`
	// A custom name for your mutual authentication. Optional. If name is not supplied we will auto-generate one.
	Name *string `json:"name,omitempty"`
}

func (o *MutualAuthenticationDataAttributes) GetCertBundle() *string {
	if o == nil {
		return nil
	}
	return o.CertBundle
}

func (o *MutualAuthenticationDataAttributes) GetEnforced() *bool {
	if o == nil {
		return nil
	}
	return o.Enforced
}

func (o *MutualAuthenticationDataAttributes) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type MutualAuthenticationData struct {
	Attributes    *MutualAuthenticationDataAttributes        `json:"attributes,omitempty"`
	Relationships *RelationshipsForMutualAuthenticationInput `json:"relationships,omitempty"`
	// Resource type
	Type *TypeMutualAuthentication `default:"mutual_authentication" json:"type"`
}

func (m MutualAuthenticationData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MutualAuthenticationData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MutualAuthenticationData) GetAttributes() *MutualAuthenticationDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *MutualAuthenticationData) GetRelationships() *RelationshipsForMutualAuthenticationInput {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *MutualAuthenticationData) GetType() *TypeMutualAuthentication {
	if o == nil {
		return nil
	}
	return o.Type
}
