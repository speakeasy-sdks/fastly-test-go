// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"time"
)

type VersionResponse struct {
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

func (v VersionResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *VersionResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *VersionResponse) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *VersionResponse) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *VersionResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *VersionResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *VersionResponse) GetDeployed() *bool {
	if o == nil {
		return nil
	}
	return o.Deployed
}

func (o *VersionResponse) GetLocked() *bool {
	if o == nil {
		return nil
	}
	return o.Locked
}

func (o *VersionResponse) GetNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *VersionResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *VersionResponse) GetStaging() *bool {
	if o == nil {
		return nil
	}
	return o.Staging
}

func (o *VersionResponse) GetTesting() *bool {
	if o == nil {
		return nil
	}
	return o.Testing
}

func (o *VersionResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}