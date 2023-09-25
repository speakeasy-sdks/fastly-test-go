// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/types"
	"Fastly/pkg/utils"
)

type RelationshipMemberService struct {
	ID *string `json:"id,omitempty"`
	// Resource type
	type_ *string `const:"service" json:"type"`
}

func (r RelationshipMemberService) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationshipMemberService) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RelationshipMemberService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RelationshipMemberService) GetType() *string {
	return types.String("service")
}

type RelationshipMemberServiceInput struct {
	// Resource type
	type_ *string `const:"service" json:"type"`
}

func (r RelationshipMemberServiceInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationshipMemberServiceInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RelationshipMemberServiceInput) GetType() *string {
	return types.String("service")
}
