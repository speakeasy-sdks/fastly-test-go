// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"time"
)

type ResourceResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The path to the resource.
	Href *string `json:"href,omitempty"`
	// An alphanumeric string identifying the resource link.
	ID *string `json:"id,omitempty"`
	// The name of the resource link.
	Name *string `json:"name,omitempty"`
	// The ID of the underlying linked resource.
	ResourceID *string `json:"resource_id,omitempty"`
	// Resource type
	ResourceType *TypeResource `default:"object-store" json:"resource_type"`
	// Alphanumeric string identifying the service.
	ServiceID *string `json:"service_id,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Integer identifying a service version.
	Version *int64 `json:"version,omitempty"`
}

func (r ResourceResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ResourceResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ResourceResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ResourceResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *ResourceResponse) GetHref() *string {
	if o == nil {
		return nil
	}
	return o.Href
}

func (o *ResourceResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ResourceResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ResourceResponse) GetResourceID() *string {
	if o == nil {
		return nil
	}
	return o.ResourceID
}

func (o *ResourceResponse) GetResourceType() *TypeResource {
	if o == nil {
		return nil
	}
	return o.ResourceType
}

func (o *ResourceResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *ResourceResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ResourceResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
