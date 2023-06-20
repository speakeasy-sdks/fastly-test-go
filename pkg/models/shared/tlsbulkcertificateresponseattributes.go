// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type TLSBulkCertificateResponseAttributes struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Time-stamp (GMT) when the certificate will expire. Must be in the future to be used to terminate TLS traffic.
	NotAfter *time.Time `json:"not_after,omitempty"`
	// Time-stamp (GMT) when the certificate will become valid. Must be in the past to be used to terminate TLS traffic.
	NotBefore *time.Time `json:"not_before,omitempty"`
	// A recommendation from Fastly indicating the key associated with this certificate is in need of rotation.
	Replace *bool `json:"replace,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}