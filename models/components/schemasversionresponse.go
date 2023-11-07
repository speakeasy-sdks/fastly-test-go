// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"time"
)

type SchemasVersionResponse struct {
	// Whether this is the active version or not.
	Active *bool `default:"false" json:"active"`
	// A freeform descriptive note.
	Comment *string `json:"comment,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Unused at this time.
	Deployed *bool `json:"deployed,omitempty"`
	// Whether this version is locked or not. Objects can not be added or edited on locked versions.
	Locked *bool `default:"false" json:"locked"`
	// The number of this version.
	Number    *int64  `json:"number,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	// Unused at this time.
	Staging *bool `default:"false" json:"staging"`
	// Unused at this time.
	Testing *bool `default:"false" json:"testing"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (s SchemasVersionResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SchemasVersionResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SchemasVersionResponse) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *SchemasVersionResponse) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *SchemasVersionResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SchemasVersionResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *SchemasVersionResponse) GetDeployed() *bool {
	if o == nil {
		return nil
	}
	return o.Deployed
}

func (o *SchemasVersionResponse) GetLocked() *bool {
	if o == nil {
		return nil
	}
	return o.Locked
}

func (o *SchemasVersionResponse) GetNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *SchemasVersionResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *SchemasVersionResponse) GetStaging() *bool {
	if o == nil {
		return nil
	}
	return o.Staging
}

func (o *SchemasVersionResponse) GetTesting() *bool {
	if o == nil {
		return nil
	}
	return o.Testing
}

func (o *SchemasVersionResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
