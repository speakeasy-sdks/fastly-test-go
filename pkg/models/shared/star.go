// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
)

type StarDataInput struct {
	Relationships *RelationshipsForStarInput `json:"relationships,omitempty"`
	// Resource type
	Type *TypeStar `default:"star" json:"type"`
}

func (s StarDataInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *StarDataInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *StarDataInput) GetRelationships() *RelationshipsForStarInput {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *StarDataInput) GetType() *TypeStar {
	if o == nil {
		return nil
	}
	return o.Type
}

type StarInput struct {
	Data *StarDataInput `json:"data,omitempty"`
}

func (o *StarInput) GetData() *StarDataInput {
	if o == nil {
		return nil
	}
	return o.Data
}
