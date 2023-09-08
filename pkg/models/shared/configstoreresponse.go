// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type ConfigStoreResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// An alphanumeric string identifying the config store.
	ID *string `json:"id,omitempty"`
	// The name of the config store.
	Name *string `json:"name,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (o *ConfigStoreResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ConfigStoreResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *ConfigStoreResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ConfigStoreResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConfigStoreResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
