// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type StarData struct {
	Relationships *RelationshipsForStarInput `json:"relationships,omitempty"`
	// Resource type
	Type *TypeStar `default:"star" json:"type"`
}

func (s StarData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *StarData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *StarData) GetRelationships() *RelationshipsForStarInput {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *StarData) GetType() *TypeStar {
	if o == nil {
		return nil
	}
	return o.Type
}

type Star struct {
	Data *StarData `json:"data,omitempty"`
}

func (o *Star) GetData() *StarData {
	if o == nil {
		return nil
	}
	return o.Data
}
