// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type RelationshipMemberTLSActivationInput struct {
	// Resource type.
	Type *TypeTLSActivation `default:"tls_activation" json:"type"`
}

func (r RelationshipMemberTLSActivationInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationshipMemberTLSActivationInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RelationshipMemberTLSActivationInput) GetType() *TypeTLSActivation {
	if o == nil {
		return nil
	}
	return o.Type
}
