// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type SchemasUserResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	CustomerID *string    `json:"customer_id,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The alphanumeric string identifying a email login.
	EmailHash *string `json:"email_hash,omitempty"`
	ID        *string `json:"id,omitempty"`
	// Indicates that the user has limited access to the customer's services.
	LimitServices *bool `json:"limit_services,omitempty"`
	// Indicates whether the is account is locked for editing or not.
	Locked *bool   `json:"locked,omitempty"`
	Login  *string `json:"login,omitempty"`
	// The real life name of the user.
	Name *string `json:"name,omitempty"`
	// Indicates if a new password is required at next login.
	RequireNewPassword *bool `json:"require_new_password,omitempty"`
	// The permissions role assigned to the user. Can be `user`, `billing`, `engineer`, or `superuser`.
	Role *RoleUser `json:"role,omitempty"`
	// Indicates if 2FA is enabled on the user.
	TwoFactorAuthEnabled *bool `json:"two_factor_auth_enabled,omitempty"`
	// Indicates if 2FA is required by the user's customer account.
	TwoFactorSetupRequired *bool `json:"two_factor_setup_required,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}