// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
	"time"
)

// ContactType - The type of contact.
type ContactType string

const (
	ContactTypePrimary           ContactType = "primary"
	ContactTypeBilling           ContactType = "billing"
	ContactTypeTechnical         ContactType = "technical"
	ContactTypeSecurity          ContactType = "security"
	ContactTypeEmergency         ContactType = "emergency"
	ContactTypeGeneralCompliance ContactType = "general compliance"
)

func (e ContactType) ToPointer() *ContactType {
	return &e
}

func (e *ContactType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "primary":
		fallthrough
	case "billing":
		fallthrough
	case "technical":
		fallthrough
	case "security":
		fallthrough
	case "emergency":
		fallthrough
	case "general compliance":
		*e = ContactType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ContactType: %v", v)
	}
}

type SchemasContactResponse struct {
	// The type of contact.
	ContactType *ContactType `json:"contact_type,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The alphanumeric string representing the customer for this customer contact.
	CustomerID *string `json:"customer_id,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The email of this contact, when a user_id is not provided.
	Email *string `json:"email,omitempty"`
	ID    *string `json:"id,omitempty"`
	// The name of this contact, when user_id is not provided.
	Name *string `json:"name,omitempty"`
	// The phone number for this contact. Required for primary, technical, and security contact types.
	Phone *string `json:"phone,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The alphanumeric string representing the user for this customer contact.
	UserID *string `json:"user_id,omitempty"`
}

func (s SchemasContactResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SchemasContactResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SchemasContactResponse) GetContactType() *ContactType {
	if o == nil {
		return nil
	}
	return o.ContactType
}

func (o *SchemasContactResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SchemasContactResponse) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *SchemasContactResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *SchemasContactResponse) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *SchemasContactResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SchemasContactResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SchemasContactResponse) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *SchemasContactResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *SchemasContactResponse) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}
