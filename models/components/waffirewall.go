// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type WafFirewall struct {
	Data *WafFirewallData `json:"data,omitempty"`
}

func (o *WafFirewall) GetData() *WafFirewallData {
	if o == nil {
		return nil
	}
	return o.Data
}
