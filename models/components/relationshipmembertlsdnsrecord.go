// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type RelationshipMemberTLSDNSRecord struct {
	ID *string `json:"id,omitempty"`
	// Resource type
	Type *TypeTLSDNSRecord `default:"dns_record" json:"type"`
}

func (r RelationshipMemberTLSDNSRecord) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationshipMemberTLSDNSRecord) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RelationshipMemberTLSDNSRecord) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RelationshipMemberTLSDNSRecord) GetType() *TypeTLSDNSRecord {
	if o == nil {
		return nil
	}
	return o.Type
}
