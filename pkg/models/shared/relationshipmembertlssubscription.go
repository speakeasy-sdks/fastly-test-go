// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/types"
	"Fastly/pkg/utils"
)

type RelationshipMemberTLSSubscription struct {
	ID *string `json:"id,omitempty"`
	// Resource type
	type_ *string `const:"tls_subscription" json:"type"`
}

func (r RelationshipMemberTLSSubscription) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationshipMemberTLSSubscription) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RelationshipMemberTLSSubscription) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RelationshipMemberTLSSubscription) GetType() *string {
	return types.String("tls_subscription")
}
