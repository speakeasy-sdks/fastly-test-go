// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type WafExclusion struct {
	Data *WafExclusionData `json:"data,omitempty"`
}

func (o *WafExclusion) GetData() *WafExclusionData {
	if o == nil {
		return nil
	}
	return o.Data
}
