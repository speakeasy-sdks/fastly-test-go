// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"encoding/json"
	"fmt"
	"time"
)

// StatusCode - HTTP status code used to redirect the client.
type StatusCode int64

const (
	StatusCodeThreeHundredAndOne   StatusCode = 301
	StatusCodeThreeHundredAndTwo   StatusCode = 302
	StatusCodeThreeHundredAndSeven StatusCode = 307
	StatusCodeThreeHundredAndEight StatusCode = 308
)

func (e StatusCode) ToPointer() *StatusCode {
	return &e
}

func (e *StatusCode) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 301:
		fallthrough
	case 302:
		fallthrough
	case 307:
		fallthrough
	case 308:
		*e = StatusCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatusCode: %v", v)
	}
}

type ApexRedirect struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Array of apex domains that should redirect to their WWW subdomain.
	Domains []string `json:"domains,omitempty"`
	// Revision number of the apex redirect feature implementation. Defaults to the most recent revision.
	FeatureRevision *int64  `json:"feature_revision,omitempty"`
	ServiceID       *string `json:"service_id,omitempty"`
	// HTTP status code used to redirect the client.
	StatusCode *StatusCode `json:"status_code,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}

func (a ApexRedirect) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApexRedirect) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApexRedirect) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ApexRedirect) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *ApexRedirect) GetDomains() []string {
	if o == nil {
		return nil
	}
	return o.Domains
}

func (o *ApexRedirect) GetFeatureRevision() *int64 {
	if o == nil {
		return nil
	}
	return o.FeatureRevision
}

func (o *ApexRedirect) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *ApexRedirect) GetStatusCode() *StatusCode {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

func (o *ApexRedirect) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ApexRedirect) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}

type ApexRedirectInput struct {
	// Array of apex domains that should redirect to their WWW subdomain.
	Domains []string `form:"name=domains"`
	// Revision number of the apex redirect feature implementation. Defaults to the most recent revision.
	FeatureRevision *int64 `form:"name=feature_revision"`
	// HTTP status code used to redirect the client.
	StatusCode *StatusCode `form:"name=status_code"`
}

func (o *ApexRedirectInput) GetDomains() []string {
	if o == nil {
		return nil
	}
	return o.Domains
}

func (o *ApexRedirectInput) GetFeatureRevision() *int64 {
	if o == nil {
		return nil
	}
	return o.FeatureRevision
}

func (o *ApexRedirectInput) GetStatusCode() *StatusCode {
	if o == nil {
		return nil
	}
	return o.StatusCode
}
