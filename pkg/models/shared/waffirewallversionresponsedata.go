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
