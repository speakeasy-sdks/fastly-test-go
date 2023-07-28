// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// GzipResponse - OK
type GzipResponse struct {
	// Name of the cache condition controlling when this configuration applies.
	CacheCondition *string `json:"cache_condition,omitempty"`
	// Space-separated list of content types to compress. If you omit this field a default list will be used.
	ContentTypes *string `json:"content_types,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Space-separated list of file extensions to compress. If you omit this field a default list will be used.
	Extensions *string `json:"extensions,omitempty"`
	// Name of the gzip configuration.
	Name      *string `json:"name,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}

func (o *GzipResponse) GetCacheCondition() *string {
	if o == nil {
		return nil
	}
	return o.CacheCondition
}

func (o *GzipResponse) GetContentTypes() *string {
	if o == nil {
		return nil
	}
	return o.ContentTypes
}

func (o *GzipResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GzipResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *GzipResponse) GetExtensions() *string {
	if o == nil {
		return nil
	}
	return o.Extensions
}

func (o *GzipResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GzipResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *GzipResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GzipResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
