// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type WafExclusionResponse struct {
	Data *WafExclusionResponseData `json:"data,omitempty"`
}

func (o *WafExclusionResponse) GetData() *WafExclusionResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}