// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RelationshipMemberWafFirewallVersion struct {
	// Alphanumeric string identifying a Firewall version.
	ID *string `json:"id,omitempty"`
	// Resource type.
	Type *TypeWafFirewallVersion `json:"type,omitempty"`
}

func (o *RelationshipMemberWafFirewallVersion) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RelationshipMemberWafFirewallVersion) GetType() *TypeWafFirewallVersion {
	if o == nil {
		return nil
	}
	return o.Type
}

type RelationshipMemberWafFirewallVersionInput struct {
	// Resource type.
	Type *TypeWafFirewallVersion `json:"type,omitempty"`
}

func (o *RelationshipMemberWafFirewallVersionInput) GetType() *TypeWafFirewallVersion {
	if o == nil {
		return nil
	}
	return o.Type
}
