// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type WafTagAttributes struct {
	Name *string `json:"name,omitempty"`
}

func (o *WafTagAttributes) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type WafTag struct {
	Attributes *WafTagAttributes `json:"attributes,omitempty"`
	// Alphanumeric string identifying a WAF tag.
	ID *string `json:"id,omitempty"`
	// Resource type.
	Type *TypeWafTag `default:"waf_tag" json:"type"`
}

func (w WafTag) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafTag) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *WafTag) GetAttributes() *WafTagAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *WafTag) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WafTag) GetType() *TypeWafTag {
	if o == nil {
		return nil
	}
	return o.Type
}
