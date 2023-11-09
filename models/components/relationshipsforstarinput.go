// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
)

type RelationshipsForStarDataInput struct {
	// Resource type
	Type *TypeUser `default:"user" json:"type"`
}

func (r RelationshipsForStarDataInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationshipsForStarDataInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RelationshipsForStarDataInput) GetType() *TypeUser {
	if o == nil {
		return nil
	}
	return o.Type
}

type RelationshipsForStarUserInput struct {
	Data *RelationshipsForStarDataInput `json:"data,omitempty"`
}

func (o *RelationshipsForStarUserInput) GetData() *RelationshipsForStarDataInput {
	if o == nil {
		return nil
	}
	return o.Data
}

type RelationshipsForStarInput struct {
	Service *RelationshipMemberServiceInput `json:"service,omitempty"`
	User    *RelationshipsForStarUserInput  `json:"user,omitempty"`
}

func (o *RelationshipsForStarInput) GetService() *RelationshipMemberServiceInput {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *RelationshipsForStarInput) GetUser() *RelationshipsForStarUserInput {
	if o == nil {
		return nil
	}
	return o.User
}
