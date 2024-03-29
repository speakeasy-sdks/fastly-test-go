// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type RelationshipMemberWafTag struct {
	// Alphanumeric string identifying a WAF tag.
	ID *string `json:"id,omitempty"`
	// Resource type.
	Type *TypeWafTag `default:"waf_tag" json:"type"`
}

func (r RelationshipMemberWafTag) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationshipMemberWafTag) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RelationshipMemberWafTag) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RelationshipMemberWafTag) GetType() *TypeWafTag {
	if o == nil {
		return nil
	}
	return o.Type
}
