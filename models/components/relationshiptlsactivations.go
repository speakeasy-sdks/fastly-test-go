// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type RelationshipTLSActivationsTLSActivations struct {
	Data []RelationshipMemberTLSActivationInput `json:"data,omitempty"`
}

func (o *RelationshipTLSActivationsTLSActivations) GetData() []RelationshipMemberTLSActivationInput {
	if o == nil {
		return nil
	}
	return o.Data
}

type RelationshipTLSActivationsInput struct {
	TLSActivations *RelationshipTLSActivationsTLSActivations `json:"tls_activations,omitempty"`
}

func (o *RelationshipTLSActivationsInput) GetTLSActivations() *RelationshipTLSActivationsTLSActivations {
	if o == nil {
		return nil
	}
	return o.TLSActivations
}

type TLSActivations struct {
	Data []RelationshipMemberTLSActivation `json:"data,omitempty"`
}

func (o *TLSActivations) GetData() []RelationshipMemberTLSActivation {
	if o == nil {
		return nil
	}
	return o.Data
}

type RelationshipTLSActivations struct {
	TLSActivations *TLSActivations `json:"tls_activations,omitempty"`
}

func (o *RelationshipTLSActivations) GetTLSActivations() *TLSActivations {
	if o == nil {
		return nil
	}
	return o.TLSActivations
}
