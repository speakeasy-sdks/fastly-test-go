// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WafFirewallVersionResponseData struct {
	Attributes *WafFirewallVersionResponseDataAttributes `json:"attributes,omitempty"`
	// Alphanumeric string identifying a Firewall version.
	ID            *string     `json:"id,omitempty"`
	Relationships interface{} `json:"relationships,omitempty"`
	// Resource type.
	Type *TypeWafFirewallVersion `json:"type,omitempty"`
}

func (o *WafFirewallVersionResponseData) GetAttributes() *WafFirewallVersionResponseDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *WafFirewallVersionResponseData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WafFirewallVersionResponseData) GetRelationships() interface{} {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *WafFirewallVersionResponseData) GetType() *TypeWafFirewallVersion {
	if o == nil {
		return nil
	}
	return o.Type
}
