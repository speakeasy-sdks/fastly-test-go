// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type BulkWafActiveRulesInput struct {
	Data []WafActiveRuleDataInput `json:"data,omitempty"`
}

func (o *BulkWafActiveRulesInput) GetData() []WafActiveRuleDataInput {
	if o == nil {
		return nil
	}
	return o.Data
}
