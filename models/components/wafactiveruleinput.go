// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type WafActiveRuleInput struct {
	Data *WafActiveRuleDataInput `json:"data,omitempty"`
}

func (o *WafActiveRuleInput) GetData() *WafActiveRuleDataInput {
	if o == nil {
		return nil
	}
	return o.Data
}
