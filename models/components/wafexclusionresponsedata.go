// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
)

type WafExclusionResponseData struct {
	Attributes *WafExclusionResponseDataAttributes `json:"attributes,omitempty"`
	// Alphanumeric string identifying a WAF exclusion.
	ID            *string                                `json:"id,omitempty"`
	Relationships *WafExclusionResponseDataRelationships `json:"relationships,omitempty"`
	// Resource type.
	Type *TypeWafExclusion `default:"waf_exclusion" json:"type"`
}

func (w WafExclusionResponseData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafExclusionResponseData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafExclusionResponseData) GetAttributes() *WafExclusionResponseDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *WafExclusionResponseData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WafExclusionResponseData) GetRelationships() *WafExclusionResponseDataRelationships {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *WafExclusionResponseData) GetType() *TypeWafExclusion {
	if o == nil {
		return nil
	}
	return o.Type
}
