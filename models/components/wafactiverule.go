// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type WafActiveRule struct {
	Data *WafActiveRuleData `json:"data,omitempty"`
}

func (o *WafActiveRule) GetData() *WafActiveRuleData {
	if o == nil {
		return nil
	}
	return o.Data
}