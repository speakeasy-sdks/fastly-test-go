// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RelationshipTLSActivationsTLSActivationsInput struct {
	Data []RelationshipMemberTLSActivationInput `json:"data,omitempty"`
}

type RelationshipTLSActivationsInput struct {
	TLSActivations *RelationshipTLSActivationsTLSActivationsInput `json:"tls_activations,omitempty"`
}

type RelationshipTLSActivationsTLSActivations struct {
	Data []RelationshipMemberTLSActivation `json:"data,omitempty"`
}

type RelationshipTLSActivations struct {
	TLSActivations *RelationshipTLSActivationsTLSActivations `json:"tls_activations,omitempty"`
}
