// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type SchemasWafFirewallVersion struct {
	Data *SchemasWafFirewallVersionData `json:"data,omitempty"`
}

func (o *SchemasWafFirewallVersion) GetData() *SchemasWafFirewallVersionData {
	if o == nil {
		return nil
	}
	return o.Data
}