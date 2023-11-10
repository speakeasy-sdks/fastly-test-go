// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"time"
)

type VclResponse struct {
	// The VCL code to be included.
	Content *string `json:"content,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Set to `true` when this is the main VCL, otherwise `false`.
	Main *bool `json:"main,omitempty"`
	// The name of this VCL.
	Name      *string `json:"name,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}

func (v VclResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *VclResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *VclResponse) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *VclResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *VclResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *VclResponse) GetMain() *bool {
	if o == nil {
		return nil
	}
	return o.Main
}

func (o *VclResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *VclResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *VclResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *VclResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
