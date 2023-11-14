// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"time"
)

type Http3 struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Revision number of the HTTP/3 feature implementation. Defaults to the most recent revision.
	FeatureRevision *int64  `json:"feature_revision,omitempty"`
	ServiceID       *string `json:"service_id,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}

func (h Http3) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *Http3) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Http3) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Http3) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Http3) GetFeatureRevision() *int64 {
	if o == nil {
		return nil
	}
	return o.FeatureRevision
}

func (o *Http3) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *Http3) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Http3) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}